package utils

import (
	"cook-book-backend/internal/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// WechatUserInfo 微信用户信息结构体
type WechatUserInfo struct {
	OpenID     string `json:"openid"`
	UnionID    string `json:"unionid"`
	Nickname   string `json:"nickname"`
	HeadImgURL string `json:"headimgurl"`
	Sex        int    `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
}

// WechatAccessTokenResponse 微信access_token响应结构体
type WechatAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionID      string `json:"unionid"`
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

// WechatConfig 微信配置
type WechatConfig struct {
	AppID     string
	AppSecret string
	JWTSecret string // 新增: JWT 签名密钥
	JWTConfig config.JWTConfig
}

// WechatMiniProgramSession 小程序登录会话响应
type WechatMiniProgramSession struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid,omitempty"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// CustomClaims JWT claims 结构
type CustomClaims struct {
	OpenID string `json:"openid"`
	jwt.RegisteredClaims
}

// GetMiniProgramSession 小程序登录 - 通过code获取session
func GetMiniProgramSession(code string, config WechatConfig) (*WechatMiniProgramSession, error) {
	// 构建请求URL
	apiURL := "https://api.weixin.qq.com/sns/jscode2session"
	params := url.Values{}
	params.Set("appid", config.AppID)
	params.Set("secret", config.AppSecret)
	params.Set("js_code", code)
	params.Set("grant_type", "authorization_code")

	requestURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	// 发送HTTP请求
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("请求微信API失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 解析JSON响应
	var session WechatMiniProgramSession
	if err := json.Unmarshal(body, &session); err != nil {
		return nil, fmt.Errorf("解析JSON响应失败: %w", err)
	}

	if session.ErrCode != 0 {
		return nil, fmt.Errorf("微信返回错误: %s (错误码: %d)", session.ErrMsg, session.ErrCode)
	}

	return &session, nil
}

// LoginWithCode 使用code登录并返回JWT token
func LoginWithCode(code string, config WechatConfig) (*WechatMiniProgramSession, string, error) {
	session, err := GetMiniProgramSession(code, config)
	if err != nil {
		return nil, "", fmt.Errorf("获取session失败: %w", err)
	}

	claims := CustomClaims{
		OpenID: session.OpenID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.JWTConfig.AccessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return nil, "", fmt.Errorf("生成token失败: %w", err)
	}

	return session, signedToken, nil
}
