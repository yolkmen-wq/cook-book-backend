package repositories

import (
	"cook-book-backend/internal/interfaces/dto"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return ArticleRepository{db: db}
}

// GetArticleList 获取文章列表
func (r *ArticleRepository) GetArticleList(req *dto.GetArticleListRequest) ([]dto.ArticleResponse, int64, error) {
	var list []dto.ArticleResponse
	var total int64

	err := r.db.Table("articles").
		Where("title like?", "%"+req.Keyword+"%").
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

// GetLatestArticle 获取最新前十篇文章
func (r *ArticleRepository) GetLatestArticles(limit int) ([]dto.ArticleResponse, int64, error) {
	var list []dto.ArticleResponse
	var total int64

	err := r.db.Table("articles").
		Order("created_at desc").
		Limit(limit).
		Count(&total).
		Find(&list).Error

	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// GetArticleDetail 获取文章详情
func (r *ArticleRepository) GetArticleDetail(id int64) (dto.ArticleResponse, error) {
	var article dto.ArticleResponse

	err := r.db.Table("articles").
		Where("article_id = ?", id).
		Take(&article).Error

	if err != nil {
		return article, err
	}
	return article, nil
}

// CollectArticle 收藏文章
func (r *ArticleRepository) CollectArticle(id int64) error {
	return nil
}
