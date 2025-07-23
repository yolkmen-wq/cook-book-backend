package logger

import (
	"context"
)

// Logger 定义日志接口
type Logger interface {
	// 基础日志方法
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})

	// 带上下文的日志方法
	DebugContext(ctx context.Context, msg string, fields ...interface{})
	InfoContext(ctx context.Context, msg string, fields ...interface{})
	WarnContext(ctx context.Context, msg string, fields ...interface{})
	ErrorContext(ctx context.Context, msg string, fields ...interface{})

	// 结构化日志方法
	WithFields(fields Fields) Logger
	WithField(key string, value interface{}) Logger
	WithError(err error) Logger
	WithContext(ctx context.Context) Logger

	// 获取子日志器
	WithComponent(component string) Logger
}

// Fields 定义日志字段类型
type Fields map[string]interface{}

// Level 定义日志级别
type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// String 返回日志级别的字符串表示
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	default:
		return "unknown"
	}
}

// ParseLevel 从字符串解析日志级别
func ParseLevel(level string) Level {
	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn", "warning":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	default:
		return InfoLevel
	}
}

// Config 日志配置
type Config struct {
	Level      string `json:"level" yaml:"level"`           // 日志级别
	Format     string `json:"format" yaml:"format"`         // 日志格式 (json/text)
	Output     string `json:"output" yaml:"output"`         // 输出目标 (stdout/file)
	Filename   string `json:"filename" yaml:"filename"`     // 日志文件名
	MaxSize    int    `json:"maxSize" yaml:"maxSize"`       // 单个日志文件最大大小(MB)
	MaxBackups int    `json:"maxBackups" yaml:"maxBackups"` // 保留的旧日志文件数量
	MaxAge     int    `json:"maxAge" yaml:"maxAge"`         // 保留旧日志文件的最大天数
	Compress   bool   `json:"compress" yaml:"compress"`     // 是否压缩旧日志文件
}

// DefaultConfig 返回默认配置
func DefaultConfig() Config {
	return Config{
		Level:      "info",
		Format:     "json",
		Output:     "stdout",
		Filename:   "logs/app.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}
}
