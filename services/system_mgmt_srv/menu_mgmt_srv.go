package system_mgmt_srv

import (
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/respositories/system_mgmt_repo"
)

type MenuMgmtService interface {
	GetAsyncRoutes() ([]models.Router, error)
	GetMenuList(searchInfo models.GetMenuRequest) ([]models.Menu, int64, error)
	GetMenuDetail(id int64) (models.Menu, error)
	CreateMenu(menu models.Menu) error
	UpdateMenu(menu models.Menu) error
	DeleteMenu(id int64) error
}

type menuMgmtService struct {
	menuMgmtRepo system_mgmt_repo.MenuMgmtRepository
}

func NewMenuMgmtService(menuMgmtRepo *system_mgmt_repo.MenuMgmtRepository) MenuMgmtService {
	return &menuMgmtService{
		menuMgmtRepo: *menuMgmtRepo,
	}
}

// GetAsyncRoutes function is used to get async routes for user
func (us *menuMgmtService) GetAsyncRoutes() ([]models.Router, error) {
	return us.menuMgmtRepo.GetRoutesWithLock()
}

// GetMenu returns the menu of the restaurant
func (s *menuMgmtService) GetMenuList(searchInfo models.GetMenuRequest) ([]models.Menu, int64, error) {
	return s.menuMgmtRepo.GetMenuList(searchInfo)
}

// GetMenuDetail returns the detail of the menu
func (s *menuMgmtService) GetMenuDetail(id int64) (models.Menu, error) {
	return s.menuMgmtRepo.GetMenuDetail(id)
}

// CreateMenu creates a new menu
func (s *menuMgmtService) CreateMenu(menu models.Menu) error {
	return s.menuMgmtRepo.CreateMenu(menu)
}

// UpdateMenu updates the menu
func (s *menuMgmtService) UpdateMenu(menu models.Menu) error {
	return s.menuMgmtRepo.UpdateMenu(menu)
}

// DeleteMenu deletes the menu
func (s *menuMgmtService) DeleteMenu(id int64) error {
	return s.menuMgmtRepo.DeleteMenu(id)
}
