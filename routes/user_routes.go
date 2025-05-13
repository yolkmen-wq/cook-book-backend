package routes

import (
	"cook-book-backend/config"
	"cook-book-backend/controllers"
	"cook-book-backend/respositories"
	"cook-book-backend/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, rg *gin.RouterGroup) {
	userController := controllers.NewUserController(services.NewUserService(respositories.NewUserRepository(config.DB)))
	fmt.Println("userController", userController)
	//r.POST("/admin/login", userController.AdminLogin)
	//r.POST("/admin/refresh-token", userController.AdminRefreshToken)
	//rg.GET("/admin/get-async-routes", userController.GetAsyncRoutes)
	//rg.POST("/admin/logout", userController.AdminLogout)
}
