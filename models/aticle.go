package models

import "time"

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
