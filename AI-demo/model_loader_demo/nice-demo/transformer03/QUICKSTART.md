# EasyOCR API 快速开始指南

## 🚀 5分钟快速开始

### 1. 环境准备

确保您的系统满足以下要求：
- macOS (推荐 M3 芯片)
- Python 3.8+
- 网络连接

### 2. 安装依赖

```bash
# 创建虚拟环境（推荐）
python3 -m venv venv
source venv/bin/activate

# 安装依赖
pip install -r requirements.txt
```

### 3. 启动服务

```bash
# 使用启动脚本（推荐）
python start_server.py

# 或直接启动
python api_server.py
```

### 4. 测试服务

```bash
# 运行测试脚本
python test_api.py

# 或访问API文档
# 浏览器打开: http://localhost:8000/docs
```

## 📖 基本使用

### 方法一：使用API接口

#### 1. 单张图片识别

```bash
curl -X POST "http://localhost:8000/ocr/upload" \
  -H "accept: application/json" \
  -H "Content-Type: multipart/form-data" \
  -F "file=@your_image.jpg" \
  -F "languages=ch_sim,en"
  
# 我的测试
curl --location 'http://localhost:8000/ocr/upload' \
--header 'accept: application/json' \
--form 'file=@"/AI-demo/OCR/images/20250527121324.png"' \
--form 'languages="ch_sim,en"'
```

#### 2. 批量图片识别

```bash
curl -X POST "http://localhost:8000/ocr/batch" \
  -H "accept: application/json" \
  -H "Content-Type: multipart/form-data" \
  -F "files=@image1.jpg" \
  -F "files=@image2.jpg"
```

### 方法二：使用Python代码

```python
import requests

# 识别单张图片
with open('image.jpg', 'rb') as f:
    files = {'file': f}
    data = {'languages': 'ch_sim,en'}
    response = requests.post('http://localhost:8000/ocr/upload', files=files, data=data)
    result = response.json()
    print(result)
```

### 方法三：直接使用模型

```python
from model_loader import get_model_loader

# 初始化模型
model_loader = get_model_loader(languages=['ch_sim', 'en'])

# 识别图片
results = model_loader.recognize_text('image.jpg')
for result in results:
    print(f"文本: {result['text']}")
    print(f"置信度: {result['confidence']}")
```

## 🔧 常用配置

### 环境变量配置

```bash
# 设置端口
export PORT=8080

# 启用调试模式
export DEBUG=true

# 禁用GPU（使用CPU）
export GPU_ENABLED=false

# 设置最大文件大小（MB）
export MAX_FILE_SIZE=20

# 设置日志级别
export LOG_LEVEL=DEBUG
```

### 支持的图片格式

- JPEG (.jpg, .jpeg)
- PNG (.png)
- BMP (.bmp)
- TIFF (.tiff)
- WebP (.webp)

### 支持的语言

常用语言：
- `ch_sim` - 简体中文
- `ch_tra` - 繁体中文
- `en` - 英文
- `ja` - 日文
- `ko` - 韩文

完整语言列表请参考 `config.py` 文件。

## 📊 API响应格式

### 成功响应

```json
{
  "success": true,
  "message": "OCR识别成功，共识别到 3 个文本区域",
  "data": [
    {
      "bbox": [[20, 20], [120, 20], [120, 50], [20, 50]],
      "text": "Hello World",
      "confidence": 0.95
    }
  ],
  "model_info": {
    "model_name": "Qualcomm/EasyOCR",
    "languages": ["ch_sim", "en"],
    "gpu_enabled": true,
    "device": "MPS"
  }
}
```

### 错误响应

```json
{
  "detail": "只支持图片文件"
}
```

## 🎯 性能优化

### Mac M3 芯片优化

1. **启用MPS加速**（默认启用）
   ```python
   model_loader = get_model_loader(gpu=True)
   ```

2. **调整画布大小**
   ```python
   # 对于大图片，可以减小画布大小
   results = model_loader.recognize_text('image.jpg', canvas_size=1280)
   ```

3. **调整识别参数**
   ```python
   # 提高识别精度（但会降低速度）
   results = model_loader.recognize_text('image.jpg', 
                                       text_threshold=0.8,
                                       link_threshold=0.5)
   ```

### 批量处理优化

```python
# 批量处理多张图片
image_paths = ['image1.jpg', 'image2.jpg', 'image3.jpg']
for path in image_paths:
    results = model_loader.recognize_text(path)
    # 处理结果...
```

## 🐛 常见问题

### Q1: 首次运行很慢？
A: 首次运行需要下载模型文件，请耐心等待。模型文件会保存在 `./models` 目录中。

### Q2: 内存不足？
A: 可以尝试以下方法：
- 减小 `canvas_size` 参数
- 使用CPU模式：`gpu=False`
- 关闭其他应用程序释放内存

### Q3: 识别准确率不高？
A: 可以尝试：
- 提高图片质量
- 调整 `text_threshold` 和 `link_threshold` 参数
- 确保图片中的文字清晰可见

### Q4: 支持的语言不显示？
A: 检查语言代码是否正确，常用语言代码：
- 简体中文：`ch_sim`
- 繁体中文：`ch_tra`
- 英文：`en`
- 日文：`ja`
- 韩文：`ko`

## 📚 更多资源

- [完整文档](README.md)
- [API文档](http://localhost:8000/docs)
- [配置说明](config.py)
- [测试脚本](test_api.py)

## 🆘 获取帮助

如果遇到问题，请：

1. 查看日志输出
2. 运行测试脚本：`python test_api.py`
3. 检查网络连接
4. 查看 [README.md](README.md) 中的故障排除部分

---

**提示**: 首次使用建议先运行 `python example_usage.py` 查看完整示例。 