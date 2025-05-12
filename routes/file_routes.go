package routes

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/controllers"
	"cook-book-admin-backend/respositories"
	"cook-book-admin-backend/services"
	"github.com/gin-gonic/gin"
)

func SetupFileRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	fileController := controllers.NewFileController(services.NewFileService(respositories.NewFileRepository(config.DB)))

	r.POST("/admin/file/upload", fileController.UploadFile)
}
