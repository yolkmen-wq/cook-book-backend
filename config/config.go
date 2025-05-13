package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义全局变量
var (
	SecretKey = []byte("lyf123456") // 密钥
	DB        *gorm.DB
)

// 定义统一的返回格式结构体
type Response struct {
	Code    int         `json:"code"`    // 状态码，0表示成功，非0表示错误
	Success bool        `json:"success"` // 是否成功
	Message string      `json:"message"` // 返回的消息
	Data    interface{} `json:"data"`    // 返回的数据，可以是任意类型
}

// 定义分页返回格式结构体
type ListResponse struct {
	Total    int64       `json:"total"`    // 总数
	PageNum  int         `json:"pageNum"`  // 当前页码
	PageSize int         `json:"pageSize"` // 每页数量
	List     interface{} `json:"list"`     // 数据列表
}

func NewResponse(code int, success bool, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Success: success,
		Message: message,
		Data:    data,
	}
}

func init() {
	// 建立数据库连接
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:wuqi9457@tcp(47.121.201.137:3306)/cook_book?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                    // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("初始化数据库成功！")
}
