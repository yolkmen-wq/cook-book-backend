package middlewares

import (
	"bytes"
	"cook-book-admin-backend/config"
	"cook-book-admin-backend/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"io"
	"net/http"
	"strings"
	"time"
)

type IPInfo struct {
	Data []map[string]interface{} `json:"data"`
}

type bodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

// 创建一个 channel 用于传递日志信息
var logChan = make(chan models.SystemLog)

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

var CommonLogInterceptor = commonLogInterceptor()

func init() {
	// 启动一个 goroutine 来处理日志记录
	go func() {
		for log := range logChan {
			err := config.DB.Create(&log).Error
			if err != nil {
				fmt.Println("admin log error", err)
			}
		}
	}()
}

/*
1 使用goroutine和channel实现操作日志的入库保存，尽可能的不影响主程序
2 goroutine协程，提高并发量
3 channel通道
*/
func commonLogInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminLogs(c)
	}
}

// 获得每次请求返回的code和message
func adminLogs(c *gin.Context) {
	// 解析请求参数

	method := c.Request.Method
	url := c.Request.URL.Path
	ip := c.Request.RemoteAddr
	if ip != "" {
		idx := strings.LastIndex(ip, ":")
		if idx > 0 {
			ip = ip[:idx]
		}
	}
	location, _ := getIPLocation(ip)
	userAgent := c.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	os := ua.OS()
	browser, _ := ua.Browser()

	// 记录请求日志
	var blw bodyLogWriter
	blw = bodyLogWriter{bodyBuf: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	start := time.Now()           // 开始时间
	c.Next()                      // 继续处理请求
	duration := time.Since(start) // 计算耗时

	// 异步记录日志
	systemLog := models.SystemLog{
		Module:      "admin",
		Url:         fmt.Sprintf("%s %s", method, url),
		Method:      method,
		Ip:          ip,
		Address:     location,
		System:      os,
		Browser:     browser,
		TakesTime:   duration.Milliseconds(),
		RequestTime: start,
	}
	logChan <- systemLog

}

func getIPLocation(ip string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://opendata.baidu.com/api.php?query=%s&co=&resource_id=6006&oe=utf8", ip))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var info IPInfo
	err = json.Unmarshal(body, &info)

	if err != nil {
		return "", err
	}
	if len(info.Data) == 0 {
		return "", errors.New("未找到IP地址信息")
	}
	location := info.Data[0]["location"]

	return location.(string), nil
}
