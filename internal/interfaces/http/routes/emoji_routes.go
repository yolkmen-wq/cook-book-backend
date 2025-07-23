package routes

import (
	"cook-book-backend/internal/container"

	"github.com/gin-gonic/gin"
)

func SetupEmojiRoutes(r *gin.Engine, rg *gin.RouterGroup, container *container.Container) {
	r.POST("/app/emoji/list", container.EmojiHandler.GetAllEmoji)
}
