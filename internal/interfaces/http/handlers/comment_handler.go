package handlers

import (
	"cook-book-backend/internal/domain/services"
	"cook-book-backend/internal/infrastructure/logger"
	"cook-book-backend/internal/interfaces/dto"
	"cook-book-backend/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	*BaseHandler
	commentService services.CommentService
}

func NewCommentHandler(commentService services.CommentService, logger logger.Logger) *CommentHandler {
	return &CommentHandler{
		BaseHandler:    NewBaseHandler(logger),
		commentService: commentService,
	}
}

// GetCommentList 获取评论列表
func (h *CommentHandler) GetCommentList(c *gin.Context) {
	var req dto.GetCommentRequest

	// 设置默认值
	req.PageNum = h.GetIntQueryParam(c, "pageNum", 1)
	req.PageSize = h.GetIntQueryParam(c, "pageSize", 10)
	req.ArticleId = int64(h.GetIntQueryParam(c, "articleId", 0))

	if err := h.BindAndValidate(c, &req); err != nil {
		h.HandleError(c, err, "bind_comment_list_params")
		return
	}

	comments, total, err := h.commentService.GetCommentList(&req)
	if err != nil {
		h.HandleError(c, err, "get_comment_list")
		return
	}

	pagination := &response.Pagination{
		Page:      req.PageNum,
		PageSize:  req.PageSize,
		Total:     total,
		TotalPage: int((total + int64(req.PageSize) - 1) / int64(req.PageSize)),
	}

	h.SuccessWithPagination(c, comments, pagination)
}

// CreateComment 创建评论
func (h *CommentHandler) CreateComment(c *gin.Context) {
	var req dto.CommentResponse

	if err := h.BindAndValidate(c, &req); err != nil {
		h.HandleError(c, err, "bind_create_comment_params")
		return
	}

	err := h.commentService.CreateComment(&req)
	if err != nil {
		h.HandleError(c, err, "create_comment")
		return
	}

	h.Success(c, gin.H{"message": "创建成功"})
}

// LikeComment 点赞评论
func (h *CommentHandler) LikeComment(c *gin.Context) {
	commentID, err := h.GetIDParam(c, "id")
	if err != nil {
		h.HandleError(c, err, "get_comment_id_param")
		return
	}

	if err := h.commentService.CreateLike(commentID); err != nil {
		h.HandleError(c, err, "like_comment")
		return
	}

	h.Success(c, gin.H{"message": "点赞成功"})
}

// // UnlikeComment 取消点赞
// func (h *CommentHandler) UnlikeComment(c *gin.Context) {
// 	commentID, err := h.GetIDParam(c, "id")
// 	if err != nil {
// 		h.HandleError(c, err, "get_comment_id_param")
// 		return
// 	}

// 	if err := h.commentService.UnlikeComment(commentID); err != nil {
// 		h.HandleError(c, err, "unlike_comment")
// 		return
// 	}

// 	h.Success(c, gin.H{"message": "取消点赞成功"})
// }

// // DeleteComment 删除评论
// func (h *CommentHandler) DeleteComment(c *gin.Context) {
// 	commentID, err := h.GetIDParam(c, "id")
// 	if err != nil {
// 		h.HandleError(c, err, "get_comment_id_param")
// 		return
// 	}

// 	if err := h.commentService.DeleteComment(commentID); err != nil {
// 		h.HandleError(c, err, "delete_comment")
// 		return
// 	}

// 	h.Success(c, gin.H{"message": "删除成功"})
// }
