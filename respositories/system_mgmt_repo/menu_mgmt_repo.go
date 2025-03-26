package system_mgmt_repo

import (
	"cook-book-admin-backend/models"
	"fmt"
	"gorm.io/gorm"
)

type MenuMgmtRepository struct {
	db *gorm.DB
}

func NewMenuMgmtRepository(db *gorm.DB) *MenuMgmtRepository {
	return &MenuMgmtRepository{db: db}
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
	return nil
}

// 更新菜单
func (m *MenuMgmtRepository) UpdateMenu(menu models.Menu) error {
	err := m.db.Save(&menu).Error
	if err != nil {
		fmt.Println("更新菜单失败", err)
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
	return nil
}
