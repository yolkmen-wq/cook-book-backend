package repositories

import (
	"cook-book-backend/internal/interfaces/dto"

	"gorm.io/gorm"
)

type ArticleCateRepository struct {
	db *gorm.DB
}

func NewArticleCateRepository(db *gorm.DB) ArticleCateRepository {
	return ArticleCateRepository{db: db}
}

// GetArticleCateList 获取文章分类列表
func (acr ArticleCateRepository) GetArticleCateList(req *dto.GetArticleCateRequest) ([]dto.ArticleCateResponse, int64, error) {
	var list []dto.ArticleCateResponse
	var total int64

	err := acr.db.Table("article_categories").
		Where("category_name like?", "%"+req.Name+"%").
		Count(&total).
		Order("created_at desc").
		Limit(req.PageSize).
		Offset((req.PageNum - 1) * req.PageSize).
		Find(&list).Error

	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
