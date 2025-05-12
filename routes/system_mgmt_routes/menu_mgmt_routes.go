package system_mgmt_routes

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/controllers/system_mgmt_ctrl"
	"cook-book-admin-backend/respositories/system_mgmt_repo"
	"cook-book-admin-backend/services/system_mgmt_srv"
	"github.com/gin-gonic/gin"
)

func SetupMenuMgmtRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	menuMgmtController := system_mgmt_ctrl.NewMenuMgmtController(system_mgmt_srv.NewMenuMgmtService(system_mgmt_repo.NewMenuMgmtRepository(config.DB)))
	rg.POST("/admin/menu", menuMgmtController.GetMenus)
	rg.GET("admin/menu/detail", menuMgmtController.GetMenuDetail)
	rg.POST("admin/menu/update", menuMgmtController.UpdateMenu)
	rg.POST("admin/menu/create", menuMgmtController.AddMenu)
	rg.GET("admin/menu/delete/:id", menuMgmtController.DeleteMenu)
	rg.GET("/admin/get-async-routes", menuMgmtController.GetAsyncRoutes)

}
