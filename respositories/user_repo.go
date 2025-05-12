package respositories

import (
	"cook-book-backEnd/models"
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
	admin.Roles = make([]interface{}, len(userRoles))
	for i, role := range roles {
		admin.Roles[i] = role.Code
	}

	return &admin, nil
}

// 前台登录
func (lr *UserRepository) FindUser(username, password string) (*models.User, error) {
	var user models.User
	err := lr.db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 获取路由
func (lr *UserRepository) GetRoutes(userId int64) ([]models.Router, error) {
	//var parentRoutes []models.Router
	//var childRoutes []models.Router
	var routes []models.Router

	// 一次性查询所有路由
	err := lr.db.Table("menus").Where("path is NOT NULL").Find(&routes).Error
	if err != nil {
		return nil, fmt.Errorf("获取路由失败: %w", err)
	}

	// 使用切片映射构建角色名缓存，避免重复查询
	roleNamesCache := make(map[int64][]string)

	// 预先构建子路由映射
	childRouteMap := make(map[int64][]models.Router)
	for _, route := range routes {
		if route.ParentID != 0 {
			childRouteMap[route.ParentID] = append(childRouteMap[route.ParentID], route)
		}
	}

	// 预先填充角色名缓存
	for _, route := range routes {
		roles, err := lr.getRoleNamesFromRoute(route.ID)
		if err != nil {
			return nil, fmt.Errorf("获取路由角色名称失败: %w", err)
		}
		roleNamesCache[route.ID] = roles
	}

	// 构建父子路由关系
	var parentRoutes []models.Router
	for _, route := range routes {
		if route.ParentID == 0 {
			route.Meta = models.Meta{
				Title:      route.Title,
				Icon:       route.Icon,
				Rank:       route.Rank,
				ShowLink:   route.ShowLink,
				ActivePath: route.ActivePath,
			}
			// 使用缓存中的角色名称列表
			route.Meta.Roles = roleNamesCache[route.ID]
			//// 查询权限角色，如果已缓存则直接使用，避免重复查询
			//if roles, found := roleNamesCache[route.ID]; found {
			//	route.Meta.Roles = roles
			//} else {
			//	roles, err := lr.getRoleNamesFromRoute(route.ID)
			//	if err != nil {
			//		return nil, err
			//	}
			//	roleNamesCache[route.ID] = roles
			//	route.Meta.Roles = roles
			//}

			// 递归地处理子路由的子路由
			route.Children = lr.buildNestedRoutes(childRouteMap, route.ID, roleNamesCache)

			// 将父路由添加到列表中
			parentRoutes = append(parentRoutes, route)
		}
	}

	return parentRoutes, nil
}

// 递归函数用于构建多级子路由关系
func (lr *UserRepository) buildNestedRoutes(childRouteMap map[int64][]models.Router, parentID int64, roleNamesCache map[int64][]string) []models.Router {
	var nestedChildren []models.Router
	for _, route := range childRouteMap[parentID] {
		if route.ParentID == parentID {
			route.Meta = models.Meta{
				Title:      route.Title,
				Icon:       route.Icon,
				Rank:       route.Rank,
				ShowLink:   route.ShowLink,
				ActivePath: route.ActivePath,
			}

			// 使用缓存中的角色名称列表
			route.Meta.Roles = roleNamesCache[route.ID]
			//// 查询权限角色，如果已缓存则直接使用，避免重复查询
			//if roles, found := roleNamesCache[route.ID]; found {
			//	route.Meta.Roles = roles
			//} else {
			//	roles, err := lr.getRoleNamesFromRoute(route.ID)
			//	if err != nil {
			//		return nil
			//	}
			//	roleNamesCache[route.ID] = roles
			//	route.Meta.Roles = roles
			//}

			// 递归处理当前路由的子路由
			route.Children = lr.buildNestedRoutes(childRouteMap, route.ID, roleNamesCache)

			// 将子路由添加到嵌套的子路由列表中
			nestedChildren = append(nestedChildren, route)
		}
	}
	return nestedChildren
}

// 辅助函数从路由获取角色名称列表
func (lr *UserRepository) getRoleNamesFromRoute(menuId int64) ([]string, error) {
	var roleCodes []string

	err := lr.db.Table("role_menus").Select("roles.role_code").Joins("JOIN roles ON role_menus.role_id = roles.role_id").Where("role_menus.menu_id = ?", menuId).Scan(&roleCodes).Error

	if err != nil {
		return nil, err
	}

	return roleCodes, nil
}
