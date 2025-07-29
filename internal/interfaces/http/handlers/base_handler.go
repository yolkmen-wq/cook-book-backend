package handlers

import (
	"strconv"

	"cook-book-backend/internal/infrastructure/logger"
	"cook-book-backend/internal/pkg/errors"
	"cook-book-backend/internal/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BaseHandler struct {
	logger    logger.Logger
	validator *validator.Validate
}

func NewBaseHandler(logger logger.Logger) *BaseHandler {
	return &BaseHandler{
		logger:    logger,
		validator: validator.New(),
	}
}

// BindAndValidate 统一的参数绑定和验证
func (h *BaseHandler) BindAndValidate(c *gin.Context, req interface{}) error {
	if err := c.ShouldBind(req); err != nil {
		h.logger.Error("Parameter binding failed", "error", err.Error())
		return errors.Wrap(err, 400, "参数绑定失败")
	}

	if err := h.validator.Struct(req); err != nil {
		h.logger.Error("Parameter validation failed", "error", err.Error())
		return errors.Wrap(err, 400, "参数验证失败")
	}

	return nil
}

// GetIDParam 获取路径参数中的ID
func (h *BaseHandler) GetIDParam(c *gin.Context, paramName string) (int64, error) {
	idStr := c.Param(paramName)
	if idStr == "" {
		return 0, errors.New(400, "缺少ID参数")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.logger.Error("Invalid ID parameter", "param", paramName, "value", idStr)
		return 0, errors.New(400, "无效的ID参数")
	}

	return id, nil
}

// GetQueryParam 获取查询参数
func (h *BaseHandler) GetQueryParam(c *gin.Context, key string, defaultValue string) string {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetIntQueryParam 获取整数类型的查询参数
func (h *BaseHandler) GetIntQueryParam(c *gin.Context, key string, defaultValue int) int {
	valueStr := c.Query(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}

// HandleError 统一的错误处理
func (h *BaseHandler) HandleError(c *gin.Context, err error, operation string) {
	h.logger.Error("Operation failed", "operation", operation, "error", err.Error())

	if appErr, ok := err.(*errors.AppError); ok {
		response.Error(c, appErr)
		return
	}

	// 未知错误，返回内部服务器错误
	response.Error(c, errors.ErrInternalServer)
}

// Success 统一的成功响应
func (h *BaseHandler) Success(c *gin.Context, data interface{}) {
	response.Success(c, data)
}

// SuccessWithPagination 带分页的成功响应
func (h *BaseHandler) SuccessWithPagination(c *gin.Context, data interface{}, pagination *response.Pagination) {
	response.SuccessWithPagination(c, data, pagination)
}
