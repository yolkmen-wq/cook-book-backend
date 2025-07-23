package routes

import (
	"cook-book-backend/internal/container"

	"github.com/gin-gonic/gin"
)

func SetupArticleCarouselRoutes(r *gin.Engine, rg *gin.RouterGroup, container *container.Container) {
	// articleCarouselRepo := repositories.NewArticleCarouselRepo(config.DB)
	// articleCarouselSrv := services.NewArticleCarouselSrv(articleCarouselRepo)
	// articleCarouselCtrl := controllers.NewArticleCarouselCtrl(articleCarouselSrv)

	r.POST("/app/carousel/list", container.ArticleCarouselHandler.GetArticleCarouselList)
}
