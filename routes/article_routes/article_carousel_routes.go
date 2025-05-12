package article_routes

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/controllers/article_ctrl"
	"cook-book-admin-backend/respositories/article_repo"
	"cook-book-admin-backend/services/article_srv"
	"github.com/gin-gonic/gin"
)

func SetupCarouselRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	articleCatController := article_ctrl.NewCarouselController(article_srv.NewCarouselService(article_repo.NewCarouselRepository(config.DB)))
	r.POST("/admin/carousel", articleCatController.GetCarousels)
	r.POST("/admin/carousel/create", articleCatController.CreateCarousel)
	r.POST("/admin/carousel/update", articleCatController.UpdateCarousel)
	r.GET("/admin/carousel/delete/:id", articleCatController.DeleteCarousel)
	r.POST("/admin/carousel/item", articleCatController.GetCarouselItems)
	r.POST("/admin/carousel/item/create", articleCatController.CreateCarouselItem)
	r.POST("/admin/carousel/item/update", articleCatController.UpdateCarouselItem)
	r.GET("/admin/carousel/item/delete/:id", articleCatController.DeleteCarouselItem)
}
