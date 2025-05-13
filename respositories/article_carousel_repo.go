package respositories

import (
	"cook-book-backend/models"
	"fmt"
	"gorm.io/gorm"
)

type ArticleCarouselRepo struct {
	db *gorm.DB
}

func NewArticleCarouselRepo(db *gorm.DB) ArticleCarouselRepo {
	return ArticleCarouselRepo{db: db}
}

// 获取文章轮播图列表
func (repo *ArticleCarouselRepo) GetArticleCarouselList(position int) ([]models.ArticleCarouselItem, error) {
	var articleCarouselItems []models.ArticleCarouselItem
	fmt.Println(19, position)
	err := repo.db.Table("carousel_items").Joins("LEFT JOIN carousel ON carousel.carousel_id = carousel_items.carousel_id").Where("carousel.position =?", position).Find(&articleCarouselItems).Error
	return articleCarouselItems, err
}
