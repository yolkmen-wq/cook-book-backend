package repositories

import (
	"cook-book-backend/internal/interfaces/dto"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return CommentRepository{db: db}
}

// GetCommentList 获取评论列表
func (cr *CommentRepository) GetCommentList(req *dto.GetCommentRequest) ([]dto.CommentResponse, int64, error) {
	var list []dto.CommentResponse
	var total int64

	err := cr.db.Table("comments").
		Where("article_id = ?", req.ArticleId).
		Count(&total).
		Order("created_at asc").
		Limit(req.PageSize).
		Offset((req.PageNum - 1) * req.PageSize).
		Find(&list).Error

	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// CreateComment 创建评论
func (cr *CommentRepository) CreateComment(comment *dto.CommentResponse) error {
	return cr.db.Create(comment).Error
}

// CreateLike 点赞
func (cr *CommentRepository) CreateLike(commentId int64) error {
	err := cr.db.Model(&dto.CommentResponse{}).Where("id = ?", commentId).Update("like_count", gorm.Expr("like_count + 1")).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateCommentCount 创建评论数
