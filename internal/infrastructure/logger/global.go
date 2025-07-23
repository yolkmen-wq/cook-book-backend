package logger

import (
    "context"
    "sync"
)

var (
    globalLogger Logger
    once         sync.Once
)

// SetGlobalLogger 设置全局日志器
func SetGlobalLogger(logger Logger) {
    globalLogger = logger
}

// GetGlobalLogger 获取全局日志器
func GetGlobalLogger() Logger {
    once.Do(func() {
        if globalLogger == nil {
            globalLogger = New(DefaultConfig())
        }
    })
    return globalLogger
}

// 全局日志方法
func Debug(msg string, fields ...interface{}) {
    GetGlobalLogger().Debug(msg, fields...)
}

func Info(msg string, fields ...interface{}) {
    GetGlobalLogger().Info(msg, fields...)
}

func Warn(msg string, fields ...interface{}) {
    GetGlobalLogger().Warn(msg, fields...)
}

func Error(msg string, fields ...interface{}) {
    GetGlobalLogger().Error(msg, fields...)
}

func Fatal(msg string, fields ...interface{}) {
    GetGlobalLogger().Fatal(msg, fields...)
}

// 带上下文的全局日志方法
func DebugContext(ctx context.Context, msg string, fields ...interface{}) {
    GetGlobalLogger().DebugContext(ctx, msg, fields...)
}

func InfoContext(ctx context.Context, msg string, fields ...interface{}) {
    GetGlobalLogger().InfoContext(ctx, msg, fields...)
}

func WarnContext(ctx context.Context, msg string, fields ...interface{}) {
    GetGlobalLogger().WarnContext(ctx, msg, fields...)
}

func ErrorContext(ctx context.Context, msg string, fields ...interface{}) {
    GetGlobalLogger().ErrorContext(ctx, msg, fields...)
}

// 便捷方法
func WithFields(fields Fields) Logger {
    return GetGlobalLogger().WithFields(fields)
}

func WithField(key string, value interface{}) Logger {
    return GetGlobalLogger().WithField(key, value)
}

func WithError(err error) Logger {
    return GetGlobalLogger().WithError(err)
}

func WithContext(ctx context.Context) Logger {
    return GetGlobalLogger().WithContext(ctx)
}

func WithComponent(component string) Logger {
    return GetGlobalLogger().WithComponent(component)
}