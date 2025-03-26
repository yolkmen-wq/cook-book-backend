package system_mon_ctrl

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/services/system_mon_srv"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OnlineUserController struct {
	onlineUserService system_mon_srv.OnlineUserService
}

func NewOnlineUserController(onlineUserService system_mon_srv.OnlineUserService) *OnlineUserController {
	return &OnlineUserController{
		onlineUserService: onlineUserService,
	}
}

func (m *OnlineUserController) GetOnlineUsers(c *gin.Context) {
	var getUsersRequest models.GetUsersRequest
	err := c.ShouldBind(&getUsersRequest)

	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	users, total, pageNum, pageSize, err := m.onlineUserService.GetOnlineUsers(getUsersRequest)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", config.ListResponse{
		List:        users,
		Total:       total,
		CurrentPage: pageNum,
		PageSize:    pageSize,
	})
	c.JSONP(http.StatusOK, response)
}
