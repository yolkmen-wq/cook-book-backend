package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cook-book-backend/internal/domain/services"
)

func main() {
	// 从环境变量获取配置
	apiKey := os.Getenv("AI_API_KEY")
	model := os.Getenv("AI_MODEL")
	provider := os.Getenv("AI_PROVIDER")

	if apiKey == "" || model == "" {
		log.Fatal("请设置 AI_API_KEY 和 AI_MODEL 环境变量")
	}

	// 创建AI服务实例
	var aiService services.AIService
	switch provider {
	case "openai":
		aiService = services.NewOpenAIService(apiKey, model)
	case "deepseek":
		aiService = services.NewDeepSeekService(apiKey, model)
	default:
		// 默认使用DeepSeek
		aiService = services.NewDeepSeekService(apiKey, model)
	}

	// 示例1: 普通聊天
	fmt.Println("=== 普通聊天示例 ===")
	ctx := context.Background()
	response, err := aiService.Chat(ctx, "你好，请简单介绍一下自己")
	if err != nil {
		log.Printf("普通聊天错误: %v", err)
	} else {
		fmt.Printf("AI回复: %s\n\n", response)
	}

	// 示例2: 流式聊天
	fmt.Println("=== 流式聊天示例 ===")
	streamExample(aiService)
}

func streamExample(aiService services.AIService) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 开始流式聊天
	respChan, errorChan := aiService.ChatStream(ctx, "请写一首关于春天的短诗")

	fmt.Print("AI流式回复: ")
	for {
		select {
		case content, ok := <-respChan:
			if !ok {
				// 响应通道已关闭，流式响应结束
				fmt.Println("\n\n流式响应完成")
				return
			}
			// 实时打印接收到的内容
			fmt.Print(content)
		case err, ok := <-errorChan:
			if ok && err != nil {
				fmt.Printf("\n流式聊天错误: %v\n", err)
				return
			}
		case <-ctx.Done():
			fmt.Printf("\n请求超时: %v\n", ctx.Err())
			return
		}
	}
}

// 高级流式聊天示例 - 处理多轮对话
func advancedStreamExample(aiService services.AIService) {
	ctx := context.Background()
	
	messages := []string{
		"请介绍一下Go语言的特点",
		"Go语言在并发编程方面有什么优势？",
		"能给一个Go协程的简单示例吗？",
	}

	for i, message := range messages {
		fmt.Printf("\n=== 第%d轮对话 ===\n", i+1)
		fmt.Printf("用户: %s\n", message)
		fmt.Print("AI: ")

		respChan, errorChan := aiService.ChatStream(ctx, message)
		
		for {
			select {
			case content, ok := <-respChan:
				if !ok {
					fmt.Println("\n")
					goto nextMessage
				}
				fmt.Print(content)
			case err := <-errorChan:
				if err != nil {
					fmt.Printf("\n错误: %v\n", err)
					goto nextMessage
				}
			}
		}
		nextMessage:
		time.Sleep(1 * time.Second) // 稍作停顿
	}
}