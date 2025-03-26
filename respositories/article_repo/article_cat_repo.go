package article_repo

import (
	"cook-book-admin-backend/models"
	"gorm.io/gorm"
)

type ArticleCatRepository struct {
	db *gorm.DB
}

func NewArticleCatRepository(db *gorm.DB) *ArticleCatRepository {
	return &ArticleCatRepository{db: db}
}

// 新增文章分类
func (r *ArticleCatRepository) CreateArticleCat(articleCat *models.ArticleCategory) error {
	return r.db.Create(articleCat).Error
}

//// 根据ID获取文章分类
//func (r *ArticleCatRepository) GetArticleCatById(id uint) (*models.ArticleCategory, error) {
//	articleCat := &models.ArticleCategory{}
//	err := r.db.First(articleCat, id).Error
//	if err != nil {
//		return nil, err
//	}
//	return articleCat, nil
//}
//
//// 根据名称获取文章分类
//func (r *ArticleCatRepository) GetArticleCatByName(name string) (*models.ArticleCategory, error) {
//	articleCat := &models.ArticleCategory{}
//	err := r.db.Where("name = ?", name).First(articleCat).Error
//	if err != nil {
//		return nil, err
//	}
//	return articleCat, nil
//}

// 获取所有文章分类
func (r *ArticleCatRepository) GetArticleCats() ([]*models.ArticleCategory, error) {
	var articleCats []*models.ArticleCategory
	err := r.db.Find(&articleCats).Error
	if err != nil {
		return nil, err
	}
	return articleCats, nil
}

// 更新文章分类
func (r *ArticleCatRepository) UpdateArticleCat(articleCat *models.ArticleCategory) error {
	return r.db.Model(&articleCat).Updates(map[string]interface{}{
		"category_name": articleCat.CategoryName,
		"category_desc": articleCat.CategoryDesc,
		"category_pic":  articleCat.CategoryPic,
		"category_sort": articleCat.CategorySort,
		"show_category": articleCat.ShowCategory,
	}).Error
}

// 删除文章分类
func (r *ArticleCatRepository) DeleteArticleCat(id int64) error {
	articleCat := &models.ArticleCategory{}
	err := r.db.Delete(articleCat, id).Error
	if err != nil {
		return err
	}
	return nil
}
