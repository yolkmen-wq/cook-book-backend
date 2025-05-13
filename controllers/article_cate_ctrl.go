package controllers

import (
	"cook-book-backend/config"
	"cook-book-backend/models"
	"cook-book-backend/services"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type ArticleCateCtrl struct {
	articleCateSrv services.ArticleCateSrv
}

func NewArticleCateCtrl(articleCateSrv services.ArticleCateSrv) ArticleCateCtrl {
	return ArticleCateCtrl{
		articleCateSrv: articleCateSrv,
	}
}

func (acc ArticleCateCtrl) GetArticleCateList(c *gin.Context) {
	var req models.GetArticleCateRequest
	articleCateSrv := acc.articleCateSrv

	// 绑定请求数据到 req
	if err := c.ShouldBind(&req); err != nil {
		// 如果请求体为空，使用默认值
		if errors.Is(err, io.EOF) || strings.Contains(err.Error(), "EOF") {
			req.PageNum = 1   // 设置默认页码
			req.PageSize = 10 // 设置默认每页数量
			// 可以根据需要为其他字段设置默认值
		} else {
			if req.PageNum == 0 {
				req.PageNum = 1
			}

			if req.PageSize == 0 {
				req.PageSize = 10
			}
			// 其他绑定错误，返回错误响应
			response := config.NewResponse(http.StatusBadRequest, false, err.Error(), nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	// 调用服务获取文章列表
	res := articleCateSrv.GetArticleCateList(&req)
	if res.Error != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, res.Error.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", config.ListResponse{
		List:     res.List,
		Total:    res.Total,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	c.JSONP(http.StatusOK, response)
}
