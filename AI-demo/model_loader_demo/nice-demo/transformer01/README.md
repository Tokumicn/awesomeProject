# Qwen3-0.6B-MLX API Service

这是一个使用 FastAPI 构建的 API 服务，用于提供 Qwen3-0.6B-MLX-4bit 模型的文本生成功能。该服务使用 Apple 的 MLX 框架运行模型。

## 系统要求

- macOS 系统（MLX 框架要求）
- Python 3.8 或更高版本

## 安装依赖

```bash
pip install -r requirements.txt
```

## 运行服务

```bash
python app.py
```

服务将在 http://localhost:8000 上运行。

## API 使用说明

### 生成文本

**端点**: POST /generate

**请求体**:
```json
{
    "prompt": "你的输入文本",
    "max_tokens": 100,  // 可选，默认100
    "temperature": 0.7  // 可选，默认0.7
}
```

**响应**:
```json
{
    "generated_text": "生成的文本"
}
```

### 测试 API

你可以使用 curl 测试 API：

```bash
curl -X POST "http://localhost:8000/generate" \
     -H "Content-Type: application/json" \
     -d '{"prompt": "你好，请介绍一下自己", "max_tokens": 100, "temperature": 0.7}'
```

或者使用 Python requests：

```python
import requests

response = requests.post(
    "http://localhost:8000/generate",
    json={
        "prompt": "你好，请介绍一下自己",
        "max_tokens": 100,
        "temperature": 0.7
    }
)
print(response.json())
```

## 注意事项

1. 该服务使用 Apple 的 MLX 框架，因此只能在 macOS 系统上运行
2. 模型是 4bit 量化版本，可以在 Apple Silicon 芯片上高效运行
3. 第一次运行时，模型会被下载到本地，这可能需要一些时间 