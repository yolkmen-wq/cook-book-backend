package handlers

import (
	"cook-book-backend/internal/domain/services"
	"cook-book-backend/internal/infrastructure/logger"
	"cook-book-backend/internal/interfaces/dto"
	"cook-book-backend/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	*BaseHandler
	articleService services.ArticleService
}

func NewArticleHandler(articleService services.ArticleService, logger logger.Logger) *ArticleHandler {
	return &ArticleHandler{
		BaseHandler:    NewBaseHandler(logger),
		articleService: articleService,
	}
}

// GetArticleList 获取文章列表
func (h *ArticleHandler) GetArticleList(c *gin.Context) {
	var req dto.GetArticleListRequest

	// 设置默认值
	req.PageNum = h.GetIntQueryParam(c, "pageNum", 1)
	req.PageSize = h.GetIntQueryParam(c, "pageSize", 10)
	req.CategoryID = h.GetIntQueryParam(c, "categoryId", 0)
	req.Keyword = h.GetQueryParam(c, "keyword", "")

	// 调用服务层
	list, total, err := h.articleService.GetArticleList(&req)
	if err != nil {
		h.HandleError(c, err, "get_article_list")
		return
	}

	// 构建分页信息
	pagination := &response.Pagination{
		Page:      req.PageNum,
		PageSize:  req.PageSize,
		Total:     total,
		TotalPage: int((total + int64(req.PageSize) - 1) / int64(req.PageSize)),
	}

	h.SuccessWithPagination(c, list, pagination)
}

// GetLatestArticles 获取最新文章
func (h *ArticleHandler) GetLatestArticles(c *gin.Context) {
	limit := h.GetIntQueryParam(c, "limit", 10)

	list, total, err := h.articleService.GetLatestArticles(limit)
	if err != nil {
		h.HandleError(c, err, "get_latest_articles")
		return
	}

	// 构建分页信息
	pagination := &response.Pagination{
		Page:      1,
		PageSize:  limit,
		Total:     total,
		TotalPage: int((total + int64(limit) - 1) / int64(limit)),
	}

	h.SuccessWithPagination(c, list, pagination)
}

// GetArticleDetail 获取文章详情
func (h *ArticleHandler) GetArticleDetail(c *gin.Context) {
	id, err := h.GetIDParam(c, "id")
	if err != nil {
		h.HandleError(c, err, "get_article_id_param")
		return
	}

	article, err := h.articleService.GetArticleDetail(id)
	if err != nil {
		h.HandleError(c, err, "get_article_detail")
		return
	}

	h.Success(c, article)
}

// // CreateArticle 创建文章
// func (h *ArticleHandler) CreateArticle(c *gin.Context) {
// 	var req dto.CreateArticleRequest

// 	if err := h.BindAndValidate(c, &req); err != nil {
// 		h.HandleError(c, err, "bind_create_article_params")
// 		return
// 	}

// 	article, err := h.articleService.CreateArticle(&req)
// 	if err != nil {
// 		h.HandleError(c, err, "create_article")
// 		return
// 	}

// 	h.Success(c, article)
// }

// // UpdateArticle 更新文章
// func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
// 	id, err := h.GetIDParam(c, "id")
// 	if err != nil {
// 		h.HandleError(c, err, "get_article_id_param")
// 		return
// 	}

// 	var req dto.UpdateArticleRequest
// 	if err := h.BindAndValidate(c, &req); err != nil {
// 		h.HandleError(c, err, "bind_update_article_params")
// 		return
// 	}

// 	article, err := h.articleService.UpdateArticle(id, &req)
// 	if err != nil {
// 		h.HandleError(c, err, "update_article")
// 		return
// 	}

// 	h.Success(c, article)
// }

// // DeleteArticle 删除文章
// func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
// 	id, err := h.GetIDParam(c, "id")
// 	if err != nil {
// 		h.HandleError(c, err, "get_article_id_param")
// 		return
// 	}

// 	if err := h.articleService.DeleteArticle(id); err != nil {
// 		h.HandleError(c, err, "delete_article")
// 		return
// 	}

// 	h.Success(c, nil)
// }
