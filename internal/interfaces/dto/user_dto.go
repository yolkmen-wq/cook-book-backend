package dto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type AdminUserResponse struct {
	ID          int64         `json:"id"`
	Username    string        `json:"username"`
	Nickname    string        `json:"nickname"`
	Avatar      string        `json:"avatar"`
	Location    string        `json:"location"`
	IP          string        `json:"ip"`
	Os          string        `json:"os"`
	Browser     string        `json:"browser"`
	LoginStatus int           `json:"loginStatus"`
	LoginTime   time.Time     `json:"loginTime"`
	Status      int           `json:"status"`
	CreateAt    time.Time     `json:"createAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	Roles       []interface{} `json:"roles"`
	Permissions []string      `json:"permissions"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=6"`
	Nickname string `json:"nickname" validate:"required"`
	Avatar   string `json:"avatar"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CustomClaims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}
