package services

import (
	"cook-book-backEnd/models"
	"cook-book-backEnd/respositories"
)

type UserService interface {
	AdminLogin(adminUser models.AdminUser) (*models.AdminUser, error)
	UserLogin(username string, password string) (*models.User, error)
	GetAsyncRoutes(userId int64) ([]models.Router, error)
}

type userService struct {
	userRepo respositories.UserRepository
}

func NewUserService(userRepo *respositories.UserRepository) UserService {
	return &userService{userRepo: *userRepo}
}

// AdminLogin function is used to login admin user
func (us *userService) AdminLogin(adminUser models.AdminUser) (*models.AdminUser, error) {
	return us.userRepo.FindAdminUser(adminUser)
}

// UserLogin function is used to login user
func (us *userService) UserLogin(username string, password string) (*models.User, error) {
	return us.userRepo.FindUser(username, password)
}

// GetAsyncRoutes function is used to get async routes for user
func (us *userService) GetAsyncRoutes(userId int64) ([]models.Router, error) {
	return us.userRepo.GetRoutes(userId)
}
