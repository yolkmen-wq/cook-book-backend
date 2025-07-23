package handlers

import (
	"cook-book-backend/config"
	"cook-book-backend/internal/domain/services"
	"cook-book-backend/internal/infrastructure/logger"
	"cook-book-backend/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmojiHandler struct {
	*BaseHandler
	emojiSrv services.EmojiService
}

func NewEmojiCtrl(emojiSrv services.EmojiService, logger logger.Logger) EmojiHandler {
	return EmojiHandler{
		BaseHandler: NewBaseHandler(logger),
		emojiSrv:    emojiSrv,
	}
}

func (h EmojiHandler) GetAllEmoji(c *gin.Context) {
	limit := h.GetIntQueryParam(c, "limit", 10)
	emojiSrv := h.emojiSrv

	// 调用服务获取Emoji列表
	list, total, err := emojiSrv.GetAllEmoji()
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 构建分页信息
	pagination := &response.Pagination{
		Page:      1,
		PageSize:  limit,
		Total:     total,
		TotalPage: int((total + int64(limit) - 1) / int64(limit)),
	}

	h.SuccessWithPagination(c, list, pagination)
}
