package respositories

import (
	"cook-book-backend/models"
	"gorm.io/gorm"
)

type ArticleCateRepo struct {
	db *gorm.DB
}

func NewArticleCateRepo(db *gorm.DB) ArticleCateRepo {
	return ArticleCateRepo{db: db}
}

// 获取文章分类列表
func (acr ArticleCateRepo) GetArticleCateList(req *models.GetArticleCateRequest) models.GetArticleCateResponse {
	var response models.GetArticleCateResponse

	response.PageNum = req.PageNum
	response.PageSize = req.PageSize

	err := acr.db.Table("article_categories").
		Where("category_name like?", "%"+req.Name+"%").
		Count(&response.Total).
		Order("created_at desc").
		Limit(response.PageSize).
		Offset((response.PageNum - 1) * response.PageSize).
		Find(&response.List).Error

	if err != nil {
		response.Error = err
		return response
	}
	return response
}
