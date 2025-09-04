# 使用OpenAI SDK调用DeepSeek API

本文档介绍如何在Golang中使用OpenAI SDK来调用DeepSeek API。DeepSeek提供了与OpenAI兼容的API接口，因此可以直接使用OpenAI的SDK。

## 为什么使用OpenAI SDK调用DeepSeek？

1. **API兼容性**：DeepSeek API与OpenAI API完全兼容
2. **成熟的SDK**：OpenAI SDK功能完善，社区支持好
3. **统一接口**：可以轻松在不同AI提供商之间切换
4. **丰富功能**：支持流式响应、函数调用等高级功能

## 安装依赖

```bash
go get github.com/sashabaranov/go-openai
```

## 基本配置

### 方法1：直接配置BaseURL

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    openai "github.com/sashabaranov/go-openai"
)

func main() {
    // 创建配置，指向DeepSeek API
    config := openai.DefaultConfig("your-deepseek-api-key")
    config.BaseURL = "https://api.deepseek.com/v1"
    
    // 创建客户端
    client := openai.NewClientWithConfig(config)
    
    // 发送请求
    resp, err := client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: "deepseek-chat",
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    openai.ChatMessageRoleUser,
                    Content: "你好，DeepSeek！",
                },
            },
        },
    )
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(resp.Choices[0].Message.Content)
}
```

### 方法2：封装服务类

参考项目中的 `examples/deepseek_with_openai_sdk.go` 文件：

```go
type DeepSeekService struct {
    client *openai.Client
    model  string
}

func NewDeepSeekServiceWithOpenAI(apiKey, model string) *DeepSeekService {
    config := openai.DefaultConfig(apiKey)
    config.BaseURL = "https://api.deepseek.com/v1"
    
    return &DeepSeekService{
        client: openai.NewClientWithConfig(config),
        model:  model,
    }
}
```

## 环境变量配置

在 `.env` 文件中配置：

```bash
# DeepSeek配置
AI_PROVIDER=deepseek
AI_API_KEY=your-deepseek-api-key
AI_MODEL=deepseek-chat
AI_BASE_URL=https://api.deepseek.com/v1
```

## 支持的功能

### 1. 基本聊天

```go
resp, err := client.CreateChatCompletion(
    ctx,
    openai.ChatCompletionRequest{
        Model: "deepseek-chat",
        Messages: []openai.ChatCompletionMessage{
            {Role: openai.ChatMessageRoleUser, Content: "Hello"},
        },
    },
)
```

### 2. 系统提示词

```go
messages := []openai.ChatCompletionMessage{
    {Role: openai.ChatMessageRoleSystem, Content: "你是一个专业的编程助手"},
    {Role: openai.ChatMessageRoleUser, Content: "解释Go语言的协程"},
}
```

### 3. 参数控制

```go
req := openai.ChatCompletionRequest{
    Model:       "deepseek-chat",
    Messages:    messages,
    MaxTokens:   1000,
    Temperature: 0.7,
    TopP:        0.9,
}
```

### 4. 流式响应

```go
req := openai.ChatCompletionRequest{
    Model:    "deepseek-chat",
    Messages: messages,
    Stream:   true,
}

stream, err := client.CreateChatCompletionStream(ctx, req)
if err != nil {
    return err
}
defer stream.Close()

for {
    response, err := stream.Recv()
    if errors.Is(err, io.EOF) {
        break
    }
    if err != nil {
        return err
    }
    
    fmt.Print(response.Choices[0].Delta.Content)
}
```

## DeepSeek支持的模型

- `deepseek-chat` - 主要的对话模型
- `deepseek-coder` - 专门用于代码生成的模型

## 错误处理

```go
resp, err := client.CreateChatCompletion(ctx, req)
if err != nil {
    // 检查是否是API错误
    var apiErr *openai.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API Error: %s (Code: %s)\n", apiErr.Message, apiErr.Code)
        return
    }
    
    // 其他错误
    fmt.Printf("Request failed: %v\n", err)
    return
}
```

## 与项目集成

在您的项目中，可以修改 `container.go` 来支持使用OpenAI SDK：

```go
// 在container.go中添加
func (c *Container) NewDeepSeekWithOpenAI() *DeepSeekService {
    config := openai.DefaultConfig(c.cfg.AI.APIKey)
    config.BaseURL = "https://api.deepseek.com/v1"
    
    return &DeepSeekService{
        client: openai.NewClientWithConfig(config),
        model:  c.cfg.AI.Model,
    }
}
```

## 性能优化建议

1. **连接复用**：重用client实例，避免频繁创建
2. **超时设置**：设置合适的请求超时时间
3. **并发控制**：使用context控制并发请求
4. **错误重试**：实现指数退避重试机制

## 注意事项

1. **API密钥安全**：不要在代码中硬编码API密钥
2. **速率限制**：注意DeepSeek的API调用频率限制
3. **模型选择**：根据使用场景选择合适的模型
4. **成本控制**：合理设置MaxTokens参数控制成本

## 完整示例

运行示例代码：

```bash
cd examples
go run deepseek_with_openai_sdk.go
```

确保设置了正确的API密钥环境变量或在代码中配置。