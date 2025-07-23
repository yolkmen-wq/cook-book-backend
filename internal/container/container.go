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
	UserRepo            repositories.UserRepository
	ArticleRepo         repositories.ArticleRepository
	ArticleCarouselRepo repositories.ArticleCarouselRepository
	ArticleCateRepo     repositories.ArticleCateRepository
	CommentRepo         repositories.CommentRepository
	EmojiRepo           repositories.EmojiRepository

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
	articleCarouselRepo := repositories.NewArticleCarouselRepository(db)
	articleCateRepo := repositories.NewArticleCateRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	emojiRepo := repositories.NewEmojiRepository(db)

	// 初始化 Services
	userService := services.NewUserService(userRepo)
	articleService := services.NewArticleService(articleRepo)
	articleCarouselService := services.NewArticleCarouselService(articleCarouselRepo)
	articleCateService := services.NewArticleCateService(articleCateRepo)
	commentService := services.NewCommentService(commentRepo)
	emojiService := services.NewEmojiService(emojiRepo)

	// 初始化 Handlers
	userHandler := handlers.NewUserHandler(userService, logger)
	articleHandler := handlers.NewArticleHandler(articleService, logger)
	articleCarouselHandler := handlers.NewArticleCarouselHandler(articleCarouselService, logger)
	articleCateHandler := handlers.NewArticleCateHandler(articleCateService, logger)
	commentCateHandler := handlers.NewCommentHandler(commentService, logger)
	emojiHandler := handlers.NewEmojiHandler(emojiService, logger)

	return &Container{
		UserRepo:               userRepo,
		ArticleRepo:            articleRepo,
		ArticleCarouselRepo:    articleCarouselRepo,
		ArticleCateRepo:        articleCateRepo,
		CommentRepo:            commentRepo,
		EmojiRepo:              emojiRepo,
		UserService:            userService,
		ArticleService:         articleService,
		ArticleCarouselService: articleCarouselService,
		ArticleCateService:     articleCateService,
		CommentService:         commentService,
		EmojiService:           emojiService,
		UserHandler:            userHandler,
		ArticleHandler:         articleHandler,
		ArticleCarouselHandler: articleCarouselHandler,
		ArticleCateHandler:     articleCateHandler,
		CommentHandler:         commentCateHandler,
		EmojiHandler:           emojiHandler,
	}
}
