package routes

import (
	"cook-book-backEnd/controllers"
	"cook-book-backEnd/respositories"
	"cook-book-backEnd/services"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	userController := controllers.NewUserController(services.NewUserService(respositories.NewUserRepository(respositories.DB)))

	r.POST("/admin/login", userController.AdminLogin)
	r.POST("/admin/refresh-token", userController.AdminRefreshToken)
	rg.GET("/admin/get-async-routes", userController.GetAsyncRoutes)
	rg.POST("/admin/logout", userController.AdminLogout)
}
