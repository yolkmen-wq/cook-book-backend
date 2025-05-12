package system_mgmt_routes

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/controllers/system_mgmt_ctrl"
	"cook-book-admin-backend/respositories/system_mgmt_repo"
	"cook-book-admin-backend/services/system_mgmt_srv"
	"github.com/gin-gonic/gin"
)

func SetupUserMgmtRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	userMgmtController := system_mgmt_ctrl.NewUserMgmtController(system_mgmt_srv.NewUserMgmtService(system_mgmt_repo.NewUserMgmtRepository(config.DB)))

	rg.POST("/admin/user", userMgmtController.GetUsers)
	rg.POST("/admin/user/update", userMgmtController.UpdateUser)
	rg.POST("/admin/user/add", userMgmtController.AddUser)
	rg.GET("/admin/user/delete/:id", userMgmtController.DeleteUser)
	rg.POST("/admin/list-all-role", userMgmtController.GetRoles)
	rg.POST("/admin/list-role-ids", userMgmtController.GetRolesByIds)
	rg.POST("/admin/user/assignRole", userMgmtController.AssignRolesToUser)
}
