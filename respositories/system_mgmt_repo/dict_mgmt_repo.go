package system_mgmt_repo

import (
	"cook-book-admin-backend/models"
	"fmt"
	"gorm.io/gorm"
)

type DictMgmtRepository struct {
	db *gorm.DB
}

func NewDictMgmtRepository(db *gorm.DB) *DictMgmtRepository {
	return &DictMgmtRepository{db: db}
}

// GetDictList 获取字典列表
func (d *DictMgmtRepository) GetDictList(req *models.GetDictListRequest) ([]models.Dict, int64, int, int, error) {
	var dicts []models.Dict
	var total int64

	// 构建查询条件
	db := d.db.Table("dicts")

	if req.DictName != "" {
		db = db.Where("dict_name LIKE ?", "%"+req.DictName+"%")
	}
	if req.DictType != "" {
		db = db.Where("dict_name LIKE ?", "%"+req.DictType+"%")
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	if len(req.CreateTime) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", req.CreateTime[0], req.CreateTime[1])
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		fmt.Println("获取字典总数失败", err)
		return nil, 0, 0, 0, err
	}

	// 获取分页用户列表
	if req.PageNum != 0 && req.PageSize != 0 {
		if err := db.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&dicts).Error; err != nil {
			fmt.Println("获取字典列表失败", err)
			return nil, 0, 0, 0, err
		}
	} else {
		if err := db.Find(&dicts).Error; err != nil {
			fmt.Println("获取字典列表失败", err)
			return nil, 0, 0, 0, err
		}
	}

	return dicts, total, req.PageNum, req.PageSize, nil
}

// CreateDict 创建字典
func (d *DictMgmtRepository) CreateDict(dict *models.Dict) error {
	err := d.db.Create(&dict).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateDict 更新字典
func (d *DictMgmtRepository) UpdateDict(dict *models.Dict) error {
	err := d.db.Updates(dict).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteDict 删除字典
func (d *DictMgmtRepository) DeleteDict(id int64) error {
	err := d.db.Delete(&models.Dict{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// GetDictDatas 获取字典项列表
func (d *DictMgmtRepository) GetDictData(req *models.GetDictDataListRequest) ([]models.DictData, int64, int, int, error) {
	var dictData []models.DictData
	var total int64

	// 构建查询条件
	db := d.db.Table("dict_data")

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		fmt.Println("获取字典总数失败", err)
		return nil, 0, 0, 0, err
	}

	err := db.Where("dict_type = ?", req.DictType).Find(&dictData).Error
	if err != nil {
		return nil, 0, 0, 0, err
	}
	return dictData, total, req.PageNum, req.PageSize, nil
}

// AddDictData 新增字典项
func (d *DictMgmtRepository) AddDictData(dictData *models.DictData) error {
	err := d.db.Create(&dictData).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateDictData 更新字典项
func (d *DictMgmtRepository) UpdateDictData(dictData *models.DictData) error {
	err := d.db.Updates(dictData).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteDictData 删除字典项
func (d *DictMgmtRepository) DeleteDictData(id int64) error {
	err := d.db.Delete(&models.DictData{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
