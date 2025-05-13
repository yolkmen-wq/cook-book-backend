package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	// Add your routes here
	SetupArticleRoutes(r, rg)
	SetupArticleCateRoutes(r, rg)
	SetupArticleCarouselRoutes(r, rg)
}
