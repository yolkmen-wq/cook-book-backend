package system_mgmt_ctrl

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/services/system_mgmt_srv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserMgmtController struct {
	userMgmtService system_mgmt_srv.UserMgmtService
}

func NewUserMgmtController(userMgmtService system_mgmt_srv.UserMgmtService) *UserMgmtController {
	return &UserMgmtController{userMgmtService}
}

// GetUsers 获取用户列表
func (umc *UserMgmtController) GetUsers(c *gin.Context) {
	var getUsersRequest models.GetUsersRequest

	err := c.ShouldBind(&getUsersRequest)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	list, total, pageSize, pageNum, err := umc.userMgmtService.GetAdminUserList(getUsersRequest)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", &config.ListResponse{
		List:        list,
		Total:       total,
		PageSize:    pageSize,
		CurrentPage: pageNum,
	})
	c.JSONP(http.StatusOK, response)
}

// AddUser 添加用户
func (umc *UserMgmtController) AddUser(c *gin.Context) {
	var addUserRequest models.AdminUserMgmt
	err := c.ShouldBind(&addUserRequest)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	err = umc.userMgmtService.AddUser(addUserRequest)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "添加成功", nil)
	c.JSON(http.StatusOK, response)
}

// DeleteUser 删除用户
func (uc *UserMgmtController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	err = uc.userMgmtService.DeleteUser(idInt)
	// 返回数据
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	response := config.NewResponse(http.StatusOK, true, "删除成功", nil)
	c.JSON(http.StatusOK, response)
}

// UpdateUser 更新用户
func (umc *UserMgmtController) UpdateUser(c *gin.Context) {

	var updateUserRequest models.AdminUser
	err := c.ShouldBind(&updateUserRequest)
	fmt.Println("updateUserRequest", updateUserRequest)
	err = umc.userMgmtService.UpdateUser(updateUserRequest)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "更新成功", nil)
	c.JSONP(http.StatusOK, response)
}

// GetRoles 获取角色列表
func (umc *UserMgmtController) GetRoles(c *gin.Context) {
	type GetAllRoleRequest struct {
		Name     string `json:"name"`
		PageNum  int    `json:"pageNum"`
		PageSize int    `json:"pageSize"`
	}
	var getAllRoleRequest GetAllRoleRequest
	err := c.ShouldBind(&getAllRoleRequest)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, err)
		return
	}

	list, err := umc.userMgmtService.GetRoleList(getAllRoleRequest.Name, 0)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, err)
		return
	}
	// 返回数据
	var roleList []interface{}
	for _, role := range list {
		roleList = append(roleList, map[string]interface{}{"id": role.ID, "name": role.Name})
	}
	response := config.NewResponse(http.StatusOK, true, "获取成功", map[string]interface{}{
		"list": roleList,
	})
	c.JSONP(http.StatusOK, response)
}

// GetRolesByIds 获取用户角色列表
func (umc *UserMgmtController) GetRolesByIds(c *gin.Context) {
	type GetRolesByIdsRequest struct {
		UserID int64 `json:"userId"`
	}
	var getRolesByIdsRequest GetRolesByIdsRequest
	err := c.ShouldBind(&getRolesByIdsRequest)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, err)
		return
	}
	list, err := umc.userMgmtService.GetRoleList("", getRolesByIdsRequest.UserID)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	var ids []int64
	for _, role := range list {
		ids = append(ids, role.ID)
	}
	response := config.NewResponse(http.StatusOK, true, "获取成功", ids)
	c.JSON(http.StatusOK, response)

}

// AssignRolesToUser 给用户分配角色
func (umc *UserMgmtController) AssignRolesToUser(c *gin.Context) {
	type AssignRolesToUserRequest struct {
		UserID  int64   `json:"userId"`
		RoleIds []int64 `json:"roleIds"`
	}
	var assignRolesToUserRequest AssignRolesToUserRequest
	err := c.ShouldBind(&assignRolesToUserRequest)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, err)
		return
	}
	err = umc.userMgmtService.UpdateUserRoles(assignRolesToUserRequest.UserID, assignRolesToUserRequest.RoleIds)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "分配成功", nil)
	c.JSON(http.StatusOK, response)
}
