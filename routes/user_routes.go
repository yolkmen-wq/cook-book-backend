package routes

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/controllers"
	"cook-book-admin-backend/respositories"
	"cook-book-admin-backend/services"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	userController := controllers.NewUserController(services.NewUserService(respositories.NewUserRepository(config.DB)))

	r.POST("/admin/login", userController.AdminLogin)
	r.POST("/admin/refresh-token", userController.AdminRefreshToken)
	rg.POST("/admin/logout", userController.AdminLogout)
}
