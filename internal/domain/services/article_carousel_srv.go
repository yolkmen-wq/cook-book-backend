package services

import (
	"cook-book-backend/internal/domain/repositories"
	"cook-book-backend/internal/interfaces/dto"
)

type ArticleCarouselService interface {
	GetArticleCarouselList(position int) ([]dto.ArticleCarouselItemResponse, error)
}

type articleCarouselService struct {
	articleCarouselRepo repositories.ArticleCarouselRepository
}

func NewArticleCarouselService(articleCarouselRepo repositories.ArticleCarouselRepository) ArticleCarouselService {
	return articleCarouselService{
		articleCarouselRepo: articleCarouselRepo,
	}
}

// GetArticleCarouselList returns a list of article carousel items
func (srv articleCarouselService) GetArticleCarouselList(position int) ([]dto.ArticleCarouselItemResponse, error) {
	return srv.articleCarouselRepo.GetArticleCarouselList(position)
}
