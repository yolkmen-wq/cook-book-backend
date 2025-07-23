package routes

import (
	"cook-book-backend/internal/container"

	"github.com/gin-gonic/gin"
)

func SetupCommentRoutes(r *gin.Engine, rg *gin.RouterGroup, container *container.Container) {
	r.POST("/app/comment/list", container.CommentHandler.GetCommentList)
	r.POST("/app/comment/create", container.CommentHandler.CreateComment)
	r.POST("/app/commen/like", container.CommentHandler.LikeComment)
}
