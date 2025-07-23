package logger

import (
    "context"
    "fmt"
    "runtime"
    "time"
)

// LogExecutionTime 记录函数执行时间
func LogExecutionTime(logger Logger, operation string) func() {
    start := time.Now()
    return func() {
        duration := time.Since(start)
        logger.Info("Operation completed",
            "operation", operation,
            "duration_ms", duration.Milliseconds(),
        )
    }
}

// LogPanic 记录 panic 信息
func LogPanic(logger Logger) {
    if r := recover(); r != nil {
        // 获取调用栈
        buf := make([]byte, 4096)
        n := runtime.Stack(buf, false)
        
        logger.Error("Panic recovered",
            "panic", r,
            "stack", string(buf[:n]),
        )
        
        // 重新抛出 panic
        panic(r)
    }
}

// LogSlowQuery 记录慢查询
func LogSlowQuery(logger Logger, threshold time.Duration) func(query string, duration time.Duration) {
    return func(query string, duration time.Duration) {
        if duration > threshold {
            logger.Warn("Slow query detected",
                "query", query,
                "duration_ms", duration.Milliseconds(),
                "threshold_ms", threshold.Milliseconds(),
            )
        }
    }
}

// ContextWithLogger 将日志器添加到上下文
func ContextWithLogger(ctx context.Context, logger Logger) context.Context {
    return context.WithValue(ctx, "logger", logger)
}

// LoggerFromContext 从上下文获取日志器
func LoggerFromContext(ctx context.Context) Logger {
    if logger, ok := ctx.Value("logger").(Logger); ok {
        return logger
    }
    return GetGlobalLogger()
}

// LogError 便捷的错误日志记录
func LogError(logger Logger, err error, operation string, fields ...interface{}) {
    if err != nil {
        allFields := []interface{}{"operation", operation, "error", err.Error()}
        allFields = append(allFields, fields...)
        logger.Error("Operation failed", allFields...)
    }
}

// LogSuccess 便捷的成功日志记录
func LogSuccess(logger Logger, operation string, fields ...interface{}) {
    allFields := []interface{}{"operation", operation}
    allFields = append(allFields, fields...)
    logger.Info("Operation succeeded", allFields...)
}

// FormatError 格式化错误信息
func FormatError(err error, context string) string {
    if err == nil {
        return ""
    }
    return fmt.Sprintf("%s: %v", context, err)
}