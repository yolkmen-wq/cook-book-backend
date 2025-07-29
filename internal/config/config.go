package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server   ServerConfig   `json:"server" yaml:"server"`
	Database DatabaseConfig `json:"database" yaml:"database"`
	JWT      JWTConfig      `json:"jwt" yaml:"jwt"`
	Wechat   WechatConfig   `json:"wechat" yaml:"wechat"`
	Logger   LoggerConfig   `json:"logger" yaml:"logger"` // 添加日志配置
}

type LoggerConfig struct {
	Level      string `json:"level" yaml:"level"`
	Format     string `json:"format" yaml:"format"`
	Output     string `json:"output" yaml:"output"`
	Filename   string `json:"filename" yaml:"filename"`
	MaxSize    int    `json:"maxSize" yaml:"maxSize"`
	MaxBackups int    `json:"maxBackups" yaml:"maxBackups"`
	MaxAge     int    `json:"maxAge" yaml:"maxAge"`
	// Compress   bool   `json:"compress" yaml:"compress"`
}

type ServerConfig struct {
	Port         string
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	SecretKey            string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

type WechatConfig struct {
	AppID     string
	AppSecret string
}

func Load() (*Config, error) {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "7575"),
			Host:         getEnv("SERVER_HOST", "localhost"),
			ReadTimeout:  getDurationEnv("SERVER_READ_TIMEOUT", 10*time.Second),
			WriteTimeout: getDurationEnv("SERVER_WRITE_TIMEOUT", 10*time.Second),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "47.121.201.137"),
			Port:     getIntEnv("DB_PORT", 3306),
			Username: getEnv("DB_USERNAME", "root"),
			Password: getEnv("DB_PASSWORD", "wuqi9457"),
			DBName:   getEnv("DB_NAME", "cook_book"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JWTConfig{
			SecretKey:            getEnv("JWT_SECRET", "your-secret-key"),
			AccessTokenDuration:  getDurationEnv("JWT_ACCESS_DURATION", 24*time.Hour),
			RefreshTokenDuration: getDurationEnv("JWT_REFRESH_DURATION", 24*time.Hour*7),
		},
		Logger: LoggerConfig{
			Level:      getEnv("LOG_LEVEL", "info"),
			Format:     getEnv("LOG_FORMAT", "json"),
			Output:     getEnv("LOG_OUTPUT", "stdout"),
			Filename:   getEnv("LOG_FILENAME", "logs/app.log"),
			MaxSize:    getIntEnv("LOG_MAX_SIZE", 100),
			MaxBackups: getIntEnv("LOG_MAX_BACKUPS", 3),
			MaxAge:     getIntEnv("LOG_MAX_AGE", 28),
			// Compress:   parseBoolEnv("LOG_COMPRESS", true),
		},
		Wechat: WechatConfig{
			AppID:     getEnv("WECHAT_APP_ID", "wx602e2e9a008deb2b"),
			AppSecret: getEnv("WECHAT_APP_SECRET", "65e3fe1540af9fedd4bbeb33de5b6558"),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
