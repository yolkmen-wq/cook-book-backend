package middlewares

import (
	"cook-book-backEnd/config"
	"cook-book-backEnd/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 解析和验证 Token
func ValidateToken(tokenString string) (*models.CustomClaims, error) {
	claims := &models.CustomClaims{}

	// 解析 Token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return config.SecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("无效的 Token")
	}

	return claims, nil
}

// 延长 Token 有效期的函数
func extendToken(tokenString string, newExpiryDuration *jwt.NumericDate) (string, error) {
	// 验证当前 Token
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err // Token 无效
	}

	// 生成新的 Token
	return createToken(claims, newExpiryDuration)
}

// 生成新的 Token
func createToken(claims *models.CustomClaims, expiresIn *jwt.NumericDate) (string, error) {

	// 创建一个新的 Token
	newClaims := &models.CustomClaims{
		claims.ID,
		jwt.RegisteredClaims{
			ExpiresAt: expiresIn,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

	// 对 Token 进行签名
	return token.SignedString(config.SecretKey)
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AuthMiddleWare")

		response := config.Response{
			Success: false,
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
			Data:    "",
		}
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		tokenData, _ := extractJWT(authHeader)
		if authHeader == "" {
			fmt.Println("Authorization 头为空")
			c.JSON(http.StatusUnauthorized, response)
			c.Abort() // 使用c.Abort()停止后续的处理器
			return
		}
		// 验证 Token
		claims, err := ValidateToken(tokenData)
		if err != nil {
			fmt.Println("Token 无效", err)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort() // 使用c.Abort()停止后续的处理器
			return
		}
		//now := time.Now().Unix()
		//expiry := claims.RegisteredClaims.ExpiresAt.Unix()
		//if expiry-now < 1800 {
		//	// 延长token有效期
		//	newExpiryDuration := jwt.NewNumericDate(time.Now().Add(time.Hour * 2))
		//	newToken, err := extendToken(tokenData, newExpiryDuration)
		//	if err != nil {
		//		c.JSON(http.StatusUnauthorized, response)
		//	}
		//	// 更新Authorization头
		//	c.Header("Authorization", "Bearer "+newToken)
		//	claims, err = ValidateToken(newToken)
		//}
		// 设置当前用户信息
		c.Set("claims", claims)
		// 验证权限
		//if claims.Role!= "admin" {
		//	return c.JSON(403, "Forbidden")
		//}
		c.Next()
	}

}

func extractJWT(Authorization string) (string, error) {
	// 检查是否以Bearer开头
	const prefix = "Bearer "
	if !strings.HasPrefix(Authorization, prefix) {
		return "", fmt.Errorf("invalid authorization header format")
	}

	// 去除Bearer前缀，提取JWT
	jwtToken := strings.TrimPrefix(Authorization, prefix)
	return jwtToken, nil
}

func GenerateAccessToken(id int64, Audience []string, expiresTime time.Time) (string, error) {
	// 生成 Token
	//time.Now().Add(time.Hour * 2)
	expiresIn := jwt.NewNumericDate(expiresTime)
	newClaims := models.CustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: expiresIn,                      // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 设置签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 设置生效时间
			Issuer:    "test",                         // 设置签发人
			Subject:   strconv.FormatInt(id, 10),      // 设置主题
			ID:        strconv.FormatInt(id, 10),      // 设置ID
			Audience:  Audience,                       // 设置受众
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

	// 对 Token 进行签名
	return token.SignedString(config.SecretKey)
}
