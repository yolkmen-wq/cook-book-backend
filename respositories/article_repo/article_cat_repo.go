package article_repo

import (
	"cook-book-admin-backend/models"
	"fmt"
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

// 获取文章分类
func (r *ArticleCatRepository) GetArticleCats(req *models.GetArticleCatsRequest) ([]*models.ArticleCategory, int64, int, int, error) {
	var articleCats []*models.ArticleCategory
	var total int64

	// 构建查询条件
	db := r.db.Table("article_categories")

	if req.CategoryName != "" {
		db = db.Where("category_name LIKE ?", "%"+req.CategoryName+"%")
	}
	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		fmt.Println("获取轮播图项总数失败", err)
		return nil, 0, 0, 0, err
	}

	// 获取分页轮播图列表
	if req.PageNum != 0 && req.PageSize != 0 {
		if err := db.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&articleCats).Error; err != nil {
			fmt.Println("获取轮播图项列表失败", err)
			return nil, 0, 0, 0, err
		}
	} else {
		if err := db.Find(&articleCats).Error; err != nil {
			fmt.Println("获取轮播图项列表失败", err)
			return nil, 0, 0, 0, err
		}
	}
	return articleCats, total, req.PageSize, req.PageNum, nil
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
