package dto

import "time"

type CommentResponse struct {
	ID              int64     `json:"id"`
	ArticleId       int64     `json:"articleId"`
	UserId          int64     `json:"userId"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	ParentCommentId int64     `json:"parentCommentId"`
	RootCommentId   int64     `json:"rootCommentId"`
	LikeCount       int       `json:"likeCount"`
	CommentCount    int       `json:"commentCount"`
	IsLiked         bool      `json:"is_liked"` // 是否已点赞
}

type GetCommentRequest struct {
	ArticleId int64 `json:"articleId" form:"articleId" validate:"required"`
	PageNum   int   `json:"pageNum" form:"pageNum" validate:"min=1"`
	PageSize  int   `json:"pageSize" form:"pageSize" validate:"min=1,max=100"`
	UserID    int64 `json:"user_id"` // 当前用户 ID，用于检查点赞
}

type CreateCommentRequest struct {
	ArticleId       int64  `json:"articleId" validate:"required"`
	UserId          int64  `json:"userId" validate:"required"`
	Content         string `json:"content" validate:"required,min=1"`
	ParentCommentId int64  `json:"parentCommentId"`
	RootCommentId   int64  `json:"rootCommentId"`
}

type CommentLike struct {
	ID        int64     `json:"id"`
	CommentID int64     `json:"commentId"`
	UserID    int64     `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}
