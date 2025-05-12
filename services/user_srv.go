package services

import (
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/respositories"
)

type UserService interface {
	AdminLogin(adminUser models.AdminUser) (*models.AdminUser, error)
	AdminUserLogout(id int64) error
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

// AdminUserLogout function is used to logout admin user
func (us *userService) AdminUserLogout(id int64) error {
	return us.userRepo.AdminUserLogout(id)
}
