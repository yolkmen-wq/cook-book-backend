package respositories

import (
	"cook-book-admin-backend/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

type Result struct {
	RoleName string `json:"role_name"`
}

// 后台登录
func (lr *UserRepository) FindAdminUser(adminUser models.AdminUser) (*models.AdminUser, error) {
	var admin models.AdminUser
	var userRoles []models.UserRole
	var roles []models.Role
	var rolePermissions []models.RolePermission

	// 查询 admin 用户
	err := lr.db.Where("username = ? AND password = ? AND status = ?", adminUser.Username, adminUser.Password, 1).Take(&admin).Error
	if err != nil {
		return nil, err
	}

	// 更新用户登录信息
	adminUser.LoginTime = time.Now()
	adminUser.LoginStatus = 1
	err = lr.db.Where("user_id = ?", admin.ID).Updates(adminUser).Error
	if err != nil {
		fmt.Println("update admin user last login time failed", err)
		return nil, err
	}

	// 查询 admin 用户的角色
	err = lr.db.Find(&userRoles, admin.ID).Error
	if err != nil {
		fmt.Println("admin user role not found", err)
		return nil, err
	}

	// 查询角色
	for _, ur := range userRoles {
		var role models.Role
		err = lr.db.First(&role, ur.RoleID).Error

		// 查询角色权限
		err = lr.db.Find(&rolePermissions, role.ID).Error
		for _, rp := range rolePermissions {
			var permission models.Permission
			err = lr.db.First(&permission, rp.PermissionID).Error
			if err != nil {
				fmt.Println("permission not found", err)
				return nil, err
			}
			admin.Permissions = append(admin.Permissions, permission.Name)
		}

		if err != nil {
			fmt.Println("role not found", err)
			return nil, err
		}
		roles = append(roles, role)
	}

	// 将 Role 记录转换为 []string 并赋值给 admin.Roles
	admin.Roles = make([]string, len(userRoles))
	for i, role := range roles {
		admin.Roles[i] = role.Code
	}

	return &admin, nil
}

// 后台登出
func (lr *UserRepository) AdminUserLogout(userId int64) error {
	err := lr.db.Table("admin_users").Where("user_id = ?", userId).Update("login_status", 0).Error
	if err != nil {
		return err
	}
	return nil
}
