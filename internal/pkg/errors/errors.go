package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

// 预定义错误
var (
	ErrInternalServer = &AppError{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}

	ErrBadRequest = &AppError{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}

	ErrUnauthorized = &AppError{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}

	ErrNotFound = &AppError{
		Code:    http.StatusNotFound,
		Message: "Resource not found",
	}
)

func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func Wrap(err error, code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: err.Error(),
	}
}
