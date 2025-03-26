package system_mgmt_repo

import (
	"cook-book-admin-backend/models"
	"fmt"
	"gorm.io/gorm"
)

type UserMgmtRepository struct {
	db *gorm.DB
}

func NewUserMgmtRepository(db *gorm.DB) *UserMgmtRepository {
	return &UserMgmtRepository{db: db}
}

// 获取用户列表
func (lr *UserMgmtRepository) GetAdminUserList(req models.GetUsersRequest) ([]models.AdminUserMgmt, int64, int, int, error) {
	var users []models.AdminUserMgmt
	var total int64

	// 构建查询条件
	db := lr.db.Table("admin_users")
	if req.Username != "" {
		db = db.Where("username LIKE ?", "%"+req.Username+"%")
	}

	if req.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+req.Nickname+"%")
	}

	if req.Status != nil && req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		fmt.Println("获取用户总数失败", err)
		return nil, 0, 0, 0, err
	}

	// 获取分页用户列表
	if req.PageNum != 0 && req.PageSize != 0 {
		if err := db.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&users).Error; err != nil {
			fmt.Println("获取用户列表失败", err)
			return nil, 0, 0, 0, err
		}
	} else {
		if err := db.Find(&users).Error; err != nil {
			fmt.Println("获取用户列表失败", err)
			return nil, 0, 0, 0, err
		}
	}

	return users, total, req.PageSize, req.PageNum, nil
}

// 删除用户
func (lr *UserMgmtRepository) DeleteUser(id int64) error {
	var user models.AdminUserMgmt
	var db = lr.db
	// 构建查询条件
	if id == 0 {
		fmt.Println("删除用户失败，用户ID不能为空")
		return fmt.Errorf("删除用户失败，用户ID不能为空")
	}
	db = db.Table("admin_users").Where("user_id = ?", id).Delete(&user)
	if err := db.Error; err != nil {
		fmt.Println("删除用户失败", err)
		return err
	}
	return nil
}

// 新增用户
func (lr *UserMgmtRepository) AddUser(userInfo models.AdminUserMgmt) error {
	var db = lr.db
	// 构建查询条件
	username := userInfo.Username
	if username == "" {
		fmt.Println("新增用户失败，用户名不能为空")
		return fmt.Errorf("新增用户失败，用户名不能为空")
	}
	password := userInfo.Password
	if password == "" {
		fmt.Println("新增用户失败，密码不能为空")
		return fmt.Errorf("新增用户失败，密码不能为空")
	}
	nickname := userInfo.Nickname
	if nickname == "" {
		fmt.Println("新增用户失败，昵称不能为空")
		return fmt.Errorf("新增用户失败，昵称不能为空")
	}
	status := userInfo.Status
	if status == nil {
		fmt.Println("新增用户失败，状态不能为空")
		return fmt.Errorf("新增用户失败，状态不能为空")
	}
	db = db.Table("admin_users").Create(&userInfo)
	if err := db.Error; err != nil {
		fmt.Println("新增用户失败", err)
		return err
	}
	return nil
}

// 修改用户信息
func (lr *UserMgmtRepository) UpdateUser(userInfo models.AdminUser) error {
	var user models.AdminUser
	var db = lr.db
	// 构建查询条件
	id := userInfo.ID
	if id == 0 {
		fmt.Println("修改用户信息失败，用户ID不能为空")
		return fmt.Errorf("修改用户信息失败，用户ID不能为空")
	}
	status := userInfo.Status
	nickname := userInfo.Nickname
	password := userInfo.Password
	db = db.Model(&user).Where("user_id = ?", id)
	if status != nil {
		fmt.Println("修改用户状态", status)
		db = db.Update("status", *status)
	}
	if nickname != "" {
		fmt.Println("修改用户昵称", nickname)
		db = db.Update("nickname", nickname)
	}
	if password != "" {
		fmt.Println("修改用户密码", password)
		db = db.Update("password", password)
	}
	//db.Updates(userInfo)
	if err := db.Error; err != nil {
		fmt.Println("修改用户状态失败", err)
		return err
	}
	return nil
}

// 用户分配角色
func (lr *UserMgmtRepository) UpdateUserRoles(userId int64, roleIds []int64) error {
	var db = lr.db
	// 构建查询条件
	if userId == 0 {
		fmt.Println("分配角色失败，用户ID不能为空")
		return fmt.Errorf("分配角色失败，用户ID不能为空")
	}
	if roleIds == nil {
		fmt.Println("分配角色失败，角色ID不能为空")
		return fmt.Errorf("分配角色失败，角色ID不能为空")
	}
	// 删除原有角色
	db.Table("user_roles").Where("user_id = ?", userId).Delete(&models.UserRole{})
	// 新增角色
	for _, roleId := range roleIds {
		db = db.Table("user_roles").Create(&models.UserRole{
			UserID: userId,
			RoleID: roleId,
		})
	}
	if err := db.Error; err != nil {
		fmt.Println("分配角色失败", err)
		return err
	}
	return nil
}

// 获取角色列表
func (lr *UserMgmtRepository) GetUserRoleList(name string, id int64) ([]models.Role, error) {
	var roles []models.Role
	var db = lr.db

	// 构建查询条件
	if name != "" {
		db = db.Where("role_name = ?", name)
	}
	if id != 0 && &id != nil {
		db = db.Joins("LEFT JOIN user_roles ON user_roles.role_id = roles.role_id").Where("user_roles.user_id = ?", id)
	}
	db = db.Find(&roles)
	if err := db.Error; err != nil {
		fmt.Println("获取角色列表失败", err)
		return nil, err
	}
	return roles, nil

}
