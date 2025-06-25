package models

import (
	"time"
)

type Emoji struct {
	ID           int       `json:"id" gorm:"column:id;primary_key"`
	CategoryId   string    `json:"categoryId" gorm:"column:category_id"`
	CategoryName string    `json:"categoryName" gorm:"column:category_name"`
	Name         string    `json:"name" gorm:"column:name"`
	Url          string    `json:"url" gorm:"column:url"`
	Unicode      string    `json:"unicode" gorm:"column:unicode;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci"`
	Status       int8      `json:"status" gorm:"column:status"`
	CreatedAt    time.Time `json:"createdTime" gorm:"column:created_at;type:datetime;autoCreateTime;"`
	UpdatedAt    time.Time `json:"updatedTime" gorm:"column:updated_at;type:datetime;autoUpdateTime;"`
}

//
//type GetEmojisRequest struct {
//	Name        string   `json:"name"`
//	CategoryId  string   `json:"categoryId"`
//	Status      *int8    `json:"status"`
//	CreatedTime []string `json:"createdTime"`
//	PageNum     int      `json:"pageNum"`
//	PageSize    int      `json:"pageSize"`
//}
//
//func (r *GetEmojisRequest) UnmarshalJSON(data []byte) error {
//	type Alias GetEmojisRequest
//	aux := &struct {
//		Status interface{} `json:"status"`
//		*Alias
//	}{
//		Alias: (*Alias)(r),
//	}
//	if err := json.Unmarshal(data, &aux); err != nil {
//		return err
//	}
//
//	if aux.Status != "" {
//		// 根据实际接收到的类型，处理 Status 字段
//		if aux.Status != nil {
//			switch v := aux.Status.(type) {
//			case string:
//				statusInt, _ := strconv.ParseInt(v, 10, 8)
//				status := int8(statusInt)
//				r.Status = &status
//			case float64:
//				status := int8(v)
//				r.Status = &status
//			case int:
//				status := int8(v)
//				r.Status = &status
//			case int8:
//				r.Status = &v
//			default:
//				return fmt.Errorf("unmarshaling 'status': got %T, want float64, int, or int8", v)
//			}
//		}
//	} else {
//		r.Status = nil
//	}
//
//	return nil
//}
