package repositories

import (
	"cook-book-backend/internal/interfaces/dto"

	"fmt"

	"gorm.io/gorm"
)

type ArticleCarouselRepository struct {
	db *gorm.DB
}

func NewArticleCarouselRepo(db *gorm.DB) ArticleCarouselRepository {
	return ArticleCarouselRepository{db: db}
}

// GetArticleCarouselList 获取文章轮播图列表
func (repo *ArticleCarouselRepository) GetArticleCarouselList(position int) ([]dto.ArticleCarouselItemResponse, error) {
	var articleCarouselItems []dto.ArticleCarouselItemResponse
	fmt.Println(19, position)
	err := repo.db.Table("carousel_items").Joins("LEFT JOIN carousel ON carousel.carousel_id = carousel_items.carousel_id").Where("carousel.position =?", position).Find(&articleCarouselItems).Error
	return articleCarouselItems, err
}
