package system_mon_repo

import (
	"cook-book-admin-backend/models"
	"fmt"
	"gorm.io/gorm"
)

type OnlineUserRepository struct {
	db *gorm.DB
}

func NewOnlineUserRepository(db *gorm.DB) *OnlineUserRepository {
	return &OnlineUserRepository{db: db}
}

func (r *OnlineUserRepository) GetOnlineUsers(req models.GetUsersRequest) ([]models.AdminUserMgmt, int64, int, int, error) {
	var onlineUsers []models.AdminUserMgmt
	var total int64

	// 构建查询条件
	db := r.db.Table("admin_users").Where("login_status =1")

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		fmt.Println("获取在线用户总数失败", err)
		return nil, 0, 0, 0, err
	}

	// 获取分页在线用户列表
	if req.PageNum != 0 && req.PageSize != 0 {
		if err := db.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&onlineUsers).Error; err != nil {
			fmt.Println("获取在线用户列表失败", err)
			return nil, 0, 0, 0, err
		}
	} else {
		if err := db.Find(&onlineUsers).Error; err != nil {
			fmt.Println("获取在线用户列表失败", err)
			return nil, 0, 0, 0, err
		}
	}

	return onlineUsers, total, req.PageNum, req.PageSize, nil
}
