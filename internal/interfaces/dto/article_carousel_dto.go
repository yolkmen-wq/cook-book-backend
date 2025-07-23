package dto

type ArticleCarouselResponse struct {
	ID           int64  `json:"id"`
	CarouselName string `json:"carouselName"`
	CarouselDesc string `json:"carouselDesc"`
	Position     int    `json:"position"`
	UpdatedAt    string `json:"updatedAt"`
	CreatedAt    string `json:"createdAt"`
}

type ArticleCarouselItemResponse struct {
	ID         int64  `json:"id"`
	CarouselID int64  `json:"carouselId"`
	Name       string `json:"name"`
	JumpType   int    `json:"jumpType"`
	KeyWord    string `json:"keyWord"`
	ImageURL   string `json:"imageUrl"`
	Sort       int    `json:"sort"`
	UpdatedAt  string `json:"updatedAt"`
	CreatedAt  string `json:"createdAt"`
}

type GetArticleCarouselListRequest struct {
	PageNum  int `json:"pageNum" form:"pageNum" validate:"min=1"`
	PageSize int `json:"pageSize" form:"pageSize" validate:"min=1,max=100"`
}