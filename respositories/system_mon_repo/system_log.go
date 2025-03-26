package system_mon_repo

import (
	"cook-book-admin-backend/models"
	"gorm.io/gorm"
)

type SystemLogRepository struct {
	db *gorm.DB
}

func NewSystemLogRepository(db *gorm.DB) *SystemLogRepository {
	return &SystemLogRepository{db: db}
}

// AddSystemLog 新增系统日志
func (r *SystemLogRepository) AddSystemLog(log models.SystemLog) error {
	return r.db.Create(&log).Error
}

// DeleteSystemLogs 删除系统日志
func (r *SystemLogRepository) DeleteSystemLogs() error {
	err := r.db.Exec("DELETE FROM system_logs").Error
	if err != nil {
		return err
	}
	return nil
}

// GetSystemLogs 获取系统日志列表
func (r *SystemLogRepository) GetSystemLogs(page, pageSize int) ([]models.SystemLog, int64, error) {
	var logs []models.SystemLog
	var count int64

	err := r.db.Find(&logs).Count(&count).Error
	return logs, count, err
}
