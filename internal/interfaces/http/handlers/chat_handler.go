package handlers

import (
	"context"
	"cook-book-backend/internal/domain/services"
	"cook-book-backend/internal/infrastructure/logger"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// chatStreamWithWebSocket 处理WebSocket流式聊天
func (h *ChatHandler) chatStreamWithWebSocket(ctx context.Context, conn *websocket.Conn, aiService interface{}, message string) {
	// 根据AI服务类型调用相应的ChatStream方法
	var respChan <-chan string
	var errChan <-chan error

	// 类型断言以确定具体的AI服务
	switch service := aiService.(type) {
	case services.OpenAIService:
		respChan, errChan = service.ChatStream(ctx, message)
	case services.DeepSeekService:
		respChan, errChan = service.ChatStream(ctx, message)
	default:
		h.logger.Error("Unknown AI service type")
		conn.WriteJSON(map[string]string{
			"type":    "error",
			"content": "未知的AI服务类型",
		})
		return
	}

	// 处理流式响应
	for {
		select {
		case <-ctx.Done():
			// 上下文被取消，发送停止消息并退出
			h.logger.Info("Chat stream stopped by context cancellation")
			conn.WriteJSON(map[string]string{
				"type":    "stop",
				"content": "聊天已停止",
			})
			return
		case resp, ok := <-respChan:
			if !ok {
				// 响应channel关闭，流式响应结束
				h.logger.Info("Chat stream completed")
				conn.WriteJSON(map[string]string{
					"type":    "end",
					"content": "聊天结束",
				})
				return
			}
			// 发送响应到客户端
			if err := conn.WriteJSON(map[string]string{
				"type":    "message",
				"content": resp,
			}); err != nil {
				h.logger.Error("Failed to write message to WebSocket: " + err.Error())
				return
			}
		case err, ok := <-errChan:
			if !ok {
				// 错误channel关闭，流式响应结束
				return
			}
			// 发送错误到客户端
			h.logger.Error("Chat stream error: " + err.Error())
			conn.WriteJSON(map[string]string{
				"type":    "error",
				"content": err.Error(),
			})
			return
		}
	}
}

type ChatHandler struct {
	*BaseHandler
	openAIService   services.OpenAIService
	deepSeekService services.DeepSeekService
	upgrader       websocket.Upgrader
}

func NewChatHandler(OpenAIService services.OpenAIService, DeepSeekService services.DeepSeekService, logger logger.Logger) *ChatHandler {
	return &ChatHandler{
		BaseHandler:     NewBaseHandler(logger),
		openAIService:   OpenAIService,
		deepSeekService: DeepSeekService,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// 允许所有跨域请求，生产环境中应该限制来源
				return true
			},
		},
	}
}

// Chat 聊天
func (h *ChatHandler) Chat(c *gin.Context) {
	type Req struct {
		Message  string `json:"message"`
		Provider string `json:"provider"` // 可选参数，指定使用的AI提供商
	}
	type Res struct {
		Reply string `json:"reply"`
	}

	var req Req
	if err := c.Bind(&req); err != nil {
		h.HandleError(c, err, "参数错误")
		return
	}

	// 根据请求参数或默认配置选择AI服务
	var reply string
	var err error
	ctx := context.Background()

	// 如果指定了provider参数，则使用指定的服务
	switch req.Provider {
	case "openai":
		reply, err = h.openAIService.Chat(ctx, req.Message)
	case "deepseek":
		reply, err = h.deepSeekService.Chat(ctx, req.Message)
	default:
		// 默认使用DeepSeek服务
		reply, err = h.deepSeekService.Chat(ctx, req.Message)
	}

	if err != nil {
		h.HandleError(c, err, "chat_error")
		return
	}

	h.Success(c, reply)
}

// ChatStream 流式聊天（SSE方式）
func (h *ChatHandler) ChatStream(c *gin.Context) {
	type Req struct {
		Message  string `json:"message"`
		Provider string `json:"provider"` // 可选参数，指定使用的AI提供商
	}
	type Res struct {
		Reply string `json:"reply"`
	}

	var req Req
	if err := c.Bind(&req); err != nil {
		h.HandleError(c, err, "参数错误")
		return
	}

	// 根据请求参数选择AI服务
	var respChan <-chan string
	var errChan <-chan error
	ctx := context.Background()

	// 如果指定了provider参数，则使用指定的服务
	switch req.Provider {
	case "openai":
		respChan, errChan = h.openAIService.ChatStream(ctx, req.Message)
	case "deepseek":
		respChan, errChan = h.deepSeekService.ChatStream(ctx, req.Message)
	default:
		// 默认使用DeepSeek服务
		respChan, errChan = h.deepSeekService.ChatStream(ctx, req.Message)
	}
	// 处理流式响应
	for {
		select {
		case resp, ok := <-respChan:
			if !ok {
				// 响应channel关闭，流式响应结束
				return
			}
			// 发送响应到客户端
			c.SSEvent("message", Res{Reply: resp})
			c.Writer.Flush()
		case err, ok := <-errChan:
			if !ok {
				// 错误channel关闭，流式响应结束
				return
			}
			// 发送错误到客户端
			c.SSEvent("error", err.Error())
			c.Writer.Flush()
			return
		}
	}
}

// ChatWebSocket 通过WebSocket进行流式聊天
func (h *ChatHandler) ChatWebSocket(c *gin.Context) {
	// 将HTTP连接升级为WebSocket连接
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		h.logger.Error("Failed to upgrade to websocket: " + err.Error())
		return
	}
	defer conn.Close()

	// 定义消息结构
	type WSMessage struct {
		Message  string `json:"message"`
		Provider string `json:"provider"` // 可选参数，指定使用的AI提供商
		Action   string `json:"action"`   // 操作类型："chat" 或 "stop"
	}

	// 读取客户端发送的第一条消息
	messageType, messageData, err := conn.ReadMessage()
	if err != nil {
		h.logger.Error("Failed to read raw message: " + err.Error())
		conn.WriteJSON(map[string]string{
			"type":    "error",
			"content": "无法读取消息: " + err.Error(),
		})
		return
	}

	// 记录原始消息类型和内容
	h.logger.Info(fmt.Sprintf("Received WebSocket message type: %d", messageType))
	h.logger.Debug("Raw message data: " + string(messageData))

	// 尝试解析为JSON
	var msg WSMessage
	var messageText string
	var provider string

	// 尝试解析JSON
	if err := json.Unmarshal(messageData, &msg); err != nil {
		// 如果不是JSON格式，则将整个消息作为文本处理
		h.logger.Warn("Message is not valid JSON, treating as plain text: " + err.Error())
		messageText = string(messageData)
		provider = "" // 使用默认提供商
	} else {
		// 成功解析JSON
		messageText = msg.Message
		provider = msg.Provider
	}

	// 验证消息不为空
	if messageText == "" {
		h.logger.Error("Empty message received")
		conn.WriteJSON(map[string]string{
			"type":    "error",
			"content": "消息内容不能为空",
		})
		return
	}

	// 创建可取消的上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动goroutine监听后续WebSocket消息（用于停止功能）
	go func() {
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				// WebSocket连接关闭或出错，取消上下文
				cancel()
				return
			}

			// 尝试解析停止消息
			var stopMsg WSMessage
			if err := json.Unmarshal(data, &stopMsg); err == nil {
				if stopMsg.Action == "stop" {
					h.logger.Info("Received stop command, canceling chat stream")
					cancel()
					return
				}
			}
		}
	}()

	// 根据请求参数选择AI服务并进行流式聊天
	switch provider {
	case "openai":
		h.logger.Info("Using OpenAI service for WebSocket chat")
		h.chatStreamWithWebSocket(ctx, conn, h.openAIService, messageText)
	case "deepseek":
		h.logger.Info("Using DeepSeek service for WebSocket chat")
		h.chatStreamWithWebSocket(ctx, conn, h.deepSeekService, messageText)
	default:
		// 默认使用DeepSeek服务
		h.logger.Info("Using default DeepSeek service for WebSocket chat")
		h.chatStreamWithWebSocket(ctx, conn, h.deepSeekService, messageText)
	}
}
