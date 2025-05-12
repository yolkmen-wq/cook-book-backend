package main

import (
	_ "cook-book-admin-backend/config"
	"cook-book-admin-backend/middlewares"
	"cook-book-admin-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// Init routes
	authorized := router.Group("/")
	authorized.Use(middlewares.AuthMiddleWare())
	authorized.Use(middlewares.CommonLogInterceptor)
	routes.InitRoutes(router, authorized)
	router.Run(":5757")
}
