package routes

import (
	"cook-book-backend/config"
	"cook-book-backend/controllers"
	"cook-book-backend/respositories"
	"cook-book-backend/services"
	"github.com/gin-gonic/gin"
)

func SetupArticleCateRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	articleCateRepo := respositories.NewArticleCateRepo(config.DB)
	articleCateSrv := services.NewArticleCateSrv(articleCateRepo)
	articleCateCtrl := controllers.NewArticleCateCtrl(articleCateSrv)

	r.POST("/app/categories/list", articleCateCtrl.GetArticleCateList)
}
