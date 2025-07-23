package handlers

import (
	"cook-book-backend/config"
	"cook-book-backend/internal/domain/services"
	"cook-book-backend/internal/infrastructure/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleCarouselHandler struct {
	*BaseHandler
	articleCarouselSrv services.ArticleCarouselService
}

func NewArticleCarouselHandler(articleCarouselSrv services.ArticleCarouselService, logger logger.Logger) *ArticleCarouselHandler {
	return &ArticleCarouselHandler{
		BaseHandler:        NewBaseHandler(logger),
		articleCarouselSrv: articleCarouselSrv,
	}
}

func (acc ArticleCarouselHandler) GetArticleCarouselList(c *gin.Context) {
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
