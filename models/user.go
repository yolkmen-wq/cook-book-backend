package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type User struct {
	ID       int64  `json:"id" gorm:"column:user_id;primary_key"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Nickname string `json:"nickname" gorm:"column:nickname"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
	//Roles      []string `json:"roles"`
	//Permission int `json:"permission"`
}

type AdminUser struct {
	ID          int64      `json:"id" gorm:"column:user_id;primary_key"`
	Username    string     `json:"username" gorm:"column:username"`
	Password    string     `json:"password" gorm:"column:password"`
	Nickname    string     `json:"nickname" gorm:"column:nickname"`
	Avatar      string     `json:"avatar" gorm:"column:avatar"`
	Location    string     `json:"location" gorm:"column:location"`
	IP          string     `json:"ip" gorm:"column:ip"`
	Os          string     `json:"os" gorm:"column:os"`
	Browser     string     `json:"browser" gorm:"column:browser"`
	LoginStatus int        `json:"loginStatus,omitempty" gorm:"column:login_status;default:0"`
	LoginTime   time.Time  `json:"loginTime" gorm:"column:login_time"`
	Status      *int       `json:"status,omitempty" gorm:"column:status;default:0"`
	Createat    *time.Time `json:"createat" gorm:"column:create_at"`
	Updatedat   *time.Time `json:"updatedat" gorm:"column:updated_at"`
	Roles       []string   `json:"roles"  gorm:"type:json"`
	Permissions []string   `json:"permissions" gorm:"type:json"`
}

type Permission struct {
	ID   int64  `gorm:"column:permission_id;primary_key"`
	Name string `gorm:"column:permission_name"`
}

type CustomClaims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}

type Router struct {
	ID         int64    `json:"-" gorm:"column:id;primary_key"`
	Path       string   `json:"path" gorm:"column:path"`
	Component  string   `json:"component,omitempty" gorm:"column:component"`
	Name       string   `json:"name" gorm:"column:name"`
	Redirect   string   `json:"redirect,omitempty" gorm:"column:redirect"`
	ParentID   int64    `json:"-" gorm:"column:parent_id"`
	Title      string   `json:"-" gorm:"column:title"`
	Icon       string   `json:"-" gorm:"column:icon"`
	Rank       int64    `json:"-" gorm:"column:rank"`
	ActivePath string   `json:"-" gorm:"column:active_path"`
	ShowLink   bool     `json:"-" gorm:"column:show_link"`
	Meta       Meta     `json:"meta" gorm:"foreignKey:Meta"`
	Children   []Router `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	Createdat  string   `json:"createdTime" gorm:"column:created_at"`
	Updatedat  string   `json:"updatedTime" gorm:"column:updated_at"`
}

type Meta struct {
	Title      string   `json:"title,omitempty"`
	Icon       string   `json:"icon,omitempty"`
	ActivePath string   `json:"activePath,omitempty"`
	ShowLink   bool     `json:"showLink"`
	Rank       int64    `json:"rank,omitempty"`
	Roles      []string `json:"roles,omitempty"`
}

// 实现 Valuer 接口
func (m Meta) Value() (driver.Value, error) {
	// 将 Meta 转换为数据库可以存储的值，例如 JSON 字符串
	return json.Marshal(m)
}

// 实现 Scanner 接口
func (m *Meta) Scan(value interface{}) error {
	// 将数据库中的值转换回 Meta 类型，例如 JSON 字符串
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}

type RouterPermission struct {
	MenuID       int64     `gorm:"column:menu_id"`
	PermissionID int64     `gorm:"column:permission_id"`
	CreatedTime  time.Time `gorm:"column:created_at"`
	UpdatedTime  time.Time `gorm:"column:updated_at"`
}
