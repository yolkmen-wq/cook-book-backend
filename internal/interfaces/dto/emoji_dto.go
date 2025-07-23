package dto

import "time"

type EmojiResponse struct {
	ID           int       `json:"id"`
	CategoryId   int       `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
	Name         string    `json:"name"`
	Url          string    `json:"url"`
	Unicode      string    `json:"unicode"`
	Status       int8      `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type GetEmojiListRequest struct {
	Name       string `json:"name" form:"name"`
	CategoryId int    `json:"categoryId" form:"categoryId"`
	Status     int8   `json:"status" form:"status"`
	PageNum    int    `json:"pageNum" form:"pageNum" validate:"min=1"`
	PageSize   int    `json:"pageSize" form:"pageSize" validate:"min=1,max=100"`
}