package respositories

import (
	"context"
	"cook-book-admin-backend/config"
	"errors"
	"fmt"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{db: db}
}

// 上传文件
func (r *FileRepository) UploadFile(objectName string, file multipart.File) (string, error) {
	bucketName := "cook-book-yolkmen"
	client := config.Client

	// 检查 OSS 客户端是否初始化
	if client == nil {
		fmt.Println("OSS client is not initialized")
		return "", errors.New("OSS client is not initialized")
	}

	// 创建上传对象的请求
	request := &oss.PutObjectRequest{
		Bucket:  oss.Ptr(bucketName), // 存储空间名称
		Key:     oss.Ptr(objectName), // 对象名称
		Expires: oss.Ptr(time.Hour.String()),
		Body:    file,
	}

	// 执行上传对象的请求
	result, err := client.PutObject(context.TODO(), request)
	if err != nil {
		fmt.Println("failed to put object %v", err)
		return "", err
	}
	fmt.Printf("put object %s success,  %s\n", objectName, result)
	fmt.Println("上传成功", result.OpMetadata)
	// 构造文件 URL
	endpoint := config.Endpoint // 假设 config.Endpoint 提供 OSS 的 endpoint
	if endpoint == "" {
		return "", errors.New("OSS endpoint is not configured")
	}
	url := fmt.Sprintf("https://%s.%s/%s", bucketName, endpoint, objectName)
	return url, nil
}
