package services

import (
	"context"
	"fmt"
	"io"

	"github.com/gorilla/websocket"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAIService struct {
	client *openai.Client
	model  string
}

func NewOpenAIService(apiKey, model string) OpenAIService {
	return OpenAIService{
		client: openai.NewClient(apiKey),
		model:  model,
	}
}

func (s *OpenAIService) Chat(ctx context.Context, message string) (string, error) {
	// 验证API密钥（通过检查model是否为空来间接验证配置）
	if s.model == "" {
		return "", fmt.Errorf("OpenAI API configuration is incomplete. Please set AI_API_KEY and AI_MODEL environment variables")
	}

	resp, err := s.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: s.model,
			Messages: []openai.ChatCompletionMessage{
				{Role: "user", Content: message},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

// ChatStream 流式聊天方法，返回一个channel用于接收流式响应
func (s *OpenAIService) ChatStream(ctx context.Context, message string) (<-chan string, <-chan error) {
	respChan := make(chan string, 10)
	errorChan := make(chan error, 1)

	go func() {
		defer close(respChan)
		defer close(errorChan)

		// 验证API密钥
		if s.model == "" {
			errorChan <- fmt.Errorf("OpenAI API configuration is incomplete. Please set AI_API_KEY and AI_MODEL environment variables")
			return
		}

		stream, err := s.client.CreateChatCompletionStream(
			ctx,
			openai.ChatCompletionRequest{
				Model: s.model,
				Messages: []openai.ChatCompletionMessage{
					{Role: "user", Content: message},
				},
				Stream: true,
			},
		)

		if err != nil {
			errorChan <- err
			return
		}
		defer stream.Close()

		for {
			response, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					// 流结束
					return
				}
				errorChan <- err
				return
			}

			if len(response.Choices) > 0 {
				content := response.Choices[0].Delta.Content
				if content != "" {
					select {
					case respChan <- content:
					case <-ctx.Done():
						errorChan <- ctx.Err()
						return
					}
				}
			}
		}
	}()

	return respChan, errorChan
}

// ChatWebSocket 通过WebSocket进行流式聊天
func (s *OpenAIService) ChatWebSocket(conn *websocket.Conn, message string) {
	if s.model == "" {
		responseError(conn, "OpenAI API configuration is incomplete. Please set AI_API_KEY and AI_MODEL environment variables")
		return
	}

	// 创建上下文
	ctx := context.Background()

	// 创建流式请求
	stream, err := s.client.CreateChatCompletionStream(
		ctx,
		openai.ChatCompletionRequest{
			Model: s.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
			Stream: true,
		},
	)

	if err != nil {
		responseError(conn, err.Error())
		return
	}
	defer stream.Close()

	// 处理流式响应
	for {
		select {
		case <-ctx.Done():
			// 上下文取消或超时
			responseError(conn, "request timeout or canceled")
			return
		default:
			// 接收流式响应
			response, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					// 流结束
					return
				}
				// 发生错误
				responseError(conn, err.Error())
				return
			}

			// 处理响应内容
			if len(response.Choices) > 0 {
				content := response.Choices[0].Delta.Content
				if content != "" {
					// 发送内容到WebSocket
					if err := sendWSMessage(conn, "message", content); err != nil {
						// WebSocket发送失败
						return
					}
				}
			}
		}
	}
}
