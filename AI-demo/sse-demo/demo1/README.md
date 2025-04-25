# SSE Demo in Go

这是一个使用 Go 语言实现的 Server-Sent Events (SSE) 示例应用。该应用程序通过 SSE 协议向客户端实时推送文本内容。

## 功能特点

- 提供 SSE 协议的 HTTP 接口
- 支持实时文本流式输出
- 包含简单的 Web 演示页面
- 处理客户端断开连接

## 运行方法

1. 确保已安装 Go 环境
2. 启动服务器:

```bash
go run main.go
```

3. 访问 http://localhost:8080 进行演示

## API 使用方法

- SSE 端点: `GET /stream`
- 返回的事件类型: `message`
- 每个事件包含 ID 和文本数据

## 客户端示例代码

```javascript
const eventSource = new EventSource('/stream');

eventSource.addEventListener('message', function(e) {
    console.log(e.data);
});

eventSource.onerror = function(e) {
    console.log('SSE error', e);
    eventSource.close();
};
```

## 自定义使用

可以修改 `main.go` 中的 `SSEHandler` 函数来改变输出的文本内容或增加更多功能。 