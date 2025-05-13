package controllers

import (
	"cook-book-backend/config"
	"cook-book-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleCarouselCtrl struct {
	articleCarouselSrv services.ArticleCarouselSrv
}

func NewArticleCarouselCtrl(articleCarouselSrv services.ArticleCarouselSrv) ArticleCarouselCtrl {
	return ArticleCarouselCtrl{
		articleCarouselSrv: articleCarouselSrv,
	}
}

func (acc ArticleCarouselCtrl) GetArticleCarouselList(c *gin.Context) {
	var position int
	articleCarouselSrv := acc.articleCarouselSrv

	position, err := strconv.Atoi(c.Query("position"))
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 调用服务获取文章列表
	res, err := articleCarouselSrv.GetArticleCarouselList(position)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", res)
	c.JSONP(http.StatusOK, response)
}
