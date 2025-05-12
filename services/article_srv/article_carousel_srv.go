package article_srv

import (
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/respositories/article_repo"
)

type CarouselService interface {
	GetCarousels(req *models.GetCarouselsRequest) ([]models.Carousel, int64, int, int, error)
	CreateCarousel(carousel *models.Carousel) error
	UpdateCarousel(carousel *models.Carousel) error
	DeleteCarousel(id int64) error
	GetCarouselItems(req *models.GetCarouselItemsRequest) ([]models.CarouselItem, int64, int, int, error)
	CreateCarouselItem(carouselItem *models.CarouselItem) error
	UpdateCarouselItem(carouselItem *models.CarouselItem) error
	DeleteCarouselItem(id int64) error
}

type carouselService struct {
	repo *article_repo.CarouselRepository
}

func NewCarouselService(repo *article_repo.CarouselRepository) CarouselService {
	return &carouselService{
		repo: repo,
	}
}

// GetCarousels get carousels
func (cs *carouselService) GetCarousels(req *models.GetCarouselsRequest) ([]models.Carousel, int64, int, int, error) {
	return cs.repo.GetCarouses(req)
}

// CreateCarousel create carousel
func (cs *carouselService) CreateCarousel(carousel *models.Carousel) error {
	return cs.repo.CreateCarousel(carousel)
}

// UpdateCarousel update carousel
func (cs *carouselService) UpdateCarousel(carousel *models.Carousel) error {
	return cs.repo.UpdateCarousel(carousel)
}

// DeleteCarousel delete carousel
func (cs *carouselService) DeleteCarousel(id int64) error {
	return cs.repo.DeleteCarousel(id)
}

// GetCarouselItems get carousel items
func (cs *carouselService) GetCarouselItems(req *models.GetCarouselItemsRequest) ([]models.CarouselItem, int64, int, int, error) {
	return cs.repo.GetCarouselItems(req)
}

// CreateCarouselItem create carousel item
func (cs *carouselService) CreateCarouselItem(carouselItem *models.CarouselItem) error {
	return cs.repo.CreateCarouselItem(carouselItem)
}

// UpdateCarouselItem update carousel item
func (cs *carouselService) UpdateCarouselItem(carouselItem *models.CarouselItem) error {
	return cs.repo.UpdateCarouselItem(carouselItem)
}

// DeleteCarouselItem delete carousel item
func (cs *carouselService) DeleteCarouselItem(id int64) error {
	return cs.repo.DeleteCarouselItem(id)
}
