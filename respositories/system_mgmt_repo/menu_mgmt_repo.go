package system_mgmt_repo

import (
	"context"
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/models"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type MenuMgmtRepository struct {
	db *gorm.DB
}

var redisClient = config.RedisClient

func NewMenuMgmtRepository(db *gorm.DB) *MenuMgmtRepository {
	return &MenuMgmtRepository{db: db}
}

// 清除路由缓存
func (m *MenuMgmtRepository) ClearRoutesCache() error {
	ctx := context.Background()
	cacheKey := "system:routes:all"
	return redisClient.Del(ctx, cacheKey).Err()
}

// 获取全部路由
func (m *MenuMgmtRepository) GetRoutesWithLock() ([]models.Router, error) {
	ctx := context.Background()
	cacheKey := "system:routes:all"
	lockKey := "system:routes:lock"

	// 尝试从Redis获取缓存
	cachedData, err := redisClient.Get(ctx, cacheKey).Bytes()
	if err == nil {
		// 缓存命中，解析JSON数据
		var routes []models.Router
		err = json.Unmarshal(cachedData, &routes)
		if err == nil {
			return routes, nil // 成功返回缓存数据
		}
	}

	// 获取分布式锁，防止缓存击穿
	lockAcquired, err := redisClient.SetNX(ctx, lockKey, "1", 30*time.Second).Result()
	if err != nil {
		return nil, err
	}

	if !lockAcquired {
		// 未获取到锁，等待一段时间后重试获取缓存
		time.Sleep(100 * time.Millisecond)
		cachedData, err = redisClient.Get(ctx, cacheKey).Bytes()
		if err == nil {
			var routes []models.Router
			if err = json.Unmarshal(cachedData, &routes); err == nil {
				return routes, nil
			}
		}
		// 如果仍未获取到缓存，返回指示需要重试的错误
		return nil, fmt.Errorf("正在重建缓存，请稍后重试")
	}

	// 获取到锁，重新查询缓存（双重检查）
	cachedData, err = redisClient.Get(ctx, cacheKey).Bytes()
	if err == nil {
		var routes []models.Router
		err = json.Unmarshal(cachedData, &routes)
		if err == nil {
			// 释放锁
			redisClient.Del(ctx, lockKey)
			return routes, nil
		}
	}

	// 设置锁自动释放
	defer redisClient.Del(ctx, lockKey)

	// 缓存未命中或解析失败，从数据库获取
	return m.loadRoutesFromDB(ctx, cacheKey)
}

// loadRoutesFromDB 从数据库加载路由并缓存
func (m *MenuMgmtRepository) loadRoutesFromDB(ctx context.Context, cacheKey string) ([]models.Router, error) {
	var routes []models.Router

	// 一次性查询所有路由，优化查询
	err := m.db.Table("menus").
		Select("id, parent_id, path, name, component, redirect, title, icon, `rank`, show_link, active_path").
		Where("path IS NOT NULL").
		Find(&routes).Error
	if err != nil {
		return nil, fmt.Errorf("获取路由失败: %w", err)
	}

	// 使用单次查询获取所有路由角色关联
	var roleAssociations []struct {
		MenuID   int64
		RoleID   int64
		RoleCode string
	}
	err = m.db.Table("role_menus rms").
		Select("rms.menu_id, rms.role_id, roles.role_code").
		Joins("LEFT JOIN roles ON rms.role_id = roles.role_id").
		Where("rms.menu_id IN ?", getMenuIDs(routes)).
		Find(&roleAssociations).Error
	if err != nil {
		return nil, fmt.Errorf("获取角色关联失败: %w", err)
	}

	// 构建角色映射
	roleNamesCache := make(map[int64][]string)
	for _, assoc := range roleAssociations {
		roleNamesCache[assoc.MenuID] = append(roleNamesCache[assoc.MenuID], assoc.RoleCode)
	}

	// 预先构建子路由映射
	childRouteMap := make(map[int64][]models.Router)
	for _, route := range routes {
		if route.ParentID != 0 {
			childRouteMap[route.ParentID] = append(childRouteMap[route.ParentID], route)
		}
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
				Roles:      roleNamesCache[route.ID],
			}

			// 递归地处理子路由
			route.Children = m.buildNestedRoutes(childRouteMap, route.ID, roleNamesCache)

			// 将父路由添加到列表中
			parentRoutes = append(parentRoutes, route)
		}
	}

	// 将处理后的结果存入Redis缓存
	resultJSON, err := json.Marshal(parentRoutes)
	if err == nil {
		// 设置缓存，有效期延长到24小时
		redisClient.Set(ctx, cacheKey, resultJSON, 24*time.Hour)
	}

	return parentRoutes, nil
}

// getMenuIDs 辅助函数，提取所有菜单ID
func getMenuIDs(routes []models.Router) []int64 {
	ids := make([]int64, 0, len(routes))
	for _, route := range routes {
		ids = append(ids, route.ID)
	}
	return ids
}

// 递归函数用于构建多级子路由关系
func (m *MenuMgmtRepository) buildNestedRoutes(childRouteMap map[int64][]models.Router, parentID int64, roleNamesCache map[int64][]string) []models.Router {
	children := childRouteMap[parentID]
	nestedChildren := make([]models.Router, 0, len(children))
	for _, route := range children {
		if route.ParentID == parentID {
			route.Meta = models.Meta{
				Title:      route.Title,
				Icon:       route.Icon,
				Rank:       route.Rank,
				ShowLink:   route.ShowLink,
				ActivePath: route.ActivePath,
				Roles:      roleNamesCache[route.ID],
			}

			// 递归处理当前路由的子路由
			route.Children = m.buildNestedRoutes(childRouteMap, route.ID, roleNamesCache)

			// 将子路由添加到嵌套的子路由列表中
			nestedChildren = append(nestedChildren, route)
		}
	}
	return nestedChildren
}

// 获取菜单列表
func (m *MenuMgmtRepository) GetMenuList(searchInfo models.GetMenuRequest) ([]models.Menu, int64, error) {
	var menus []models.Menu
	var count int64

	err := m.db.Where("name like?", "%"+searchInfo.Name+"%").Find(&menus).Count(&count).Error
	if err != nil {
		fmt.Println("获取菜单列表失败", err)
		return nil, 0, err
	}
	return menus, count, nil
}

// 获取菜单详情
func (m *MenuMgmtRepository) GetMenuDetail(id int64) (models.Menu, error) {
	var menu models.Menu
	err := m.db.First(&menu, id).Error
	if err != nil {
		fmt.Println("获取菜单详情失败", err)
		return models.Menu{}, err
	}
	return menu, nil
}

// 创建菜单
func (m *MenuMgmtRepository) CreateMenu(menu models.Menu) error {
	err := m.db.Create(&menu).Error
	if err != nil {
		fmt.Println("创建菜单失败", err)
		return err
	}
	err = m.ClearRoutesCache()
	if err != nil {
		return err
	}
	return nil
}

// 更新菜单
func (m *MenuMgmtRepository) UpdateMenu(menu models.Menu) error {
	err := m.db.Save(&menu).Error
	if err != nil {
		fmt.Println("更新菜单失败", err)
		return err
	}
	err = m.ClearRoutesCache()
	if err != nil {
		fmt.Println("删除缓存失败", err)
		return err
	}
	return nil
}

// 删除菜单
func (m *MenuMgmtRepository) DeleteMenu(id int64) error {
	err := m.db.Delete(&models.Menu{}, id).Error
	if err != nil {
		fmt.Println("删除菜单失败", err)
		return err
	}
	err = m.ClearRoutesCache()
	if err != nil {
		fmt.Println("删除缓存失败", err)
		return err
	}
	return nil
}
