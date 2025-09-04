package routes

import (
	"cook-book-backend/internal/container"

	"github.com/gin-gonic/gin"
)

func SetupAiRoutes(r *gin.Engine, rg *gin.RouterGroup, container *container.Container) {
	r.POST("/chat", container.ChatHandler.Chat)
	r.POST("/chat-stream", container.ChatHandler.ChatStream)
	r.GET("/chat-ws", container.ChatHandler.ChatWebSocket) // WebSocket端点
}
