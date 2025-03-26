package models

import "time"

type SystemLog struct {
	ID          int       `json:"id" gorm:"column:id;primary_key;"`
	Module      string    `json:"module" gorm:"column:module"`
	Url         string    `json:"url" gorm:"column:url"`
	Method      string    `json:"method" gorm:"column:method"`
	Ip          string    `json:"ip" gorm:"column:ip"`
	Address     string    `json:"address" gorm:"column:address"`
	System      string    `json:"system" gorm:"column:system"`
	Browser     string    `json:"browser" gorm:"column:browser"`
	TakesTime   int64     `json:"takesTime" gorm:"column:takes_time"`
	RequestTime time.Time `json:"requestTime" gorm:"column:request_time"`
}

type AdminUserMgmt struct {
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
	LoginTime   time.Time  `json:"loginTime" gorm:"column:login_time;default:null"`
	Status      *int       `json:"status,omitempty" gorm:"column:status;default:0"`
	Createat    *time.Time `json:"createat" gorm:"column:create_at;autoCreateTime"`
	Updatedat   *time.Time `json:"updatedat" gorm:"column:updated_at;autoUpdateTime"`
}

type GetUsersRequest struct {
	Username string      `json:"username"`
	Nickname string      `json:"nickname"`
	Status   interface{} `json:"status,omitempty" `
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
}
