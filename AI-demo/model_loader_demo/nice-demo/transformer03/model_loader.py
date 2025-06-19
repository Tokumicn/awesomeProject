"""
EasyOCR模型加载器
用于加载Qualcomm/EasyOCR模型并提供OCR功能
"""

import torch
import easyocr
from typing import List, Dict, Any, Optional
import logging

# 配置日志
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

class EasyOCRModelLoader:
    """
    EasyOCR模型加载器类
    
    该类负责加载和初始化EasyOCR模型，提供OCR文本识别功能
    支持多种语言，特别针对Mac M3芯片进行了优化
    """
    
    def __init__(self, languages: List[str] = ['ch_sim', 'en'], gpu: bool = True):
        """
        初始化EasyOCR模型加载器
        
        Args:
            languages (List[str]): 支持的语言列表
                - 'ch_sim': 简体中文
                - 'en': 英文
                - 'ja': 日文
                - 'ko': 韩文
                - 等等
            gpu (bool): 是否使用GPU加速，Mac M3芯片建议设置为True
        """
        self.languages = languages
        self.gpu = gpu and torch.backends.mps.is_available()
        self.reader = None
        self._load_model()
    
    def _load_model(self):
        """加载EasyOCR模型"""
        try:
            logger.info(f"正在加载EasyOCR模型，语言: {self.languages}")
            logger.info(f"GPU加速: {self.gpu}")
            
            # 初始化EasyOCR读取器
            self.reader = easyocr.Reader(
                self.languages,
                gpu=self.gpu,
                model_storage_directory='./models',
                download_enabled=True
            )
            
            logger.info("EasyOCR模型加载成功")
            
        except Exception as e:
            logger.error(f"模型加载失败: {str(e)}")
            raise
    
    def recognize_text(self, image_path: str, **kwargs) -> List[Dict[str, Any]]:
        """
        识别图片中的文本
        
        Args:
            image_path (str): 图片文件路径
            **kwargs: 额外的识别参数
                - detail (int): 返回详细程度 (0=仅文本, 1=包含位置信息)
                - paragraph (bool): 是否按段落分组
                - height_ths (float): 高度阈值
                - width_ths (float): 宽度阈值
                - contrast_ths (float): 对比度阈值
                - text_threshold (float): 文本检测阈值
                - link_threshold (float): 文本连接阈值
                - low_text (float): 低文本阈值
                - canvas_size (int): 画布大小
                - mag_ratio (float): 放大比例
        
        Returns:
            List[Dict[str, Any]]: 识别结果列表，每个元素包含：
                - bbox: 文本框坐标 [[x1,y1], [x2,y2], [x3,y3], [x4,y4]]
                - text: 识别的文本内容
                - confidence: 置信度 (0-1)
        """
        try:
            if self.reader is None:
                raise ValueError("模型未正确加载")
            
            logger.info(f"开始识别图片: {image_path}")
            
            # 执行文本识别
            results = self.reader.readtext(image_path, **kwargs)
            
            # 格式化结果
            formatted_results = []
            for bbox, text, confidence in results:
                # 转换bbox为Python原生类型
                formatted_bbox = [[int(x), int(y)] for x, y in bbox]
                formatted_results.append({
                    'bbox': formatted_bbox,
                    'text': str(text),
                    'confidence': float(confidence)
                })
            
            logger.info(f"识别完成，共识别到 {len(formatted_results)} 个文本区域")
            return formatted_results
            
        except Exception as e:
            logger.error(f"文本识别失败: {str(e)}")
            raise
    
    def recognize_text_from_bytes(self, image_bytes: bytes, **kwargs) -> List[Dict[str, Any]]:
        """
        从字节数据识别文本
        
        Args:
            image_bytes (bytes): 图片字节数据
            **kwargs: 额外的识别参数（同recognize_text方法）
        
        Returns:
            List[Dict[str, Any]]: 识别结果列表
        """
        try:
            if self.reader is None:
                raise ValueError("模型未正确加载")
            
            logger.info("开始识别字节数据中的文本")
            
            # 执行文本识别
            results = self.reader.readtext(image_bytes, **kwargs)
            
            # 格式化结果
            formatted_results = []
            for bbox, text, confidence in results:
                # 转换bbox为Python原生类型
                formatted_bbox = [[int(x), int(y)] for x, y in bbox]
                formatted_results.append({
                    'bbox': formatted_bbox,
                    'text': str(text),
                    'confidence': float(confidence)
                })
            
            logger.info(f"识别完成，共识别到 {len(formatted_results)} 个文本区域")
            return formatted_results
            
        except Exception as e:
            logger.error(f"文本识别失败: {str(e)}")
            raise
    
    def get_supported_languages(self) -> List[str]:
        """
        获取支持的语言列表
        
        Returns:
            List[str]: 支持的语言代码列表
        """
        return self.languages
    
    def get_model_info(self) -> Dict[str, Any]:
        """
        获取模型信息
        
        Returns:
            Dict[str, Any]: 模型信息字典
        """
        return {
            'model_name': 'Qualcomm/EasyOCR',
            'languages': self.languages,
            'gpu_enabled': self.gpu,
            'device': 'MPS' if self.gpu and torch.backends.mps.is_available() else 'CPU'
        }

# 全局模型实例
_model_loader = None

def get_model_loader(languages: List[str] = ['ch_sim', 'en'], gpu: bool = True) -> EasyOCRModelLoader:
    """
    获取全局模型加载器实例（单例模式）
    
    Args:
        languages (List[str]): 支持的语言列表
        gpu (bool): 是否使用GPU加速
    
    Returns:
        EasyOCRModelLoader: 模型加载器实例
    """
    global _model_loader
    if _model_loader is None:
        _model_loader = EasyOCRModelLoader(languages, gpu)
    return _model_loader 