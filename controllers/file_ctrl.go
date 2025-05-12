package controllers

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type FileController struct {
	fileService services.FileService
}

func NewFileController(fileService services.FileService) *FileController {
	return &FileController{fileService}
}

func (fc *FileController) UploadFile(c *gin.Context) {
	form, _ := c.MultipartForm()
	name := form.Value["key"][0]
	file := form.File["file"][0]
	src, err := file.Open()
	if err != nil {
		errResp := config.NewResponse(http.StatusInternalServerError, false, "打开文件失败", nil)
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			fmt.Println("close file error", err)
		}
	}(src)
	url, err := fc.fileService.UploadFile(name, src)
	if err != nil {
		errResp := config.NewResponse(http.StatusInternalServerError, false, "上传文件失败", nil)
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "上传成功", map[string]interface{}{
		"url": url,
	})
	c.JSONP(http.StatusOK, response)
}
