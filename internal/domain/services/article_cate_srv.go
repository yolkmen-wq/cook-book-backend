package services

import (
	"cook-book-backend/internal/domain/repositories"
	"cook-book-backend/internal/interfaces/dto"
)

type ArticleCateService interface {
	GetArticleCateList(req *dto.GetArticleCateRequest) ([]dto.ArticleCateResponse, int64, error)
}

type articleCateSrv struct {
	articleCateRepo repositories.ArticleCateRepository
}

func NewArticleCateSrv(articleCateRepo repositories.ArticleCateRepository) ArticleCateService {
	return articleCateSrv{
		articleCateRepo: articleCateRepo,
	}
}

// GetArticleCateList get article cate list
func (acs articleCateSrv) GetArticleCateList(req *dto.GetArticleCateRequest) ([]dto.ArticleCateResponse, int64, error) {
	return acs.articleCateRepo.GetArticleCateList(req)
}
