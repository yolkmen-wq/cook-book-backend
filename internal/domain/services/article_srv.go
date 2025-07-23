package services

import (
	"cook-book-backend/internal/domain/repositories"
	"cook-book-backend/internal/interfaces/dto"
)

type ArticleService interface {
	GetArticleList(req *dto.GetArticleListRequest) ([]dto.ArticleResponse, int64, error)
	GetLatestArticles(limit int) ([]dto.ArticleResponse, int64, error)
	GetArticleDetail(id int64) (dto.ArticleResponse, error)
}

type articleService struct {
	articleRepo repositories.ArticleRepository
}

func NewArticleService(articleRepo repositories.ArticleRepository) ArticleService {
	return &articleService{
		articleRepo: articleRepo,
	}
}

// GetArticleList returns a list of articles based on the given request
func (as *articleService) GetArticleList(req *dto.GetArticleListRequest) ([]dto.ArticleResponse, int64, error) {
	return as.articleRepo.GetArticleList(req)
}

// GetLatestArticle returns the latest article
func (as *articleService) GetLatestArticles(limit int) ([]dto.ArticleResponse, int64, error) {
	return as.articleRepo.GetLatestArticles(limit)
}

// GetArticleDetail returns the detail of an article based on the given id
func (as *articleService) GetArticleDetail(id int64) (dto.ArticleResponse, error) {
	return as.articleRepo.GetArticleDetail(id)
}
