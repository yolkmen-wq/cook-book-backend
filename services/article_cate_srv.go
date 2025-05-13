package services

import (
	"cook-book-backend/models"
	"cook-book-backend/respositories"
)

type ArticleCateSrv interface {
	GetArticleCateList(req *models.GetArticleCateRequest) models.GetArticleCateResponse
}

type articleCateSrv struct {
	articleCateRepo respositories.ArticleCateRepo
}

func NewArticleCateSrv(articleCateRepo respositories.ArticleCateRepo) ArticleCateSrv {
	return articleCateSrv{
		articleCateRepo: articleCateRepo,
	}
}

// GetArticleCateList get article cate list
func (acs articleCateSrv) GetArticleCateList(req *models.GetArticleCateRequest) models.GetArticleCateResponse {
	return acs.articleCateRepo.GetArticleCateList(req)
}
