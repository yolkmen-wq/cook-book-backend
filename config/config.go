package config

// 定义一个密钥
var SecretKey = []byte("lyf123456")

// 定义统一的返回格式结构体
type Response struct {
	Code    int         `json:"code"`    // 状态码，0表示成功，非0表示错误
	Success bool        `json:"success"` // 是否成功
	Message string      `json:"message"` // 返回的消息
	Data    interface{} `json:"data"`    // 返回的数据，可以是任意类型
}

func NewResponse(code int, success bool, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Success: success,
		Message: message,
		Data:    data,
	}
}
