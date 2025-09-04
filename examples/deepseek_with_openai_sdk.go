package main

import (
	"context"
	"fmt"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

// 使用OpenAI SDK调用DeepSeek的示例
func main() {
	// DeepSeek API配置
	apiKey := "your-deepseek-api-key"
	baseURL := "https://api.deepseek.com/v1" // DeepSeek的API端点
	model := "deepseek-chat"

	// 创建OpenAI客户端配置，指向DeepSeek的API
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseURL // 重要：将BaseURL设置为DeepSeek的API地址

	// 创建客户端
	client := openai.NewClientWithConfig(config)

	// 发送聊天请求
	ctx := context.Background()
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "你好，请介绍一下你自己",
				},
			},
			MaxTokens:   1000,
			Temperature: 0.7,
		},
	)

	if err != nil {
		log.Fatalf("ChatCompletion error: %v\n", err)
	}

	fmt.Printf("DeepSeek回复: %s\n", resp.Choices[0].Message.Content)
}

// DeepSeekService 使用OpenAI SDK的封装版本
type DeepSeekService struct {
	client *openai.Client
	model  string
}

// NewDeepSeekServiceWithOpenAI 创建使用OpenAI SDK的DeepSeek服务
func NewDeepSeekServiceWithOpenAI(apiKey, model string) *DeepSeekService {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.deepseek.com/v1"
	
	return &DeepSeekService{
		client: openai.NewClientWithConfig(config),
		model:  model,
	}
}

// Chat 使用OpenAI SDK调用DeepSeek进行聊天
func (s *DeepSeekService) Chat(ctx context.Context, message string) (string, error) {
	resp, err := s.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: s.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response choices returned")
	}

	return resp.Choices[0].Message.Content, nil
}

// ChatWithOptions 支持更多参数的聊天方法
func (s *DeepSeekService) ChatWithOptions(ctx context.Context, messages []openai.ChatCompletionMessage, options *ChatOptions) (string, error) {
	req := openai.ChatCompletionRequest{
		Model:    s.model,
		Messages: messages,
	}

	// 应用选项
	if options != nil {
		if options.MaxTokens > 0 {
			req.MaxTokens = options.MaxTokens
		}
		if options.Temperature >= 0 {
			req.Temperature = options.Temperature
		}
		if options.TopP >= 0 {
			req.TopP = options.TopP
		}
		if options.Stream {
			req.Stream = options.Stream
		}
	}

	resp, err := s.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response choices returned")
	}

	return resp.Choices[0].Message.Content, nil
}

// ChatOptions 聊天选项
type ChatOptions struct {
	MaxTokens   int     `json:"max_tokens,omitempty"`
	Temperature float32 `json:"temperature,omitempty"`
	TopP        float32 `json:"top_p,omitempty"`
	Stream      bool    `json:"stream,omitempty"`
}

// 使用示例
func ExampleUsage() {
	// 创建DeepSeek服务实例
	service := NewDeepSeekServiceWithOpenAI("your-api-key", "deepseek-chat")

	ctx := context.Background()

	// 简单聊天
	response, err := service.Chat(ctx, "解释一下什么是人工智能")
	if err != nil {
		log.Printf("Chat error: %v", err)
		return
	}
	fmt.Printf("回复: %s\n", response)

	// 带选项的聊天
	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: "你是一个专业的技术助手"},
		{Role: openai.ChatMessageRoleUser, Content: "请解释Go语言的并发模型"},
	}

	options := &ChatOptions{
		MaxTokens:   500,
		Temperature: 0.7,
		TopP:        0.9,
	}

	response, err = service.ChatWithOptions(ctx, messages, options)
	if err != nil {
		log.Printf("ChatWithOptions error: %v", err)
		return
	}
	fmt.Printf("详细回复: %s\n", response)
}