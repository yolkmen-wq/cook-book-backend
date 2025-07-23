package handlers

import (
	"cook-book-backend/config"
	"cook-book-backend/internal/domain/services"
	"cook-book-backend/internal/infrastructure/logger"
	"cook-book-backend/internal/interfaces/dto"
	"cook-book-backend/internal/pkg/response"

	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ArticleCateHandler struct {
	*BaseHandler
	articleCateSrv services.ArticleCateService
}

func NewArticleCateCtrl(articleCateSrv services.ArticleCateService, logger logger.Logger) ArticleCateHandler {
	return ArticleCateHandler{
		BaseHandler:    NewBaseHandler(logger),
		articleCateSrv: articleCateSrv,
	}
}

func (h ArticleCateHandler) GetArticleCateList(c *gin.Context) {
	var req dto.GetArticleCateRequest
	articleCateSrv := h.articleCateSrv

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
	list, total, err := articleCateSrv.GetArticleCateList(&req)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 构建分页信息
	pagination := &response.Pagination{
		Page:      req.PageNum,
		PageSize:  req.PageSize,
		Total:     total,
		TotalPage: int((total + int64(req.PageSize) - 1) / int64(req.PageSize)),
	}

	h.SuccessWithPagination(c, list, pagination)
}
