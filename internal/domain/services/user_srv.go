package services

import (
	"cook-book-backend/internal/domain/repositories"
)

type UserService interface {
	//AdminLogin(adminUser models.AdminUser) (*models.AdminUser, error)
	//UserLogin(username string, password string) (*models.User, error)
	//GetAsyncRoutes(userId int64) ([]models.Router, error)
	WechatLogin(code string) (string, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (us *userService) WechatLogin(code string) (string, error) {
	return us.userRepo.WechatLogin(code)
}

//// AdminLogin function is used to login admin user
//func (us *userService) AdminLogin(adminUser models.AdminUser) (*models.AdminUser, error) {
//	return us.userRepo.FindAdminUser(adminUser)
//}
//
//// UserLogin function is used to login user
//func (us *userService) UserLogin(username string, password string) (*models.User, error) {
//	return us.userRepo.FindUser(username, password)
//}
//
//// GetAsyncRoutes function is used to get async routes for user
//func (us *userService) GetAsyncRoutes(userId int64) ([]models.Router, error) {
//	return us.userRepo.GetRoutes(userId)
//}
