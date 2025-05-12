package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Article struct {
	ID          int64     `json:"id" gorm:"column:article_id;primary_key;"`
	Title       string    `json:"title" gorm:"column:title;type:varchar(255);not null;"`
	Content     string    `json:"content" gorm:"column:content;type:text;not null;"`
	Cover       string    `json:"cover" gorm:"column:cover;type:varchar(255);not null;"`
	Status      int8      `json:"status" gorm:"column:status;type:tinyint(1);not null;default:0;"`
	Author      string    `json:"author" gorm:"column:author;type:varchar(255);"`
	CreatedTime time.Time `json:"createdTime" gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedTime time.Time `json:"updatedTime,omitempty" gorm:"column:updated_at;type:datetime;autoUpdateTime;"`
}

type ArticleCategory struct {
	ID           int64     `json:"id" gorm:"column:category_id;primary_key;"`
	ParentID     int64     `json:"parentID" gorm:"column:parent_id;not null;default:0;"`
	CategoryName string    `json:"categoryName" gorm:"column:category_name;type:varchar(255);not null;"`
	CategoryPic  string    `json:"categoryPic" gorm:"column:category_pic;type:varchar(255);not null;"`
	CategoryDesc string    `json:"categoryDesc" gorm:"column:category_desc;type:text;not null;"`
	ShowCategory int8      `json:"showCategory" gorm:"column:show_category;type:tinyint(1);not null;default:1;"`
	CategorySort int       `json:"categorySort" gorm:"column:category_sort;type:tinyint(1);not null;default:0;"`
	CreatedTime  time.Time `json:"createdTime" gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedTime  time.Time `json:"updatedTime,omitempty" gorm:"column:updated_at;type:datetime;autoUpdateTime;"`
}

type GetArticlesRequest struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Status      *int8    `json:"status"`
	CreatedTime []string `json:"createdTime"`
	PageNum     int      `json:"pageNum"`
	PageSize    int      `json:"pageSize"`
}

type GetArticleCatsRequest struct {
	CategoryName string `json:"categoryName"`
	PageNum      int    `json:"pageNum"`
	PageSize     int    `json:"pageSize"`
}

func (r *GetArticlesRequest) UnmarshalJSON(data []byte) error {
	type Alias GetArticlesRequest
	aux := &struct {
		Status interface{} `json:"status"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.Status != "" {
		// 根据实际接收到的类型，处理 Status 字段
		if aux.Status != nil {
			switch v := aux.Status.(type) {
			case string:
				statusInt, _ := strconv.ParseInt(v, 10, 8)
				status := int8(statusInt)
				r.Status = &status
			case float64:
				status := int8(v)
				r.Status = &status
			case int:
				status := int8(v)
				r.Status = &status
			case int8:
				r.Status = &v
			default:
				return fmt.Errorf("unmarshaling 'status': got %T, want float64, int, or int8", v)
			}
		}
	} else {
		r.Status = nil
	}

	return nil
}
