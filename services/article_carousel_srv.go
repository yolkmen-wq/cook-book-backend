package services

import (
	"cook-book-backend/models"
	"cook-book-backend/respositories"
)

type ArticleCarouselSrv interface {
	GetArticleCarouselList(position int) ([]models.ArticleCarouselItem, error)
}

type articleCarouselSrv struct {
	articleCarouselRepo respositories.ArticleCarouselRepo
}

func NewArticleCarouselSrv(articleCarouselRepo respositories.ArticleCarouselRepo) ArticleCarouselSrv {
	return articleCarouselSrv{
		articleCarouselRepo: articleCarouselRepo,
	}
}

// GetArticleCarouselList returns a list of article carousel items
func (srv articleCarouselSrv) GetArticleCarouselList(position int) ([]models.ArticleCarouselItem, error) {
	return srv.articleCarouselRepo.GetArticleCarouselList(position)
}
