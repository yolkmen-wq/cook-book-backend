package routes

import (
	"cook-book-backEnd/routes/system_mgt_routes"
	"cook-book-backEnd/routes/system_mon_routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	// Add your routes here
	SetupUserRoutes(r, rg)
	system_mgt_routes.SetupUserMgmtRoutes(r, rg)
	system_mgt_routes.SetupRoleMgmtRoutes(r, rg)
	system_mgt_routes.SetupMenuMgmtRoutes(r, rg)
	system_mon_routes.SetupOnlineUserRoutes(r, rg)
	system_mon_routes.SetupSystemLogRoutes(r, rg)
}
