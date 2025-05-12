package config

import (
	"context"
	"flag"
	"fmt"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 定义一个密钥
var SecretKey = []byte("lyf123456")

// 定义全局变量
var (
	region        string // 存储区域
	bucketName    string // 存储空间名称
	AccesskeyIdid string // 阿里云AccessKeyId
	Signature     string // 阿里云Signature
	Endpoint      string // 存储空间Endpoint
	DB            *gorm.DB
	RedisClient   *redis.Client
	Client        *oss.Client
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
	Total       int64       `json:"total"`       // 总数
	CurrentPage int         `json:"currentPage"` // 当前页码
	PageSize    int         `json:"pageSize"`    // 每页数量
	List        interface{} `json:"list"`        // 数据列表
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

	// 建立Redis连接
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "47.121.201.137:6011",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("连接Redis中...", RedisClient)
	if RedisClient != nil {
		fmt.Println("连接Redis成功！")
	}
	initOss()
}

// 连接OSS
func initOss() {
	bucketName = "cook-book-yolkmen"
	region = "cn-shenzhen"
	Endpoint = "oss-cn-shenzhen.aliyuncs.com"
	AccesskeyIdid = "TMP.3Ksu45ymtAGrX4xgwLYjmhBWCrzATRu1Z1sZybvLKBpZBkt89HazdgjSZfDr45btD8J5Hrh6MuKy9npeiydDqeZD7KbJNG"
	Signature = "MN9qfwpWo3B9vb1qmyDiXt6BxtU%3D"

	// 检查bucket名称是否为空
	if len(bucketName) == 0 {
		flag.PrintDefaults()
		log.Fatalf("invalid parameters, bucket name required")
	}

	// 检查region是否为空
	if len(region) == 0 {
		flag.PrintDefaults()
		log.Fatalf("invalid parameters, region required")
	}

	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewEnvironmentVariableCredentialsProvider()).
		WithRegion(region).
		WithEnableAutoDetectCloudBoxId(true)

	// 创建OSS客户端
	Client = oss.NewClient(cfg)

	// 检查是否存在目标bucket
	// 检查存储空间是否存在
	result, err := Client.IsBucketExist(context.TODO(), bucketName)
	if err != nil {
		log.Fatalf("failed to check if bucket exists %v", err)
	}

	// 打印检查结果
	log.Printf("is bucket exist: %#v\n", result)

	if !result {
		request := &oss.PutBucketRequest{
			Bucket: oss.Ptr(bucketName), // 存储空间名称
		}
		// 发送创建存储空间的请求
		result, err := Client.PutBucket(context.TODO(), request)
		if err != nil {
			log.Fatalf("failed to put bucket %v", err)
		}

		// 打印创建存储空间的结果
		log.Printf("put bucket result:%#v\n", result)
	}

}
