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

// GetCommentList 获取评论列表，并包含是否点赞信息
func (cr *CommentRepository) GetCommentList(req *dto.GetCommentRequest) ([]dto.CommentResponse, int64, error) {
	var list []dto.CommentResponse
	var total int64

	// 先获取总数
	err := cr.db.Table("comments").
		Where("article_id = ?", req.ArticleId).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询评论列表，并使用子查询添加 is_liked 字段
	err = cr.db.Table("comments c").
		Select("c.*, (CASE WHEN EXISTS (SELECT 1 FROM comment_likes cl WHERE cl.comment_id = c.id AND cl.user_id = ?) THEN true ELSE false END) AS is_liked", req.UserID).
		Where("c.article_id = ?", req.ArticleId).
		Order("c.created_at asc").
		Limit(req.PageSize).
		Offset((req.PageNum - 1) * req.PageSize).
		Scan(&list).Error

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
	err := cr.db.Table("comments").Where("id = ?", commentId).Update("like_count", gorm.Expr("like_count + 1")).Error
	// 存入点赞记录
	like := &dto.CommentLike{
		CommentID: commentId,
	}
	like.UserID = 1 // 假设用户ID为1
	if err != nil {
		return err
	}
	return cr.db.Table("comment_likes").Create(like).Error
}

// GetLikeCount 获取点赞数
func (cr *CommentRepository) GetLikeCount(commentId int64) (int64, error) {
	var count int64
	err := cr.db.Table("comments").Where("comment_id = ?", commentId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// DeleteLike 取消点赞
func (cr *CommentRepository) DeleteLike(commentId int64) error {
	err := cr.db.Table("comments").Where("id = ?", commentId).Update("like_count", gorm.Expr("like_count - 1")).Error
	if err != nil {
		return err
	}
	return cr.db.Table("comment_likes").Where("comment_id = ? AND user_id = ?", commentId, 1).Delete(&dto.CommentLike{}).Error
}

// CreateCommentCount 创建评论数
func (cr *CommentRepository) CreateCommentCount(commentId int64) error {
	return cr.db.Model(&dto.CommentResponse{}).Where("id = ?", commentId).Update("comment_count", gorm.Expr("comment_count + 1")).Error
}
