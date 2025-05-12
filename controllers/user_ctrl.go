package controllers

import (
	"context"
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/middlewares"
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/services"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mssola/user_agent"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var ctx context.Context

type IPInfo struct {
	Data []map[string]interface{} `json:"data"`
}

func init() {
	ctx = context.Background()
	//ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
}

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

// AdminLogin 管理员登录
func (uc *UserController) AdminLogin(c *gin.Context) {
	// 解析请求参数
	ip := c.Request.RemoteAddr
	if ip != "" {
		idx := strings.LastIndex(ip, ":")
		if idx > 0 {
			ip = ip[:idx]
		}
	}
	location, err := getIPLocation(ip)
	userAgent := c.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	os := ua.OS()
	browser, _ := ua.Browser()

	var adminUser models.AdminUser
	adminUser.IP = ip
	adminUser.Location = location
	adminUser.Os = os
	adminUser.Browser = browser

	if err := c.ShouldBind(&adminUser); err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, "请求参数错误", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	user, err := uc.userService.AdminLogin(adminUser)
	if err != nil {
		if err.Error() == "record not found" {
			response := config.NewResponse(http.StatusOK, false, "用户名或密码错误", nil)
			c.JSONP(http.StatusOK, response)
			return
		}
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 生成AccessToken
	tokenStr, err := middlewares.GenerateAccessToken(user.ID, []string{user.Username}, time.Now().Add(time.Hour*1))
	if err != nil {
		errResp := config.NewResponse(http.StatusInternalServerError, false, "生成AccessToken失败", nil)
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	// 生成RefreshToken
	refreshTokenStr, err := middlewares.GenerateAccessToken(user.ID, []string{user.Username}, time.Now().Add(time.Hour*24))
	if err != nil {
		errResp := config.NewResponse(http.StatusInternalServerError, false, "生成RefreshToken失败", nil)
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	//defer cancel() // 确保在函数结束时取消上下文，释放资源

	// 存储RefreshToken到Redis，注意这里使用了不同的方法名GenerateRefreshToken
	if err := config.RedisClient.Set(ctx, "refreshTokenStr"+strconv.FormatInt(user.ID, 10), refreshTokenStr, time.Hour*24*7).Err(); err != nil {
		fmt.Println("存储RefreshToken==err", err)
		errResp := config.NewResponse(http.StatusInternalServerError, false, "存储RefreshToken到Redis失败", nil)
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "登录成功", map[string]interface{}{
		"avatar":       user.Avatar,
		"username":     user.Username,
		"nickname":     user.Nickname,
		"roles":        user.Roles,
		"permissions":  user.Permissions,
		"accessToken":  tokenStr,
		"refreshToken": refreshTokenStr,
		"expires":      time.Now().Add(time.Minute * 1).Format("2006-01-02 15:04:05"),
	})
	c.JSONP(http.StatusOK, response)
}

// AdminLogout 管理员登出
func (uc *UserController) AdminLogout(c *gin.Context) {
	// 解析token
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		response := config.NewResponse(http.StatusUnauthorized, false, "token 为空", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	tokenStr := token[7:]
	claims, err := middlewares.ValidateToken(tokenStr)
	if err != nil {
		response := config.NewResponse(http.StatusUnauthorized, false, "token 解析失败", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	// 修改登录状态
	if err := uc.userService.AdminUserLogout(claims.ID); err != nil {
		response := config.NewResponse(http.StatusInternalServerError, false, err.Error(), nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 删除RefreshToken
	if err := config.RedisClient.Del(ctx, "refreshTokenStr"+strconv.FormatInt(claims.ID, 10)).Err(); err != nil {
		fmt.Println("删除RefreshToken==err", err)
	}
	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "登出成功", nil)
	c.JSONP(http.StatusOK, response)
}

// RefreshToken 刷新token
func (uc *UserController) AdminRefreshToken(c *gin.Context) {
	type RefreshTokenRequest struct {
		RefreshToken string `json:"refreshToken"`
	}
	var refreshToken RefreshTokenRequest
	err := c.ShouldBind(&refreshToken)

	token, err := jwt.ParseWithClaims(refreshToken.RefreshToken, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.SecretKey, nil
	})

	if err != nil {
		fmt.Println("AdminRefreshToken==err", err)
		errResp := config.NewResponse(http.StatusUnauthorized, false, "token 无效", nil)
		c.JSON(http.StatusUnauthorized, errResp)
		return
	}

	var tokenStr string
	if claims, ok := token.Claims.(*models.CustomClaims); ok {
		tokenStr, err = middlewares.GenerateAccessToken(claims.ID, claims.Audience, time.Now().Add(time.Hour*1))
		if err != nil {
			fmt.Println("AdminRefreshToken==err", err)
			errResp := config.NewResponse(http.StatusInternalServerError, false, "token 刷新失败", nil)
			c.JSON(http.StatusInternalServerError, errResp)
			return
		}

	}

	// 返回数据
	response := config.NewResponse(http.StatusOK, true, "刷新成功", map[string]interface{}{
		"accessToken":  tokenStr,
		"refreshToken": refreshToken.RefreshToken,
		"expires":      time.Now().Add(time.Hour * 1).Format("2006-01-02 15:04:05"),
	})
	c.JSONP(http.StatusOK, response)
}

func getIPLocation(ip string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://opendata.baidu.com/api.php?query=%s&co=&resource_id=6006&oe=utf8", ip))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var info IPInfo
	err = json.Unmarshal(body, &info)

	if err != nil {
		return "", err
	}
	if len(info.Data) == 0 {
		return "", errors.New("未找到IP地址信息")
	}
	location := info.Data[0]["location"]

	return location.(string), nil
}
