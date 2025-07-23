package routes

import (
	"cook-book-backend/internal/infrastructure/middlewares" // 假设 middlewares 在此包

	"github.com/gin-gonic/gin"
	// 导入其他需要的包，如 handlers
	"cook-book-backend/internal/container" // Add this
)

func Setup(router *gin.Engine, c *container.Container) {
	// Use c instead of container
	// 创建授权路由组
	authorized := router.Group("/")
	authorized.Use(middlewares.AuthMiddleWare())

	// 初始化各个路由，使用 container 获取依赖
	SetupArticleRoutes(router, authorized, c)
	// SetupArticleCateRoutes(router, authorized, container)
	// SetupArticleCarouselRoutes(router, authorized, container)
	// SetupEmojiRoutes(router, authorized, container)
	// SetupCommentRoutes(router, authorized, container)
	// 如果有用户路由，可以添加 SetupUserRoutes(router, authorized, container)
}

// func InitRoutes(r *gin.Engine, rg *gin.RouterGroup) {
// 	// Add your routes here
// 	SetupArticleRoutes(r, rg)
// 	SetupArticleCateRoutes(r, rg)
// 	SetupArticleCarouselRoutes(r, rg)
// 	SetupEmojiRoutes(r, rg)
// 	SetupCommentRoutes(r, rg)
// }
