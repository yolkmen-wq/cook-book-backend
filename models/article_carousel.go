package models

type ArticleCarousel struct {
	ID           int64  `json:"carousel_id" gorm:"column:carousel_id;primary_key"`
	CarouselName string `json:"carouselName" gorm:"column:carousel_name"`
	CarouselDesc string `json:"carouselDesc" gorm:"column:carousel_desc"`
	Position     int    `json:"position" gorm:"column:position"`
	UpdatedAt    string `json:"updatedTime" gorm:"column:updated_at"`
	CreatedAt    string `json:"createdTime" gorm:"column:created_at"`
}

type ArticleCarouselItem struct {
	ID         int64  `json:"id" gorm:"column:id;primary_key"`
	CarouselID int64  `json:"carouselId" gorm:"column:carousel_id"`
	Name       string `json:"name" gorm:"column:name"`
	JumpType   int    `json:"jumpType" gorm:"column:jump_type"`
	KeyWord    string `json:"keyWord" gorm:"column:key_word"`
	ImageURL   string `json:"imageUrl" gorm:"column:image_url"`
	Sort       int    `json:"sort" gorm:"column:sort"`
	UpdatedAt  string `json:"updatedTime" gorm:"column:updated_at"`
	CreatedAt  string `json:"createdTime" gorm:"column:created_at"`
}
