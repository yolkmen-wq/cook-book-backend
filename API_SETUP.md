# API 配置说明

## DeepSeek API 配置

当前项目使用 DeepSeek API 提供 AI 聊天功能。如果遇到以下错误：

```
API request failed with status 402: {"error":{"message":"Insufficient Balance","type":"unknown_error","param":null,"code":"invalid_request_error"}}
```

或者：

```
DeepSeek API key is not configured. Please set AI_API_KEY environment variable
```

说明需要配置有效的 DeepSeek API 密钥。

## 解决方案

### 方法1：设置环境变量

1. 获取 DeepSeek API 密钥：
   - 访问 [DeepSeek 官网](https://platform.deepseek.com/)
   - 注册账号并获取 API 密钥
   - 确保账户有足够余额

2. 设置环境变量：
   ```bash
   # Windows (PowerShell)
   $env:AI_API_KEY="your_deepseek_api_key_here"
   
   # Linux/Mac
   export AI_API_KEY="your_deepseek_api_key_here"
   ```

### 方法2：创建 .env 文件

1. 复制 `.env.example` 文件为 `.env`：
   ```bash
   cp .env.example .env
   ```

2. 编辑 `.env` 文件，设置你的 API 密钥：
   ```
   AI_API_KEY=your_deepseek_api_key_here
   AI_PROVIDER=deepseek
   AI_MODEL=deepseek-chat
   ```

### 方法3：Docker 环境配置

在 `docker-compose.yml` 中添加环境变量：

```yaml
services:
  backend:
    environment:
      - AI_API_KEY=your_deepseek_api_key_here
      - AI_PROVIDER=deepseek
      - AI_MODEL=deepseek-chat
```

## 验证配置

重启应用后，尝试发送聊天请求。如果配置正确，应该能正常获得 AI 回复。

## 注意事项

- 请勿将 API 密钥提交到版本控制系统
- 定期检查 API 账户余额
- 如需切换到其他 AI 提供商，请相应修改 `AI_PROVIDER` 和相关配置