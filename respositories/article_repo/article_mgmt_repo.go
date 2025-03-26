package article_repo

import (
	"cook-book-admin-backend/models"
	"gorm.io/gorm"
)

type ArticleMgmtRepository struct {
	db *gorm.DB
}

func NewArticleMgmtRepository(db *gorm.DB) *ArticleMgmtRepository {
	return &ArticleMgmtRepository{db: db}
}

// GetArticleList 获取文章列表
func (r *ArticleMgmtRepository) GetArticleList() ([]models.Article, error) {
	// TODO: implement this method
	var articles []models.Article

	err := r.db.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

// CreateArticle 创建文章
func (r *ArticleMgmtRepository) CreateArticle(article models.Article) error {
	// TODO: implement this method
	//article.CreatedTime = time.Now()
	err := r.db.Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateArticle 更新文章
func (r *ArticleMgmtRepository) UpdateArticle(article models.Article) error {
	// TODO: implement this method
	//article.UpdatedTime = time.Now()
	err := r.db.Updates(&article).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteArticle 删除文章
func (r *ArticleMgmtRepository) DeleteArticle(id int64) error {
	// TODO: implement this method
	err := r.db.Delete(&models.Article{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
