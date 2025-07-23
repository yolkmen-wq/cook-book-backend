package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cook-book-backend/internal/config"
	"cook-book-backend/internal/infrastructure/database"
	"cook-book-backend/internal/infrastructure/logger"
	"cook-book-backend/internal/infrastructure/middlewares"
	"cook-book-backend/internal/interfaces/http/routes"

	"cook-book-backend/internal/container" // Add this

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	config *config.Config
	db     *gorm.DB
	server *http.Server
	logger logger.Logger
}

func New() (*App, error) {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	// 初始化日志
	loggerInstance := logger.New(logger.Config{
		Level:      cfg.Logger.Level,
		Format:     cfg.Logger.Format,
		Output:     cfg.Logger.Output,
		Filename:   cfg.Logger.Filename,
		MaxSize:    cfg.Logger.MaxSize,
		MaxBackups: cfg.Logger.MaxBackups,
		MaxAge:     cfg.Logger.MaxAge,
		// Compress:   cfg.Logger.Compress,
	})

	// 设置全局日志器
	logger.SetGlobalLogger(loggerInstance)

	// 初始化数据库
	db, err := database.New(cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return &App{
		config: cfg,
		db:     db,
		logger: loggerInstance,
	}, nil
}

func (a *App) Run() error {
	// 设置 Gin 模式
	if a.config.Logger.Level == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建路由
	router := gin.New()
	authorized := router.Group("/")
	authorized.Use(middlewares.AuthMiddleWare())
	// 设置中间件
	router.Use(gin.Recovery())
	router.Use(logger.GinLogger(a.logger))

	// 初始化依赖注入容器
	container := a.buildContainer()

	// 设置路由
	routes.Setup(router, container)
	// 创建服务器
	a.server = &http.Server{
		Addr:         ":" + a.config.Server.Port,
		Handler:      router,
		ReadTimeout:  a.config.Server.ReadTimeout,
		WriteTimeout: a.config.Server.WriteTimeout,
	}

	// 启动服务器
	go func() {
		a.logger.Info("Server starting on port " + a.config.Server.Port)
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal("Failed to start server: " + err.Error())
		}
	}()

	// 等待中断信号
	return a.gracefulShutdown()
}

func (a *App) gracefulShutdown() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	a.logger.Info("Server exited")
	return nil
}

func (a *App) buildContainer() *container.Container {
	return container.NewContainer(a.db, a.config, a.logger)
}
