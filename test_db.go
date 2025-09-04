package main

import (
	"fmt"
	"time"

	"cook-book-backend/internal/config"
	"cook-book-backend/internal/infrastructure/database"
)

func main() {
	fmt.Println("开始测试数据库连接...")
	
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		return
	}
	
	fmt.Printf("尝试连接数据库: %s:%d\n", cfg.Database.Host, cfg.Database.Port)
	
	// 设置连接超时
	done := make(chan bool, 1)
	var dbErr error
	
	go func() {
		_, dbErr = database.New(cfg.Database)
		done <- true
	}()
	
	// 等待连接或超时
	select {
	case <-done:
		if dbErr != nil {
			fmt.Printf("数据库连接失败: %v\n", dbErr)
		} else {
			fmt.Println("数据库连接成功!")
		}
	case <-time.After(10 * time.Second):
		fmt.Println("数据库连接超时 (10秒)")
		fmt.Println("可能的原因:")
		fmt.Println("1. 网络连接问题")
		fmt.Println("2. 数据库服务器不可达")
		fmt.Println("3. 防火墙阻止连接")
		fmt.Println("4. 数据库配置错误")
	}
}