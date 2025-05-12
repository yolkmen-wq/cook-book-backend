package main

import (
	"cook-book-backEnd/middlewares"
	_ "cook-book-backEnd/respositories"
	"cook-book-backEnd/routes"
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
