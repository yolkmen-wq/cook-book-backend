package article_routes

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/controllers/article_ctrl"
	"cook-book-admin-backend/respositories/article_repo"
	"cook-book-admin-backend/services/article_srv"
	"github.com/gin-gonic/gin"
)

func SetupArticleCatRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	articleCatController := article_ctrl.NewArticleCatController(article_srv.NewArticleCatService(article_repo.NewArticleCatRepository(config.DB)))
	r.POST("/admin/article-cat", articleCatController.GetArticleCatList)
	r.POST("/admin/article-cat/create", articleCatController.CreateArticleCat)
	r.POST("/admin/article-cat/update", articleCatController.UpdateArticleCat)
	r.GET("/admin/article-cat/delete/:id", articleCatController.DeleteArticleCat)

}
