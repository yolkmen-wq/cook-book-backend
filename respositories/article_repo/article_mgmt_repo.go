package article_repo

import (
	"cook-book-admin-backend/models"
	"fmt"
	"gorm.io/gorm"
)

type ArticleMgmtRepository struct {
	db *gorm.DB
}

func NewArticleMgmtRepository(db *gorm.DB) *ArticleMgmtRepository {
	return &ArticleMgmtRepository{db: db}
}

// GetArticleList 获取文章列表
func (r *ArticleMgmtRepository) GetArticleList(req *models.GetArticlesRequest) ([]models.Article, int64, int, int, error) {
	// TODO: implement this method
	var articles []models.Article
	var total int64
	fmt.Println("CreatedTime", req.CreatedTime)

	// 构建查询条件
	db := r.db.Table("articles")

	if req.Title != "" {
		db = db.Where("title LIKE ?", "%"+req.Title+"%")
	}

	if req.Author != "" {
		db = db.Where("author LIKE ?", "%"+req.Author+"%")
	}

	if req.Status != nil {
		db = db.Where("status = ?", req.Status)
	}

	if len(req.CreatedTime) == 2 {
		fmt.Println(11111111111)
		fmt.Println(req.CreatedTime[0], req.CreatedTime[1])
		db = db.Where("created_at BETWEEN ? AND ?", req.CreatedTime[0], req.CreatedTime[1])
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		fmt.Println("获取文章列表总数失败", err)
		return nil, 0, 0, 0, err
	}

	// 分页
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	db = db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize)

	err := db.Find(&articles).Error
	if err != nil {
		return nil, 0, 0, 0, err
	}
	return articles, total, req.PageSize, req.PageNum, nil
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

//// UploadArticleImage 上传文章图片
//func (r *ArticleMgmtRepository) UploadArticleImage(image models.ArticleImage) error {
//	// TODO: implement this method
//}
