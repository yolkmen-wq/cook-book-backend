package routes

import (
	"cook-book-admin-backend/routes/article_routes"
	"cook-book-admin-backend/routes/system_mgmt_routes"
	"cook-book-admin-backend/routes/system_mon_routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	// Add your routes here
	SetupUserRoutes(r, rg)
	SetupFileRoutes(r, rg)
	system_mgmt_routes.SetupUserMgmtRoutes(r, rg)
	system_mgmt_routes.SetupRoleMgmtRoutes(r, rg)
	system_mgmt_routes.SetupMenuMgmtRoutes(r, rg)
	system_mgmt_routes.SetupDictMgmtRoutes(r, rg)
	system_mon_routes.SetupOnlineUserRoutes(r, rg)
	system_mon_routes.SetupSystemLogRoutes(r, rg)
	article_routes.SetupArticleMgmtRoutes(r, rg)
	article_routes.SetupArticleCatRoutes(r, rg)
	article_routes.SetupCarouselRoutes(r, rg)
}
