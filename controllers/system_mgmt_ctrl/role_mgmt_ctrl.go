package system_mgmt_ctrl

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/services/system_mgmt_srv"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoleMgmtController struct {
	roleMgmtService system_mgmt_srv.RoleMgmtService
}

func NewRoleMgmtController(roleMgmtService system_mgmt_srv.RoleMgmtService) *RoleMgmtController {
	return &RoleMgmtController{roleMgmtService}
}

func (roleMgmtController *RoleMgmtController) GetRoles(c *gin.Context) {
	var getRolesRequest models.GetRolesRequest
	err := c.ShouldBind(&getRolesRequest)

	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	roles, total, pageNum, pageSize, err := roleMgmtController.roleMgmtService.GetRoleList(getRolesRequest)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", config.ListResponse{
		List:        roles,
		Total:       total,
		CurrentPage: pageNum,
		PageSize:    pageSize,
	})
	c.JSONP(http.StatusOK, response)
}

func (roleMgmtController *RoleMgmtController) UpdateRole(c *gin.Context) {
	var role models.Role
	err := c.ShouldBind(&role)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	err = roleMgmtController.roleMgmtService.UpdateRole(role)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "更新成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (roleMgmtController *RoleMgmtController) CreateRole(c *gin.Context) {
	var role models.Role
	err := c.ShouldBind(&role)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	err = roleMgmtController.roleMgmtService.CreateRole(role.Name, role.Code)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "创建成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (roleMgmtController *RoleMgmtController) DeleteRole(c *gin.Context) {
	roleId := c.Param("id")
	if roleId == "" {
		errResponse := config.NewResponse(http.StatusBadRequest, false, "参数错误", nil)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		errResponse := config.NewResponse(http.StatusBadRequest, false, "参数错误", nil)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	err = roleMgmtController.roleMgmtService.DeleteRole(roleIdInt)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "删除成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (roleMgmtController *RoleMgmtController) GetRoleMenuListByRoleId(c *gin.Context) {
	type RequestBody struct {
		RoleId int64 `json:"id"`
	}
	var requestBody RequestBody

	err := c.ShouldBind(&requestBody)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	menuList, err := roleMgmtController.roleMgmtService.GetRoleMenuListByRoleId(requestBody.RoleId)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", map[string]interface{}{
		"list": menuList,
	})
	c.JSONP(http.StatusOK, response)
}

func (roleMgmtController *RoleMgmtController) SaveRoleMenuPermission(c *gin.Context) {
	type RequestBody struct {
		RoleId  int64   `json:"roleId"`
		MenuIds []int64 `json:"menuIds"`
	}
	var requestBody RequestBody

	err := c.ShouldBind(&requestBody)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	err = roleMgmtController.roleMgmtService.SaveRoleMenuPermission(requestBody.RoleId, requestBody.MenuIds)
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "保存成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (roleMgmtController *RoleMgmtController) GetRoleCanAssignMenuList(c *gin.Context) {
	menuList, err := roleMgmtController.roleMgmtService.GetRoleCanAssignMenuList()
	if err != nil {
		errResponse := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", map[string]interface{}{
		"list": menuList,
	})
	c.JSONP(http.StatusOK, response)
}
