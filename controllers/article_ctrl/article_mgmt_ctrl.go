package article_ctrl

import (
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/services/article_srv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleMgmtController struct {
	articleService article_srv.ArticleMgmtService
}

func NewArticleMgmtController(articleService article_srv.ArticleMgmtService) *ArticleMgmtController {
	return &ArticleMgmtController{articleService}
}

func (amc *ArticleMgmtController) GetArticleList(c *gin.Context) {
	articleList, err := amc.articleService.GetArticleList()
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, "请求参数错误", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "获取成功", map[string]interface{}{
		"list": articleList,
	})
	c.JSONP(http.StatusOK, response)
}

func (amc *ArticleMgmtController) CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		fmt.Println(err)
		response := config.NewResponse(http.StatusBadRequest, false, "请求参数错误", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := amc.articleService.CreateArticle(article)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, "创建失败", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := config.NewResponse(http.StatusOK, true, "创建成功", nil)
	c.JSON(http.StatusOK, response)
}

func (amc *ArticleMgmtController) UpdateArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		fmt.Println(err)
		response := config.NewResponse(http.StatusBadRequest, false, "请求参数错误", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := amc.articleService.UpdateArticle(article)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, "更新失败", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := config.NewResponse(http.StatusOK, true, "更新成功", nil)
	c.JSON(http.StatusOK, response)
}

func (amc *ArticleMgmtController) DeleteArticle(c *gin.Context) {
	fmt.Println("delete article")
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	fmt.Println(id, idInt)
	if err != nil {
		response := config.NewResponse(http.StatusBadRequest, false, "请求参数错误", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = amc.articleService.DeleteArticle(idInt)
	if err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, "删除失败", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := config.NewResponse(http.StatusOK, true, "删除成功", nil)
	c.JSON(http.StatusOK, response)
}
