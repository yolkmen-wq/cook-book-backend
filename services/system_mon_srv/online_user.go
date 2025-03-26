package system_mon_srv

import (
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/respositories/system_mon_repo"
)

type OnlineUserService interface {
	GetOnlineUsers(req models.GetUsersRequest) ([]models.AdminUserMgmt, int64, int, int, error)
}

type onlineUserService struct {
	repo system_mon_repo.OnlineUserRepository
}

func NewOnlineUserService(repo *system_mon_repo.OnlineUserRepository) OnlineUserService {
	return &onlineUserService{
		repo: *repo,
	}
}

func (s *onlineUserService) GetOnlineUsers(req models.GetUsersRequest) ([]models.AdminUserMgmt, int64, int, int, error) {
	return s.repo.GetOnlineUsers(req)
}
