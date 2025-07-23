package response

import (
    "net/http"
    
    "cook-book-backend/internal/pkg/errors"
    "github.com/gin-gonic/gin"
)

type Response struct {
    Success bool        `json:"success"`
    Code    string      `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

type PaginatedResponse struct {
    Response
    Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
    Page      int   `json:"page"`
    PageSize  int   `json:"pageSize"`
    Total     int64 `json:"total"`
    TotalPage int   `json:"totalPage"`
}

func Success(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, Response{
        Success: true,
        Code:    "SUCCESS",
        Message: "Success",
        Data:    data,
    })
}

func SuccessWithPagination(c *gin.Context, data interface{}, pagination *Pagination) {
    c.JSON(http.StatusOK, PaginatedResponse{
        Response: Response{
            Success: true,
            Code:    "SUCCESS",
            Message: "Success",
            Data:    data,
        },
        Pagination: pagination,
    })
}

func Error(c *gin.Context, err *errors.AppError) {
    c.JSON(err.Status, Response{
        Success: false,
        Code:    err.Code,
        Message: err.Message,
    })
}