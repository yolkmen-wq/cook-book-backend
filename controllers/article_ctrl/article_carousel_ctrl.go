package article_ctrl

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/services/article_srv"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CarouselController struct {
	carouselSrv article_srv.CarouselService
}

func NewCarouselController(carouselService article_srv.CarouselService) *CarouselController {
	return &CarouselController{
		carouselService,
	}
}

func (cc *CarouselController) GetCarousels(c *gin.Context) {
	var req models.GetCarouselsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	carousels, total, pageSize, pageNum, err := cc.carouselSrv.GetCarousels(&req)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", config.ListResponse{List: carousels, Total: total, CurrentPage: pageNum, PageSize: pageSize})
	c.JSONP(http.StatusOK, response)
}

func (cc *CarouselController) CreateCarousel(c *gin.Context) {
	var req models.Carousel
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = cc.carouselSrv.CreateCarousel(&req)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "创建成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (cc *CarouselController) UpdateCarousel(c *gin.Context) {
	var carousel models.Carousel
	err := c.ShouldBindJSON(&carousel)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = cc.carouselSrv.UpdateCarousel(&carousel)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "更新成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (cc *CarouselController) DeleteCarousel(c *gin.Context) {
	var id = c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = cc.carouselSrv.DeleteCarousel(idInt)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "删除成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (cc *CarouselController) GetCarouselItems(c *gin.Context) {
	var req models.GetCarouselItemsRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	carouselItems, total, pageSize, pageNum, err := cc.carouselSrv.GetCarouselItems(&req)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", config.ListResponse{List: carouselItems, Total: total, CurrentPage: pageNum, PageSize: pageSize})
	c.JSONP(http.StatusOK, response)
}

func (cc *CarouselController) CreateCarouselItem(c *gin.Context) {
	var req models.CarouselItem
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = cc.carouselSrv.CreateCarouselItem(&req)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "创建成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (cc *CarouselController) UpdateCarouselItem(c *gin.Context) {
	var carouselItem models.CarouselItem
	err := c.ShouldBindJSON(&carouselItem)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = cc.carouselSrv.UpdateCarouselItem(&carouselItem)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "更新成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (cc *CarouselController) DeleteCarouselItem(c *gin.Context) {
	var id = c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = cc.carouselSrv.DeleteCarouselItem(idInt)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "删除成功", nil)
	c.JSONP(http.StatusOK, response)
}
