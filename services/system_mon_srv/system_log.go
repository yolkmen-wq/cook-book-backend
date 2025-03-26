package system_mon_srv

import (
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/respositories/system_mon_repo"
)

type SystemLogService interface {
	AddSystemLog(log models.SystemLog) error
	GetSystemLogs(page, pageSize int) ([]models.SystemLog, int64, error)
	DeleteSystemLogs() error
}
type systemLogService struct {
	repo system_mon_repo.SystemLogRepository
}

func NewSystemLogService(repo *system_mon_repo.SystemLogRepository) SystemLogService {
	return &systemLogService{
		repo: *repo,
	}
}

// AddSystemLog add system log
func (s *systemLogService) AddSystemLog(log models.SystemLog) error {
	return s.repo.AddSystemLog(log)
}

// DeleteSystemLogs delete system logs
func (s *systemLogService) DeleteSystemLogs() error {
	return s.repo.DeleteSystemLogs()
}

// GetSystemLogs get system logs
func (s *systemLogService) GetSystemLogs(page, pageSize int) ([]models.SystemLog, int64, error) {
	return s.repo.GetSystemLogs(page, pageSize)
}
