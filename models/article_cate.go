package models

type ArticleCate struct {
	ID           int64  `json:"id" gorm:"column:category_id;primary_key"`
	Name         string `json:"name" gorm:"column:category_name"`
	CategoryPic  string `json:"categoryPic" gorm:"column:category_pic"`
	CategoryDesc string `json:"categoryDesc" gorm:"column:category_desc"`
	ShowCategory int    `json:"showCategory" gorm:"column:show_category"`
	CategorySort int    `json:"categorySort" gorm:"column:category_sort"`
	CreatedAt    string `json:"createdTime" gorm:"column:created_at"`
	UpdatedAt    string `json:"updatedTime" gorm:"column:updated_at"`
}

type GetArticleCateRequest struct {
	Name     string `json:"name,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type GetArticleCateResponse struct {
	List     []ArticleCate `json:"data"`
	Total    int64         `json:"total"`
	PageNum  int           `json:"pageNum"`
	PageSize int           `json:"pageSize"`
	Error    error         `json:"error"`
}
