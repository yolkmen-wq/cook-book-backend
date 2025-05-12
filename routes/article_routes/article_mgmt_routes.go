package article_routes

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/controllers/article_ctrl"
	"cook-book-admin-backend/respositories/article_repo"
	"cook-book-admin-backend/services/article_srv"
	"github.com/gin-gonic/gin"
)

func SetupArticleMgmtRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	articleMgmtController := article_ctrl.NewArticleMgmtController(article_srv.NewArticleMgmtService(article_repo.NewArticleMgmtRepository(config.DB)))

	r.POST("/admin/article", articleMgmtController.GetArticleList)
	r.POST("/admin/article/create", articleMgmtController.CreateArticle)
	r.POST("/admin/article/update", articleMgmtController.UpdateArticle)
	r.GET("/admin/article/delete/:id", articleMgmtController.DeleteArticle)
}
