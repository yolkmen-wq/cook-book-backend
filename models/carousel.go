package models

import (
	"time"
)

type Carousel struct {
	ID           int64     `json:"carouselId" gorm:"column:carousel_id;primary_key;"`
	CarouselName string    `json:"carouselName" gorm:"column:carousel_name;type:varchar(255);not null;"`
	CarouselDesc string    `json:"carouselDesc" gorm:"column:carousel_desc;type:varchar(255);not null;"`
	Position     string    `json:"position" gorm:"column:position;type:text;not null;"`
	PositionName string    `json:"positionName"`
	CreatedTime  time.Time `json:"createdTime" gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedTime  time.Time `json:"updatedTime,omitempty" gorm:"column:updated_at;type:datetime;autoUpdateTime;"`
}

type CarouselItem struct {
	ID          int64     `json:"carouselItemId" gorm:"column:id;primary_key;"`
	CarouselID  int64     `json:"carouselId" gorm:"column:carousel_id;not null;foreignKey:carousel_id;"`
	Name        string    `json:"name" gorm:"column:name;type:varchar(255);not null;"`
	JumpType    string    `json:"jumpType" gorm:"column:jump_type;type:varchar(255);not null;"`
	ImageURL    string    `json:"imageUrl" gorm:"column:image_url;type:varchar(255);not null;"`
	KeyWord     string    `json:"keyWord" gorm:"column:key_word;type:json;not null;"`
	Sort        int       `json:"sort" gorm:"column:sort;not null;"`
	CreatedTime time.Time `json:"createdTime" gorm:"column:created_at;type:datetime;autoCreateTime"`
	UpdatedTime time.Time `json:"updatedTime,omitempty" gorm:"column:updated_at;type:datetime;autoUpdateTime;"`
}

type GetCarouselsRequest struct {
	CarouselName string `json:"carouselName"`
	PageNum      int    `json:"pageNum"`
	PageSize     int    `json:"pageSize"`
}

type GetCarouselItemsRequest struct {
	Name       string `json:"name"`
	CarouselID int64  `json:"carouselId" `
	PageNum    int    `json:"pageNum"`
	PageSize   int    `json:"pageSize"`
}
