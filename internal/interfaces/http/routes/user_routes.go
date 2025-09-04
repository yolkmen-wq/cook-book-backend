package routes

import (
	// "cook-book-backend/config"
	// "cook-book-backend/controllers"
	// "cook-book-backend/internal/domain/repositories"
	// "cook-book-backend/internal/domain/services"
	// "fmt"

	"cook-book-backend/internal/container"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, rg *gin.RouterGroup, container *container.Container) {
	// userController := controllers.NewUserController(services.NewUserService(repositories.NewUserRepository(config.DB)))
	// fmt.Println("userController", userController)
	//r.POST("/admin/login", userController.AdminLogin)
	//r.POST("/admin/refresh-token", userController.AdminRefreshToken)
	//rg.GET("/admin/get-async-routes", userController.GetAsyncRoutes)
	//rg.POST("/admin/logout", userController.AdminLogout)
	r.POST("/app/wx-login", container.UserHandler.WechatLogin)
	r.GET("/app/getPageControl", container.UserHandler.GetPageControl)
	r.POST("/app/setPageControl", container.UserHandler.SetPageControl)
}
