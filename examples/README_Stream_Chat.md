# 流式聊天功能使用指南

本文档介绍如何使用DeepSeek和OpenAI服务的流式聊天功能。

## 功能特点

- **实时响应**: 无需等待完整回复，实时接收AI生成的内容
- **更好的用户体验**: 类似ChatGPT的打字机效果
- **支持取消**: 可以通过context取消正在进行的流式请求
- **错误处理**: 完善的错误处理机制
- **统一接口**: DeepSeek和OpenAI服务都实现了相同的流式接口

## 接口定义

```go
type AIService interface {
    Chat(ctx context.Context, message string) (string, error)
    ChatStream(ctx context.Context, message string) (<-chan string, <-chan error)
}
```

## 基本使用方法

### 1. 创建服务实例

```go
// DeepSeek服务
deepseekService := services.NewDeepSeekService(apiKey, model)

// OpenAI服务
openaiService := services.NewOpenAIService(apiKey, model)
```

### 2. 流式聊天

```go
ctx := context.Background()
respChan, errorChan := aiService.ChatStream(ctx, "你好，请介绍一下自己")

for {
    select {
    case content, ok := <-respChan:
        if !ok {
            // 响应完成
            fmt.Println("\n流式响应完成")
            return
        }
        // 实时打印内容
        fmt.Print(content)
    case err := <-errorChan:
        if err != nil {
            fmt.Printf("错误: %v\n", err)
            return
        }
    }
}
```

## 高级用法

### 1. 带超时的流式聊天

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

respChan, errorChan := aiService.ChatStream(ctx, message)

for {
    select {
    case content, ok := <-respChan:
        if !ok {
            return
        }
        fmt.Print(content)
    case err := <-errorChan:
        if err != nil {
            log.Printf("流式聊天错误: %v", err)
            return
        }
    case <-ctx.Done():
        log.Printf("请求超时: %v", ctx.Err())
        return
    }
}
```

### 2. 可取消的流式聊天

```go
ctx, cancel := context.WithCancel(context.Background())

// 在另一个goroutine中处理用户取消操作
go func() {
    // 监听用户输入或其他取消信号
    // ...
    cancel() // 取消请求
}()

respChan, errorChan := aiService.ChatStream(ctx, message)
// 处理响应...
```

### 3. 累积响应内容

```go
var fullResponse strings.Builder

respChan, errorChan := aiService.ChatStream(ctx, message)

for {
    select {
    case content, ok := <-respChan:
        if !ok {
            // 获取完整响应
            result := fullResponse.String()
            fmt.Printf("完整响应: %s\n", result)
            return
        }
        // 实时显示并累积内容
        fmt.Print(content)
        fullResponse.WriteString(content)
    case err := <-errorChan:
        if err != nil {
            log.Printf("错误: %v", err)
            return
        }
    }
}
```

## 环境变量配置

确保设置以下环境变量：

```bash
# DeepSeek配置
export AI_PROVIDER=deepseek
export AI_API_KEY=your_deepseek_api_key
export AI_MODEL=deepseek-chat

# 或者OpenAI配置
export AI_PROVIDER=openai
export AI_API_KEY=your_openai_api_key
export AI_MODEL=gpt-3.5-turbo
```

## 运行示例

```bash
# 设置环境变量
export AI_API_KEY=your_api_key
export AI_MODEL=deepseek-chat
export AI_PROVIDER=deepseek

# 运行示例
go run examples/stream_chat_example.go
```

## 错误处理

流式聊天可能遇到的错误：

1. **API密钥未配置**: 确保设置了`AI_API_KEY`和`AI_MODEL`环境变量
2. **网络错误**: 检查网络连接和API服务状态
3. **API配额不足**: 检查API账户余额
4. **请求超时**: 适当调整context超时时间
5. **上下文取消**: 正常的取消操作，不是错误

## 性能优化建议

1. **合理设置缓冲区大小**: 响应channel默认缓冲区为10，可根据需要调整
2. **及时处理响应**: 避免channel阻塞
3. **适当的超时设置**: 根据预期响应时间设置合理的超时
4. **资源清理**: 确保在不需要时取消context，释放资源

## 注意事项

1. 流式响应是异步的，需要在goroutine中处理
2. 响应channel关闭表示流式响应结束
3. 错误channel只会发送一次错误，然后关闭
4. 使用context可以优雅地取消正在进行的请求
5. 流式响应的内容是增量的，需要累积才能获得完整回复