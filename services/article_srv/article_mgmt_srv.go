package article_srv

import (
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/respositories/article_repo"
)

type ArticleMgmtService interface {
	GetArticleList(req *models.GetArticlesRequest) ([]models.Article, int64, int, int, error)
	CreateArticle(article models.Article) error
	//GetArticle(ctx context.Context, id string) (*Article, error)
	UpdateArticle(article models.Article) error
	DeleteArticle(id int64) error
}

type articleMgmtService struct {
	repo article_repo.ArticleMgmtRepository
}

func NewArticleMgmtService(repo *article_repo.ArticleMgmtRepository) ArticleMgmtService {
	return &articleMgmtService{
		repo: *repo,
	}
}

// GetArticleList returns a list of articles
func (s *articleMgmtService) GetArticleList(req *models.GetArticlesRequest) ([]models.Article, int64, int, int, error) {
	return s.repo.GetArticleList(req)
}

// CreateArticle creates a new article
func (s *articleMgmtService) CreateArticle(article models.Article) error {
	return s.repo.CreateArticle(article)
}

// UpdateArticle updates an existing article
func (s *articleMgmtService) UpdateArticle(article models.Article) error {
	return s.repo.UpdateArticle(article)
}

// DeleteArticle deletes an article by id
func (s *articleMgmtService) DeleteArticle(id int64) error {
	return s.repo.DeleteArticle(id)
}
