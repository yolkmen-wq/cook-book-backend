package logger

import (
    "bytes"
    "io"
    "time"
    
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

// responseWriter 包装 gin.ResponseWriter 以捕获响应体
type responseWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
    w.body.Write(b)
    return w.ResponseWriter.Write(b)
}

// GinLogger 返回 Gin 日志中间件
func GinLogger(logger Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 生成请求ID
        requestID := uuid.New().String()
        c.Set("request_id", requestID)
        
        // 记录请求开始时间
        start := time.Now()
        
        // 读取请求体（如果需要记录）
        var requestBody []byte
        if c.Request.Body != nil {
            requestBody, _ = io.ReadAll(c.Request.Body)
            c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
        }
        
        // 包装响应写入器
        writer := &responseWriter{
            ResponseWriter: c.Writer,
            body:          bytes.NewBufferString(""),
        }
        c.Writer = writer
        
        // 处理请求
        c.Next()
        
        // 计算处理时间
        duration := time.Since(start)
        
        // 构建日志字段
        fields := Fields{
            "request_id":     requestID,
            "method":         c.Request.Method,
            "path":           c.Request.URL.Path,
            "query":          c.Request.URL.RawQuery,
            "status":         c.Writer.Status(),
            "duration_ms":    duration.Milliseconds(),
            "client_ip":      c.ClientIP(),
            "user_agent":     c.Request.UserAgent(),
            "request_size":   len(requestBody),
            "response_size":  writer.body.Len(),
        }
        
        // 添加用户ID（如果存在）
        if userID, exists := c.Get("user_id"); exists {
            fields["user_id"] = userID
        }
        
        // 根据状态码选择日志级别
        logLevel := InfoLevel
        if c.Writer.Status() >= 400 && c.Writer.Status() < 500 {
            logLevel = WarnLevel
        } else if c.Writer.Status() >= 500 {
            logLevel = ErrorLevel
        }
        
        // 记录日志
        loggerWithFields := logger.WithFields(fields)
        
        switch logLevel {
        case WarnLevel:
            loggerWithFields.Warn("HTTP Request")
        case ErrorLevel:
            loggerWithFields.Error("HTTP Request")
        default:
            loggerWithFields.Info("HTTP Request")
        }
        
        // 如果有错误，记录详细信息
        if len(c.Errors) > 0 {
            for _, err := range c.Errors {
                logger.WithFields(Fields{
                    "request_id": requestID,
                    "error_type": err.Type,
                }).Error("Request error: " + err.Error())
            }
        }
    }
}

// RequestIDMiddleware 请求ID中间件
func RequestIDMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestID := c.GetHeader("X-Request-ID")
        if requestID == "" {
            requestID = uuid.New().String()
        }
        
        c.Set("request_id", requestID)
        c.Header("X-Request-ID", requestID)
        c.Next()
    }
}