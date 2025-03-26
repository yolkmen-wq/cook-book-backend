package article_ctrl

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/services/article_srv"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleCatCtrl struct {
	articleCatSrv article_srv.ArticleCatService
}

func NewArticleCatController(articleService article_srv.ArticleCatService) *ArticleCatCtrl {
	return &ArticleCatCtrl{articleService}
}

func (acc *ArticleCatCtrl) CreateArticleCat(c *gin.Context) {
	var articleCat models.ArticleCategory
	err := c.ShouldBindJSON(&articleCat)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = acc.articleCatSrv.CreateArticleCat(&articleCat)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "创建成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (acc *ArticleCatCtrl) GetArticleCatList(c *gin.Context) {
	articleCats, err := acc.articleCatSrv.GetArticleCats()
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", config.ListResponse{List: articleCats, Total: int64(len(articleCats)), CurrentPage: 1, PageSize: 10})
	c.JSONP(http.StatusOK, response)
}

func (acc *ArticleCatCtrl) UpdateArticleCat(c *gin.Context) {
	var articleCat models.ArticleCategory
	err := c.ShouldBindJSON(&articleCat)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = acc.articleCatSrv.UpdateArticleCat(&articleCat)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "更新成功", nil)
	c.JSONP(http.StatusOK, response)
}

func (acc *ArticleCatCtrl) DeleteArticleCat(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = acc.articleCatSrv.DeleteArticleCat(idInt)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "删除成功", nil)
	c.JSONP(http.StatusOK, response)
}
