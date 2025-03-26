package system_mon_ctrl

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/services/system_mon_srv"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SystemLogController struct {
	SystemLogService system_mon_srv.SystemLogService
}

func NewSystemLogController(systemLogService system_mon_srv.SystemLogService) *SystemLogController {
	return &SystemLogController{
		SystemLogService: systemLogService,
	}
}

func (slc *SystemLogController) DeleteSystemLogs(c *gin.Context) {
	err := slc.SystemLogService.DeleteSystemLogs()
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "删除成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (slc *SystemLogController) GetSystemLogs(c *gin.Context) {
	list, total, err := slc.SystemLogService.GetSystemLogs(1, 10)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", map[string]interface{}{
		"list":  list,
		"total": total,
	})
	c.JSONP(http.StatusOK, response)
}
