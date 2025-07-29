package dto

import "time"

type GetArticleListRequest struct {
	PageNum    int    `json:"pageNum" form:"pageNum" validate:"min=1"`
	PageSize   int    `json:"pageSize" form:"pageSize" validate:"min=1,max=100"`
	CategoryID int    `json:"categoryId" form:"categoryId"`
	Keyword    string `json:"keyword" form:"keyword"`
}

type CreateArticleRequest struct {
	Title      string `json:"title" validate:"required,min=1,max=200"`
	Content    string `json:"content" validate:"required"`
	Summary    string `json:"summary" validate:"max=500"`
	CategoryID int64  `json:"categoryId" validate:"required"`
	Tags       string `json:"tags"`
	CoverImage string `json:"coverImage"`
}

type UpdateArticleRequest struct {
	Title      string `json:"title" validate:"min=1,max=200"`
	Content    string `json:"content"`
	Summary    string `json:"summary" validate:"max=500"`
	CategoryID int64  `json:"categoryId"`
	Tags       string `json:"tags"`
	CoverImage string `json:"coverImage"`
}

type ArticleResponse struct {
	ID         int64     `json:"id" gorm:P`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Summary    string    `json:"summary"`
	CategoryID int64     `json:"categoryId"`
	Tags       string    `json:"tags"`
	CoverImage string    `json:"coverImage" gorm:"column:cover"`
	ViewCount  int       `json:"viewCount"`
	LikeCount  int       `json:"likeCount"`
	CreatedAt  time.Time `json:"createdTime" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updatedTime" gorm:"column:updated_at"`
}
