package system_mgmt_routes

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/controllers/system_mgmt_ctrl"
	"cook-book-admin-backend/respositories/system_mgmt_repo"
	"cook-book-admin-backend/services/system_mgmt_srv"
	"github.com/gin-gonic/gin"
)

func SetupDictMgmtRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	dictMgmtController := system_mgmt_ctrl.NewDictMgmtController(system_mgmt_srv.NewDictMgmtService(system_mgmt_repo.NewDictMgmtRepository(config.DB)))
	rg.POST("/admin/dict", dictMgmtController.GetDictList)
	rg.POST("/admin/dict/create", dictMgmtController.CreateDict)
	rg.POST("/admin/dict/update", dictMgmtController.UpdateDict)
	rg.GET("/admin/dict/delete/:id", dictMgmtController.DeleteDict)
	rg.POST("/admin/dict/dictData", dictMgmtController.GetDictData)
	rg.POST("/admin/dict/dictData/add", dictMgmtController.AddDictData)
	rg.POST("/admin/dict/dictData/update", dictMgmtController.UpdateDictData)
	rg.GET("/admin/dict/dictData/delete/:id", dictMgmtController.DeleteDictData)
}
