package container

import (
	"cook-book-backend/internal/config"
	"cook-book-backend/internal/domain/repositories"
	"cook-book-backend/internal/domain/services"
	"cook-book-backend/internal/infrastructure/logger"
	"cook-book-backend/internal/interfaces/http/handlers"

	"gorm.io/gorm"
)

type Container struct {
	// Repositories
	UserRepo        repositories.UserRepository
	ArticleRepo     repositories.ArticleRepository
	ArticleCarousel repositories.ArticleCarouselRepository
	ArticleCate     repositories.ArticleCateRepository
	Comment         repositories.CommentRepository
	Emoji           repositories.EmojiRepository

	// Services
	UserService            services.UserService
	ArticleService         services.ArticleService
	ArticleCarouselService services.ArticleCarouselService
	ArticleCateService     services.ArticleCateService
	CommentService         services.CommentService
	EmojiService           services.EmojiService

	// Handlers
	UserHandler            *handlers.UserHandler
	ArticleHandler         *handlers.ArticleHandler
	ArticleCarouselHandler *handlers.ArticleCarouselHandler
	ArticleCateHandler     *handlers.ArticleCateHandler
	CommentHandler         *handlers.CommentHandler
	EmojiHandler           *handlers.EmojiHandler
}

func NewContainer(db *gorm.DB, cfg *config.Config, logger logger.Logger) *Container {
	// 初始化 Repositories
	userRepo := repositories.NewUserRepository(db)
	articleRepo := repositories.NewArticleRepository(db)

	// 初始化 Services
	userService := services.NewUserService(userRepo)
	articleService := services.NewArticleService(articleRepo)

	// 初始化 Handlers
	userHandler := handlers.NewUserHandler(userService, logger)
	articleHandler := handlers.NewArticleHandler(articleService, logger)

	return &Container{
		UserRepo:       userRepo,
		ArticleRepo:    articleRepo,
		UserService:    userService,
		ArticleService: articleService,
		UserHandler:    userHandler,
		ArticleHandler: articleHandler,
	}
}
