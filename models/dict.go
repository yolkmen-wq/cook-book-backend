package models

import "time"

type Dict struct {
	ID         int       `json:"dictId" gorm:"column:dict_id;primary_key"`
	DictName   string    `json:"dictName" gorm:"column:dict_name"`
	DictType   string    `json:"dictType" gorm:"column:dict_type"`
	Status     int8      `json:"status" gorm:"column:status"`
	Remark     string    `json:"remark" gorm:"column:remark"`
	CreateTime time.Time `json:"createTime" gorm:"column:created_at;type:datetime;autoCreateTime"`
	CreatedBy  string    `json:"createdBy" gorm:"column:created_by;default:''"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updated_at;type:datetime;autoUpdateTime"`
	UpdatedBy  string    `json:"updatedBy" gorm:"column:updated_by;default:''"`
}

type DictData struct {
	DicCode    int       `json:"dictCode" gorm:"column:dict_code;primary_key"`
	DictLabel  string    `json:"dictLabel" gorm:"column:dict_label"`
	DictValue  string    `json:"dictValue" gorm:"column:dict_value"`
	DictType   string    `json:"dictType" gorm:"column:dict_type"`
	DictSort   int       `json:"dictSort" gorm:"column:dict_sort"`
	IsDefault  string    `json:"isDefault" gorm:"column:is_default"`
	Status     int8      `json:"status" gorm:"column:status"`
	CreateTime time.Time `json:"createTime" gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updated_at;type:datetime;autoUpdateTime"`
}

type GetDictListRequest struct {
	DictName   string   `json:"dictName"`
	DictType   string   `json:"dictType"`
	Status     string   `json:"status"`
	CreateTime []string `json:"createTime"`
	PageNum    int      `json:"pageNum"`
	PageSize   int      `json:"pageSize"`
}

type GetDictDataListRequest struct {
	DictType   string   `json:"dictType"`
	CreateTime []string `json:"createTime"`
	PageNum    int      `json:"pageNum"`
	PageSize   int      `json:"pageSize"`
}
