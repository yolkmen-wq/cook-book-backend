package routes

import (
	"cook-book-backend/internal/container"

	"github.com/gin-gonic/gin"
)

func SetupArticleCateRoutes(r *gin.Engine, rg *gin.RouterGroup, container *container.Container) {
	r.POST("/app/categories/list", container.ArticleCateHandler.GetArticleCateList)
}
