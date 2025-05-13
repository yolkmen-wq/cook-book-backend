package routes

import (
	"cook-book-backend/config"
	"cook-book-backend/controllers"
	"cook-book-backend/respositories"
	"cook-book-backend/services"
	"github.com/gin-gonic/gin"
)

func SetupArticleCarouselRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	articleCarouselRepo := respositories.NewArticleCarouselRepo(config.DB)
	articleCarouselSrv := services.NewArticleCarouselSrv(articleCarouselRepo)
	articleCarouselCtrl := controllers.NewArticleCarouselCtrl(articleCarouselSrv)

	r.POST("/app/carousel/list", articleCarouselCtrl.GetArticleCarouselList)
}
