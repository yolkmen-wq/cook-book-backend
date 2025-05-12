package system_mon_routes

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/controllers/system_mon_ctrl"
	"cook-book-admin-backend/respositories/system_mon_repo"
	"cook-book-admin-backend/services/system_mon_srv"
	"github.com/gin-gonic/gin"
)

func SetupSystemLogRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	systemLogController := system_mon_ctrl.NewSystemLogController(system_mon_srv.NewSystemLogService(system_mon_repo.NewSystemLogRepository(config.DB)))
	rg.POST("/admin/system-logs", systemLogController.GetSystemLogs)
	rg.POST("/admin/system-logs/delete", systemLogController.DeleteSystemLogs)
}
