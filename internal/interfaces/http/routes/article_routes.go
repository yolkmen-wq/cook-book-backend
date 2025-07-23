package routes

import (
	"cook-book-backend/internal/container"

	"github.com/gin-gonic/gin"
)

func SetupArticleRoutes(r *gin.Engine, rg *gin.RouterGroup, container *container.Container) {
	r.POST("/app/article/list", container.ArticleHandler.GetArticleList)
	r.POST("/app/article/latest/list", container.ArticleHandler.GetLatestArticles)
	r.GET("/app/article/:id", container.ArticleHandler.GetArticleDetail)
}
