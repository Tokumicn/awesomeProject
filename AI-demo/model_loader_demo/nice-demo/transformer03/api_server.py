"""
EasyOCR API服务器
提供基于FastAPI的OCR文本识别REST API服务
"""

from fastapi import FastAPI, File, UploadFile, HTTPException, Form
from fastapi.responses import JSONResponse
from fastapi.middleware.cors import CORSMiddleware
import uvicorn
import logging
import os
import tempfile
from typing import List, Dict, Any, Optional
from pydantic import BaseModel

from model_loader import get_model_loader

# 配置日志
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# 创建FastAPI应用
app = FastAPI(
    title="EasyOCR API",
    description="基于Qualcomm/EasyOCR模型的文本识别API服务",
    version="1.0.0",
    docs_url="/docs",
    redoc_url="/redoc"
)

# 添加CORS中间件
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# 请求模型
class OCRRequest(BaseModel):
    """OCR请求模型"""
    languages: Optional[List[str]] = ['ch_sim', 'en']
    detail: Optional[int] = 1
    paragraph: Optional[bool] = False
    height_ths: Optional[float] = 0.5
    width_ths: Optional[float] = 0.5
    contrast_ths: Optional[float] = 0.1
    text_threshold: Optional[float] = 0.7
    link_threshold: Optional[float] = 0.4
    low_text: Optional[float] = 0.4
    canvas_size: Optional[int] = 2560
    mag_ratio: Optional[float] = 1.5

# 响应模型
class OCRResponse(BaseModel):
    """OCR响应模型"""
    success: bool
    message: str
    data: Optional[List[Dict[str, Any]]] = None
    model_info: Optional[Dict[str, Any]] = None

class ModelInfoResponse(BaseModel):
    """模型信息响应模型"""
    success: bool
    message: str
    model_info: Dict[str, Any]

@app.on_event("startup")
async def startup_event():
    """应用启动时初始化模型"""
    try:
        logger.info("正在初始化EasyOCR模型...")
        # 初始化模型加载器
        get_model_loader()
        logger.info("EasyOCR模型初始化完成")
    except Exception as e:
        logger.error(f"模型初始化失败: {str(e)}")
        raise

@app.get("/", response_model=Dict[str, str])
async def root():
    """
    根路径，返回API信息
    
    Returns:
        Dict[str, str]: API信息
    """
    return {
        "message": "EasyOCR API服务",
        "version": "1.0.0",
        "docs": "/docs",
        "model": "Qualcomm/EasyOCR"
    }

@app.get("/health", response_model=Dict[str, str])
async def health_check():
    """
    健康检查接口
    
    Returns:
        Dict[str, str]: 服务状态
    """
    return {"status": "healthy", "message": "EasyOCR API服务运行正常"}

@app.get("/model/info", response_model=ModelInfoResponse)
async def get_model_info():
    """
    获取模型信息
    
    Returns:
        ModelInfoResponse: 模型信息
    """
    try:
        model_loader = get_model_loader()
        model_info = model_loader.get_model_info()
        
        return ModelInfoResponse(
            success=True,
            message="获取模型信息成功",
            model_info=model_info
        )
    except Exception as e:
        logger.error(f"获取模型信息失败: {str(e)}")
        raise HTTPException(status_code=500, detail=f"获取模型信息失败: {str(e)}")

@app.get("/model/languages", response_model=Dict[str, Any])
async def get_supported_languages():
    """
    获取支持的语言列表
    
    Returns:
        Dict[str, Any]: 支持的语言列表
    """
    try:
        model_loader = get_model_loader()
        languages = model_loader.get_supported_languages()
        
        return {
            "success": True,
            "message": "获取支持语言列表成功",
            "languages": languages
        }
    except Exception as e:
        logger.error(f"获取支持语言列表失败: {str(e)}")
        raise HTTPException(status_code=500, detail=f"获取支持语言列表失败: {str(e)}")

@app.post("/ocr/upload", response_model=OCRResponse)
async def ocr_upload(
    file: UploadFile = File(..., description="要识别的图片文件"),
    languages: Optional[str] = Form('ch_sim,en', description="支持的语言，用逗号分隔"),
    detail: Optional[int] = Form(1, description="返回详细程度 (0=仅文本, 1=包含位置信息)"),
    paragraph: Optional[bool] = Form(False, description="是否按段落分组"),
    height_ths: Optional[float] = Form(0.5, description="高度阈值"),
    width_ths: Optional[float] = Form(0.5, description="宽度阈值"),
    contrast_ths: Optional[float] = Form(0.1, description="对比度阈值"),
    text_threshold: Optional[float] = Form(0.7, description="文本检测阈值"),
    link_threshold: Optional[float] = Form(0.4, description="文本连接阈值"),
    low_text: Optional[float] = Form(0.4, description="低文本阈值"),
    canvas_size: Optional[int] = Form(2560, description="画布大小"),
    mag_ratio: Optional[float] = Form(1.5, description="放大比例")
):
    """
    上传图片进行OCR文本识别
    
    Args:
        file (UploadFile): 要识别的图片文件
        languages (str): 支持的语言，用逗号分隔
        detail (int): 返回详细程度
        paragraph (bool): 是否按段落分组
        height_ths (float): 高度阈值
        width_ths (float): 宽度阈值
        contrast_ths (float): 对比度阈值
        text_threshold (float): 文本检测阈值
        link_threshold (float): 文本连接阈值
        low_text (float): 低文本阈值
        canvas_size (int): 画布大小
        mag_ratio (float): 放大比例
    
    Returns:
        OCRResponse: OCR识别结果
    """
    try:
        # 验证文件类型
        if not file.content_type.startswith('image/'):
            raise HTTPException(status_code=400, detail="只支持图片文件")
        
        # 解析语言参数
        language_list = [lang.strip() for lang in languages.split(',') if lang.strip()]
        
        # 读取文件内容
        file_content = await file.read()
        
        # 获取模型加载器
        model_loader = get_model_loader()
        
        # 执行OCR识别
        results = model_loader.recognize_text_from_bytes(
            file_content,
            detail=detail,
            paragraph=paragraph,
            height_ths=height_ths,
            width_ths=width_ths,
            contrast_ths=contrast_ths,
            text_threshold=text_threshold,
            link_threshold=link_threshold,
            low_text=low_text,
            canvas_size=canvas_size,
            mag_ratio=mag_ratio
        )
        
        # 获取模型信息
        model_info = model_loader.get_model_info()
        
        return OCRResponse(
            success=True,
            message=f"OCR识别成功，共识别到 {len(results)} 个文本区域",
            data=results,
            model_info=model_info
        )
        
    except HTTPException:
        raise
    except Exception as e:
        logger.error(f"OCR识别失败: {str(e)}")
        raise HTTPException(status_code=500, detail=f"OCR识别失败: {str(e)}")

@app.post("/ocr/path", response_model=OCRResponse)
async def ocr_path(request: OCRRequest):
    """
    通过图片路径进行OCR文本识别
    
    Args:
        request (OCRRequest): OCR请求参数
    
    Returns:
        OCRResponse: OCR识别结果
    """
    try:
        # 这里需要图片路径，实际使用时可能需要调整
        # 暂时返回错误信息
        raise HTTPException(status_code=400, detail="请使用 /ocr/upload 接口上传图片文件")
        
    except HTTPException:
        raise
    except Exception as e:
        logger.error(f"OCR识别失败: {str(e)}")
        raise HTTPException(status_code=500, detail=f"OCR识别失败: {str(e)}")

@app.post("/ocr/batch", response_model=Dict[str, Any])
async def ocr_batch(files: List[UploadFile] = File(..., description="要识别的图片文件列表")):
    """
    批量OCR识别
    
    Args:
        files (List[UploadFile]): 要识别的图片文件列表
    
    Returns:
        Dict[str, Any]: 批量识别结果
    """
    try:
        if len(files) > 10:
            raise HTTPException(status_code=400, detail="一次最多处理10个文件")
        
        results = []
        model_loader = get_model_loader()
        
        for i, file in enumerate(files):
            try:
                # 验证文件类型
                if not file.content_type.startswith('image/'):
                    results.append({
                        "file_index": i,
                        "filename": file.filename,
                        "success": False,
                        "error": "只支持图片文件"
                    })
                    continue
                
                # 读取文件内容
                file_content = await file.read()
                
                # 执行OCR识别
                ocr_results = model_loader.recognize_text_from_bytes(file_content)
                
                results.append({
                    "file_index": i,
                    "filename": file.filename,
                    "success": True,
                    "data": ocr_results,
                    "text_count": len(ocr_results)
                })
                
            except Exception as e:
                results.append({
                    "file_index": i,
                    "filename": file.filename,
                    "success": False,
                    "error": str(e)
                })
        
        return {
            "success": True,
            "message": f"批量识别完成，共处理 {len(files)} 个文件",
            "results": results,
            "model_info": model_loader.get_model_info()
        }
        
    except HTTPException:
        raise
    except Exception as e:
        logger.error(f"批量OCR识别失败: {str(e)}")
        raise HTTPException(status_code=500, detail=f"批量OCR识别失败: {str(e)}")

if __name__ == "__main__":
    # 启动服务器
    uvicorn.run(
        "api_server:app",
        host="0.0.0.0",
        port=8000,
        reload=True,
        log_level="info"
    ) 