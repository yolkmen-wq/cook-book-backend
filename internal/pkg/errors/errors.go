package errors

import (
    "fmt"
    "net/http"
)

type AppError struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Details string `json:"details,omitempty"`
    Status  int    `json:"-"`
}

func (e *AppError) Error() string {
    return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
}

// 预定义错误
var (
    ErrInternalServer = &AppError{
        Code:    "INTERNAL_SERVER_ERROR",
        Message: "Internal server error",
        Status:  http.StatusInternalServerError,
    }
    
    ErrBadRequest = &AppError{
        Code:    "BAD_REQUEST",
        Message: "Bad request",
        Status:  http.StatusBadRequest,
    }
    
    ErrUnauthorized = &AppError{
        Code:    "UNAUTHORIZED",
        Message: "Unauthorized",
        Status:  http.StatusUnauthorized,
    }
    
    ErrNotFound = &AppError{
        Code:    "NOT_FOUND",
        Message: "Resource not found",
        Status:  http.StatusNotFound,
    }
)

func New(code, message string, status int) *AppError {
    return &AppError{
        Code:    code,
        Message: message,
        Status:  status,
    }
}

func Wrap(err error, code, message string, status int) *AppError {
    return &AppError{
        Code:    code,
        Message: message,
        Details: err.Error(),
        Status:  status,
    }
}