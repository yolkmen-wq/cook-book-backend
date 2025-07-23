package dto

type ArticleCateResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	CategoryPic  string `json:"categoryPic"`
	CategoryDesc string `json:"categoryDesc"`
	ShowCategory int    `json:"showCategory"`
	CategorySort int    `json:"categorySort"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type GetArticleCateRequest struct {
	Name     string `json:"name" form:"name"`
	PageNum  int    `json:"pageNum" form:"pageNum" validate:"min=1"`
	PageSize int    `json:"pageSize" form:"pageSize" validate:"min=1,max=100"`
}