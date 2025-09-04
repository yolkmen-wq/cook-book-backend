package services

import (
	"context"

	"github.com/gorilla/websocket"
)

type AIService interface {
	Chat(ctx context.Context, message string) (string, error)
	ChatStream(ctx context.Context, message string) (<-chan string, <-chan error)
	ChatWebSocket(conn *websocket.Conn, message string)
}
