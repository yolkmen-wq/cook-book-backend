package services

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/gorilla/websocket"
	openai "github.com/sashabaranov/go-openai"
)

type DeepSeekService struct {
	client *openai.Client
	model  string
}

func NewDeepSeekService(apiKey, model string) DeepSeekService {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.deepseek.com/v1"
	return DeepSeekService{
		client: openai.NewClientWithConfig(config),
		model:  model,
	}
}

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

// ChatStream 流式聊天方法，返回一个channel用于接收流式响应
func (s *DeepSeekService) ChatStream(ctx context.Context, message string) (<-chan string, <-chan error) {
	respChan := make(chan string, 10)
	errorChan := make(chan error, 1)

	go func() {
		defer close(respChan)
		defer close(errorChan)

		// 验证API密钥
		if s.model == "" {
			errorChan <- fmt.Errorf("AI_API_KEY and AI_MODEL environment variables must be set")
			return
		}

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
func (s *DeepSeekService) ChatWebSocket(conn *websocket.Conn, message string) {
	// 创建上下文，可以通过超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// 验证API密钥
	if s.model == "" {
		responseError(conn, "AI_API_KEY and AI_MODEL environment variables must be set")
		return
	}

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
					// 流结束，发送完成消息
					sendWSMessage(conn, "done", "")
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

// 发送WebSocket错误消息
func responseError(conn *websocket.Conn, errMsg string) {
	sendWSMessage(conn, "error", errMsg)
}

// 发送WebSocket消息
func sendWSMessage(conn *websocket.Conn, msgType string, content string) error {
	msg := map[string]string{
		"type":    msgType,
		"content": content,
	}
	return conn.WriteJSON(msg)
}
