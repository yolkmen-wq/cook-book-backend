package system_mgmt_ctrl

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/services/system_mgmt_srv"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DictMgmtController struct {
	dictService system_mgmt_srv.DictMgmtService
}

func NewDictMgmtController(dictService system_mgmt_srv.DictMgmtService) *DictMgmtController {
	return &DictMgmtController{dictService: dictService}
}

func (dmc *DictMgmtController) GetDictList(c *gin.Context) {
	var req models.GetDictListRequest

	if err := c.ShouldBind(&req); err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dicts, total, pageNum, pageSize, err := dmc.dictService.GetDictList(&req)
	if err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "success", config.ListResponse{
		List:        dicts,
		Total:       total,
		CurrentPage: pageNum,
		PageSize:    pageSize,
	})
	c.JSONP(http.StatusOK, response)
}

func (dmc *DictMgmtController) CreateDict(c *gin.Context) {
	var dict models.Dict
	if err := c.ShouldBindJSON(&dict); err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := dmc.dictService.CreateDict(&dict)
	if err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "success", nil)
	c.JSONP(http.StatusOK, response)
}

func (dmc *DictMgmtController) UpdateDict(c *gin.Context) {
	var dict models.Dict
	if err := c.ShouldBindJSON(&dict); err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := dmc.dictService.UpdateDict(&dict)
	if err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "success", nil)
	c.JSONP(http.StatusOK, response)
}

func (dmc *DictMgmtController) DeleteDict(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err = dmc.dictService.DeleteDict(idInt); err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "success", nil)
	c.JSONP(http.StatusOK, response)
}

func (dmc *DictMgmtController) GetDictData(c *gin.Context) {
	var req models.GetDictDataListRequest

	if err := c.ShouldBind(&req); err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	dictData, total, pageNum, pageSize, err := dmc.dictService.GetDictData(&req)
	if err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "success", config.ListResponse{
		List:        dictData,
		Total:       total,
		CurrentPage: pageNum,
		PageSize:    pageSize,
	})
	c.JSONP(http.StatusOK, response)
}

func (dmc *DictMgmtController) AddDictData(c *gin.Context) {
	var dictData models.DictData
	if err := c.ShouldBindJSON(&dictData); err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := dmc.dictService.AddDictData(&dictData)
	if err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "success", nil)
	c.JSONP(http.StatusOK, response)
}

func (dmc *DictMgmtController) UpdateDictData(c *gin.Context) {
	var dictData models.DictData
	if err := c.ShouldBindJSON(&dictData); err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := dmc.dictService.UpdateDictData(&dictData)
	if err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "success", nil)
	c.JSONP(http.StatusOK, response)
}

func (dmc *DictMgmtController) DeleteDictData(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err = dmc.dictService.DeleteDictData(idInt); err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "success", nil)
	c.JSONP(http.StatusOK, response)
}
