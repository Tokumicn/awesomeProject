# EasyOCR API 服务

基于 Qualcomm/EasyOCR 模型的文本识别 API 服务，专为 Mac M3 芯片优化。

## 项目概述

本项目提供了一个完整的 OCR（光学字符识别）解决方案，包括：

- **模型加载器**: 自动下载和加载 EasyOCR 模型
- **REST API**: 基于 FastAPI 的 Web 服务
- **多语言支持**: 支持中文、英文等多种语言
- **GPU 加速**: 针对 Mac M3 芯片的 MPS 加速优化
- **批量处理**: 支持单张和批量图片识别

## 功能特性

- ✅ 自动模型下载和加载
- ✅ 支持多种语言（中文、英文、日文、韩文等）
- ✅ Mac M3 芯片 GPU 加速支持
- ✅ RESTful API 接口
- ✅ 批量图片处理
- ✅ 详细的识别结果（文本、位置、置信度）
- ✅ 可配置的识别参数
- ✅ 完整的错误处理和日志记录

## 系统要求

- macOS (推荐 M3 芯片)
- Python 3.8+
- 至少 4GB 可用内存
- 网络连接（首次运行需要下载模型）

## 安装步骤

### 1. 克隆项目

```bash
git clone <repository-url>
cd transformer03
```

### 2. 创建虚拟环境（推荐）

```bash
python3 -m venv venv
source venv/bin/activate
```

### 3. 安装依赖

```bash
pip install -r requirements.txt
```

### 4. 验证安装

```bash
python -c "import torch; print(f'PyTorch版本: {torch.__version__}'); print(f'MPS可用: {torch.backends.mps.is_available()}')"
```

## 使用方法

### 方法一：直接使用模型加载器

```python
from model_loader import get_model_loader

# 初始化模型加载器
model_loader = get_model_loader(languages=['ch_sim', 'en'], gpu=True)

# 识别图片
results = model_loader.recognize_text('path/to/image.jpg')

# 处理结果
for result in results:
    print(f"文本: {result['text']}")
    print(f"置信度: {result['confidence']}")
    print(f"位置: {result['bbox']}")
```

### 方法二：启动 API 服务

```bash
python api_server.py
```

服务将在 `http://localhost:8000` 启动，您可以访问：
- API 文档：`http://localhost:8000/docs`
- 交互式文档：`http://localhost:8000/redoc`

### 方法三：运行示例

```bash
python example_usage.py
```

## API 接口说明

### 基础接口

| 接口 | 方法 | 描述 |
|------|------|------|
| `/` | GET | 获取 API 信息 |
| `/health` | GET | 健康检查 |
| `/model/info` | GET | 获取模型信息 |
| `/model/languages` | GET | 获取支持的语言列表 |

### OCR 识别接口

#### 1. 单张图片识别

**接口**: `POST /ocr/upload`

**参数**:
- `file`: 图片文件（必需）
- `languages`: 支持的语言，用逗号分隔（默认: `ch_sim,en`）
- `detail`: 返回详细程度，0=仅文本，1=包含位置信息（默认: 1）
- `paragraph`: 是否按段落分组（默认: false）
- `text_threshold`: 文本检测阈值（默认: 0.7）
- `link_threshold`: 文本连接阈值（默认: 0.4）
- `canvas_size`: 画布大小（默认: 2560）
- `mag_ratio`: 放大比例（默认: 1.5）

**响应示例**:
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

#### 2. 批量图片识别

**接口**: `POST /ocr/batch`

**参数**:
- `files`: 图片文件列表（最多10个）

**响应示例**:
```json
{
  "success": true,
  "message": "批量识别完成，共处理 2 个文件",
  "results": [
    {
      "file_index": 0,
      "filename": "image1.jpg",
      "success": true,
      "data": [...],
      "text_count": 3
    }
  ]
}
```

## 模型参数说明

### 输入参数

| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `languages` | List[str] | `['ch_sim', 'en']` | 支持的语言列表 |
| `gpu` | bool | `True` | 是否使用GPU加速 |
| `detail` | int | `1` | 返回详细程度 |
| `paragraph` | bool | `False` | 是否按段落分组 |
| `text_threshold` | float | `0.7` | 文本检测阈值 |
| `link_threshold` | float | `0.4` | 文本连接阈值 |
| `canvas_size` | int | `2560` | 画布大小 |
| `mag_ratio` | float | `1.5` | 放大比例 |

### 输出参数

| 字段 | 类型 | 说明 |
|------|------|------|
| `bbox` | List[List[int]] | 文本框坐标 [[x1,y1], [x2,y2], [x3,y3], [x4,y4]] |
| `text` | str | 识别的文本内容 |
| `confidence` | float | 置信度 (0-1) |

## 支持的语言

| 语言代码 | 语言名称 |
|----------|----------|
| `ch_sim` | 简体中文 |
| `ch_tra` | 繁体中文 |
| `en` | 英文 |
| `ja` | 日文 |
| `ko` | 韩文 |
| `th` | 泰文 |
| `vi` | 越南文 |
| `ar` | 阿拉伯文 |
| `hi` | 印地文 |
| `bn` | 孟加拉文 |

## 性能优化

### Mac M3 芯片优化

1. **MPS 加速**: 自动检测并使用 Apple Metal Performance Shaders
2. **内存管理**: 优化的模型加载和内存使用
3. **并发处理**: 支持多请求并发处理

### 性能建议

1. **首次运行**: 模型下载可能需要几分钟，请耐心等待
2. **内存使用**: 建议至少 4GB 可用内存
3. **图片大小**: 建议图片分辨率不超过 4K
4. **批量处理**: 单次最多处理 10 张图片

## 故障排除

### 常见问题

1. **模型下载失败**
   ```bash
   # 检查网络连接
   curl -I https://huggingface.co
   
   # 手动下载模型
   python -c "import easyocr; easyocr.Reader(['ch_sim', 'en'])"
   ```

2. **MPS 不可用**
   ```bash
   # 检查 PyTorch 版本
   python -c "import torch; print(torch.__version__)"
   
   # 检查 MPS 支持
   python -c "import torch; print(torch.backends.mps.is_available())"
   ```

3. **内存不足**
   ```bash
   # 减少画布大小
   canvas_size=1280
   
   # 使用 CPU 模式
   gpu=False
   ```

### 日志查看

```bash
# 查看详细日志
python api_server.py --log-level debug
```

## 开发说明

### 项目结构

```
transformer03/
├── model_loader.py      # 模型加载器
├── api_server.py        # API 服务器
├── example_usage.py     # 使用示例
├── requirements.txt     # 依赖列表
├── README.md           # 项目文档
└── models/             # 模型存储目录（自动创建）
```

### 扩展开发

1. **添加新语言**: 修改 `languages` 参数
2. **自定义参数**: 在 `OCRRequest` 模型中添加新字段
3. **预处理**: 在 `recognize_text` 方法中添加图像预处理
4. **后处理**: 在返回结果前添加文本后处理

## 许可证

本项目基于 MIT 许可证开源。

## 贡献

欢迎提交 Issue 和 Pull Request！

## 联系方式

如有问题，请通过以下方式联系：
- 提交 GitHub Issue
- 发送邮件至项目维护者

---

**注意**: 首次运行时会自动下载模型文件，请确保网络连接正常。 