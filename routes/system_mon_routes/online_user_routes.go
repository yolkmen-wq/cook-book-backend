package system_mon_routes

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/controllers/system_mon_ctrl"
	"cook-book-admin-backend/respositories/system_mon_repo"
	"cook-book-admin-backend/services/system_mon_srv"
	"github.com/gin-gonic/gin"
)

func SetupOnlineUserRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	onlineUserController := system_mon_ctrl.NewOnlineUserController(system_mon_srv.NewOnlineUserService(system_mon_repo.NewOnlineUserRepository(config.DB)))

	rg.POST("/admin/online-user", onlineUserController.GetOnlineUsers)

}
