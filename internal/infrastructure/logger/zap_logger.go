package logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// zapLogger 基于 zap 的日志实现
type zapLogger struct {
	logger    *zap.Logger
	sugar     *zap.SugaredLogger
	level     Level
	component string
	fields    Fields
}

// New 创建新的日志器
func New(config Config) Logger {
	// 创建日志目录
	if config.Output == "file" && config.Filename != "" {
		dir := filepath.Dir(config.Filename)
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(fmt.Sprintf("failed to create log directory: %v", err))
		}
	}

	// 配置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 选择编码器
	var encoder zapcore.Encoder
	if config.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 配置输出
	var writeSyncer zapcore.WriteSyncer
	if config.Output == "file" {
		// 文件输出，支持日志轮转
		lumberJackLogger := &lumberjack.Logger{
			Filename:   config.Filename,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		}
		writeSyncer = zapcore.AddSync(lumberJackLogger)
	} else {
		// 控制台输出
		writeSyncer = zapcore.AddSync(os.Stdout)
	}

	// 设置日志级别
	level := ParseLevel(config.Level)
	zapLevel := toZapLevel(level)

	// 创建核心
	core := zapcore.NewCore(encoder, writeSyncer, zapLevel)

	// 创建 logger
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &zapLogger{
		logger: logger,
		sugar:  logger.Sugar(),
		level:  level,
		fields: make(Fields),
	}
}

// Debug 记录调试级别日志
func (l *zapLogger) Debug(msg string, fields ...interface{}) {
	l.log(DebugLevel, msg, fields...)
}

// Info 记录信息级别日志
func (l *zapLogger) Info(msg string, fields ...interface{}) {
	l.log(InfoLevel, msg, fields...)
}

// Warn 记录警告级别日志
func (l *zapLogger) Warn(msg string, fields ...interface{}) {
	l.log(WarnLevel, msg, fields...)
}

// Error 记录错误级别日志
func (l *zapLogger) Error(msg string, fields ...interface{}) {
	l.log(ErrorLevel, msg, fields...)
}

// Fatal 记录致命错误级别日志
func (l *zapLogger) Fatal(msg string, fields ...interface{}) {
	l.log(FatalLevel, msg, fields...)
}

// DebugContext 带上下文的调试日志
func (l *zapLogger) DebugContext(ctx context.Context, msg string, fields ...interface{}) {
	l.logWithContext(ctx, DebugLevel, msg, fields...)
}

// InfoContext 带上下文的信息日志
func (l *zapLogger) InfoContext(ctx context.Context, msg string, fields ...interface{}) {
	l.logWithContext(ctx, InfoLevel, msg, fields...)
}

// WarnContext 带上下文的警告日志
func (l *zapLogger) WarnContext(ctx context.Context, msg string, fields ...interface{}) {
	l.logWithContext(ctx, WarnLevel, msg, fields...)
}

// ErrorContext 带上下文的错误日志
func (l *zapLogger) ErrorContext(ctx context.Context, msg string, fields ...interface{}) {
	l.logWithContext(ctx, ErrorLevel, msg, fields...)
}

// WithFields 添加多个字段
func (l *zapLogger) WithFields(fields Fields) Logger {
	newFields := make(Fields)
	for k, v := range l.fields {
		newFields[k] = v
	}
	for k, v := range fields {
		newFields[k] = v
	}

	return &zapLogger{
		logger:    l.logger,
		sugar:     l.sugar,
		level:     l.level,
		component: l.component,
		fields:    newFields,
	}
}

// WithField 添加单个字段
func (l *zapLogger) WithField(key string, value interface{}) Logger {
	return l.WithFields(Fields{key: value})
}

// WithError 添加错误字段
func (l *zapLogger) WithError(err error) Logger {
	return l.WithField("error", err.Error())
}

// WithContext 添加上下文信息
func (l *zapLogger) WithContext(ctx context.Context) Logger {
	fields := make(Fields)

	// 从上下文中提取常用字段
	if requestID := ctx.Value("request_id"); requestID != nil {
		fields["request_id"] = requestID
	}
	if userID := ctx.Value("user_id"); userID != nil {
		fields["user_id"] = userID
	}
	if traceID := ctx.Value("trace_id"); traceID != nil {
		fields["trace_id"] = traceID
	}

	return l.WithFields(fields)
}

// WithComponent 添加组件名称
func (l *zapLogger) WithComponent(component string) Logger {
	return &zapLogger{
		logger:    l.logger,
		sugar:     l.sugar,
		level:     l.level,
		component: component,
		fields:    l.fields,
	}
}

// log 内部日志方法
func (l *zapLogger) log(level Level, msg string, fields ...interface{}) {
	if level < l.level {
		return
	}

	zapFields := l.buildZapFields(fields...)

	switch level {
	case DebugLevel:
		l.logger.Debug(msg, zapFields...)
	case InfoLevel:
		l.logger.Info(msg, zapFields...)
	case WarnLevel:
		l.logger.Warn(msg, zapFields...)
	case ErrorLevel:
		l.logger.Error(msg, zapFields...)
	case FatalLevel:
		l.logger.Fatal(msg, zapFields...)
	}
}

// logWithContext 带上下文的内部日志方法
func (l *zapLogger) logWithContext(ctx context.Context, level Level, msg string, fields ...interface{}) {
	contextLogger := l.WithContext(ctx)
	contextLogger.(*zapLogger).log(level, msg, fields...)
}

// buildZapFields 构建 zap 字段
func (l *zapLogger) buildZapFields(fields ...interface{}) []zap.Field {
	var zapFields []zap.Field

	// 添加组件字段
	if l.component != "" {
		zapFields = append(zapFields, zap.String("component", l.component))
	}

	// 添加预设字段
	for k, v := range l.fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}

	// 添加传入的字段
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			key := fmt.Sprintf("%v", fields[i])
			value := fields[i+1]
			zapFields = append(zapFields, zap.Any(key, value))
		}
	}

	return zapFields
}

// toZapLevel 转换为 zap 日志级别
func toZapLevel(level Level) zapcore.Level {
	switch level {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
