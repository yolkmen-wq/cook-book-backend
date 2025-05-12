package article_srv

import (
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/respositories/article_repo"
)

type ArticleCatService interface {
	CreateArticleCat(articleCat *models.ArticleCategory) error
	GetArticleCats(req *models.GetArticleCatsRequest) ([]*models.ArticleCategory, int64, int, int, error)
	UpdateArticleCat(articleCat *models.ArticleCategory) error
	DeleteArticleCat(id int64) error
}
type articleCatService struct {
	repo article_repo.ArticleCatRepository
}

func NewArticleCatService(repo *article_repo.ArticleCatRepository) ArticleCatService {
	return &articleCatService{
		repo: *repo,
	}
}

// CreateArticleCat creates a new article category
func (acs *articleCatService) CreateArticleCat(articleCat *models.ArticleCategory) error {
	return acs.repo.CreateArticleCat(articleCat)
}

// GetArticleCats returns all article categories
func (acs *articleCatService) GetArticleCats(req *models.GetArticleCatsRequest) ([]*models.ArticleCategory, int64, int, int, error) {
	return acs.repo.GetArticleCats(req)
}

// UpdateArticleCat updates an existing article category
func (acs *articleCatService) UpdateArticleCat(articleCat *models.ArticleCategory) error {
	return acs.repo.UpdateArticleCat(articleCat)
}

// DeleteArticleCat deletes an existing article category
func (acs *articleCatService) DeleteArticleCat(id int64) error {
	return acs.repo.DeleteArticleCat(id)
}
