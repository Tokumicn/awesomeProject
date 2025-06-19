"""
EasyOCR API 配置文件
管理API服务的各种设置和参数
"""

import os
from typing import List, Dict, Any

class Config:
    """配置类"""
    
    # 服务器配置
    HOST = os.getenv("HOST", "0.0.0.0")
    PORT = int(os.getenv("PORT", 8000))
    DEBUG = os.getenv("DEBUG", "False").lower() == "true"
    
    # 模型配置
    DEFAULT_LANGUAGES = ["ch_sim", "en"]
    MODEL_STORAGE_DIR = os.getenv("MODEL_STORAGE_DIR", "./models")
    GPU_ENABLED = os.getenv("GPU_ENABLED", "True").lower() == "true"
    
    # OCR参数默认值
    OCR_PARAMS = {
        "detail": 1,
        "paragraph": False,
        "height_ths": 0.5,
        "width_ths": 0.5,
        "contrast_ths": 0.1,
        "text_threshold": 0.7,
        "link_threshold": 0.4,
        "low_text": 0.4,
        "canvas_size": 2560,
        "mag_ratio": 1.5
    }
    
    # 文件上传配置
    MAX_FILE_SIZE = int(os.getenv("MAX_FILE_SIZE", 10 * 1024 * 1024))  # 10MB
    ALLOWED_EXTENSIONS = {".jpg", ".jpeg", ".png", ".bmp", ".tiff", ".webp"}
    MAX_FILES_PER_REQUEST = int(os.getenv("MAX_FILES_PER_REQUEST", 10))
    
    # 日志配置
    LOG_LEVEL = os.getenv("LOG_LEVEL", "INFO")
    LOG_FORMAT = "%(asctime)s - %(name)s - %(levelname)s - %(message)s"
    
    # CORS配置
    CORS_ORIGINS = [
        "http://localhost:3000",
        "http://localhost:8080",
        "http://127.0.0.1:3000",
        "http://127.0.0.1:8080",
        "*"  # 开发环境允许所有来源
    ]
    
    # 缓存配置
    CACHE_ENABLED = os.getenv("CACHE_ENABLED", "True").lower() == "true"
    CACHE_TTL = int(os.getenv("CACHE_TTL", 3600))  # 1小时
    
    # 性能配置
    WORKER_PROCESSES = int(os.getenv("WORKER_PROCESSES", 1))
    MAX_CONCURRENT_REQUESTS = int(os.getenv("MAX_CONCURRENT_REQUESTS", 10))
    
    # 安全配置
    API_KEY_ENABLED = os.getenv("API_KEY_ENABLED", "False").lower() == "true"
    API_KEY_HEADER = "X-API-Key"
    
    @classmethod
    def get_ocr_params(cls, **kwargs) -> Dict[str, Any]:
        """
        获取OCR参数，支持覆盖默认值
        
        Args:
            **kwargs: 要覆盖的参数
            
        Returns:
            Dict[str, Any]: OCR参数字典
        """
        params = cls.OCR_PARAMS.copy()
        params.update(kwargs)
        return params
    
    @classmethod
    def get_supported_languages(cls) -> List[str]:
        """
        获取支持的语言列表
        
        Returns:
            List[str]: 支持的语言代码列表
        """
        return [
            "ch_sim",  # 简体中文
            "ch_tra",  # 繁体中文
            "en",      # 英文
            "ja",      # 日文
            "ko",      # 韩文
            "th",      # 泰文
            "vi",      # 越南文
            "ar",      # 阿拉伯文
            "hi",      # 印地文
            "bn",      # 孟加拉文
            "ru",      # 俄文
            "fr",      # 法文
            "de",      # 德文
            "es",      # 西班牙文
            "it",      # 意大利文
            "pt",      # 葡萄牙文
            "nl",      # 荷兰文
            "pl",      # 波兰文
            "tr",      # 土耳其文
            "sv",      # 瑞典文
            "da",      # 丹麦文
            "no",      # 挪威文
            "fi",      # 芬兰文
            "cs",      # 捷克文
            "hu",      # 匈牙利文
            "ro",      # 罗马尼亚文
            "bg",      # 保加利亚文
            "hr",      # 克罗地亚文
            "sk",      # 斯洛伐克文
            "sl",      # 斯洛文尼亚文
            "et",      # 爱沙尼亚文
            "lv",      # 拉脱维亚文
            "lt",      # 立陶宛文
            "mt",      # 马耳他文
            "el",      # 希腊文
            "he",      # 希伯来文
            "fa",      # 波斯文
            "ur",      # 乌尔都文
            "ta",      # 泰米尔文
            "te",      # 泰卢固文
            "kn",      # 卡纳达文
            "ml",      # 马拉雅拉姆文
            "gu",      # 古吉拉特文
            "pa",      # 旁遮普文
            "or",      # 奥里亚文
            "as",      # 阿萨姆文
            "ne",      # 尼泊尔文
            "si",      # 僧伽罗文
            "my",      # 缅甸文
            "km",      # 高棉文
            "lo",      # 老挝文
            "ka",      # 格鲁吉亚文
            "am",      # 阿姆哈拉文
            "ti",      # 提格里尼亚文
            "so",      # 索马里文
            "sw",      # 斯瓦希里文
            "zu",      # 祖鲁文
            "af",      # 南非荷兰文
            "xh",      # 科萨文
            "st",      # 塞索托文
            "tn",      # 茨瓦纳文
            "ts",      # 聪加文
            "ve",      # 文达文
            "ss",      # 斯威士文
            "nr",      # 南恩德贝莱文
            "nd",      # 北恩德贝莱文
            "sn",      # 绍纳文
            "ny",      # 奇切瓦文
            "lg",      # 干达文
            "rw",      # 卢旺达文
            "ak",      # 阿坎文
            "yo",      # 约鲁巴文
            "ig",      # 伊博文
            "ha",      # 豪萨文
            "ff",      # 富拉文
            "wo",      # 沃洛夫文
            "bm",      # 班巴拉文
            "dy",      # 迪乌拉文
            "sg",      # 桑戈文
            "ln",      # 林加拉文
            "sw",      # 斯瓦希里文
            "rw",      # 卢旺达文
            "lg",      # 干达文
            "ny",      # 奇切瓦文
            "sn",      # 绍纳文
            "nd",      # 北恩德贝莱文
            "nr",      # 南恩德贝莱文
            "ss",      # 斯威士文
            "ve",      # 文达文
            "ts",      # 聪加文
            "tn",      # 茨瓦纳文
            "st",      # 塞索托文
            "xh",      # 科萨文
            "af",      # 南非荷兰文
            "zu",      # 祖鲁文
            "sw",      # 斯瓦希里文
            "so",      # 索马里文
            "ti",      # 提格里尼亚文
            "am",      # 阿姆哈拉文
            "ka",      # 格鲁吉亚文
            "lo",      # 老挝文
            "km",      # 高棉文
            "my",      # 缅甸文
            "si",      # 僧伽罗文
            "ne",      # 尼泊尔文
            "as",      # 阿萨姆文
            "or",      # 奥里亚文
            "pa",      # 旁遮普文
            "gu",      # 古吉拉特文
            "ml",      # 马拉雅拉姆文
            "kn",      # 卡纳达文
            "te",      # 泰卢固文
            "ta",      # 泰米尔文
            "ur",      # 乌尔都文
            "fa",      # 波斯文
            "he",      # 希伯来文
            "el",      # 希腊文
            "mt",      # 马耳他文
            "lt",      # 立陶宛文
            "lv",      # 拉脱维亚文
            "et",      # 爱沙尼亚文
            "sl",      # 斯洛文尼亚文
            "sk",      # 斯洛伐克文
            "hr",      # 克罗地亚文
            "bg",      # 保加利亚文
            "ro",      # 罗马尼亚文
            "hu",      # 匈牙利文
            "cs",      # 捷克文
            "fi",      # 芬兰文
            "no",      # 挪威文
            "da",      # 丹麦文
            "sv",      # 瑞典文
            "tr",      # 土耳其文
            "pl",      # 波兰文
            "nl",      # 荷兰文
            "pt",      # 葡萄牙文
            "it",      # 意大利文
            "es",      # 西班牙文
            "de",      # 德文
            "fr",      # 法文
            "ru",      # 俄文
            "bn",      # 孟加拉文
            "hi",      # 印地文
            "ar",      # 阿拉伯文
            "vi",      # 越南文
            "th",      # 泰文
            "ko",      # 韩文
            "ja",      # 日文
            "en",      # 英文
            "ch_tra",  # 繁体中文
            "ch_sim"   # 简体中文
        ]
    
    @classmethod
    def validate_languages(cls, languages: List[str]) -> List[str]:
        """
        验证语言列表，过滤不支持的语言
        
        Args:
            languages: 要验证的语言列表
            
        Returns:
            List[str]: 有效的语言列表
        """
        supported = set(cls.get_supported_languages())
        valid_languages = [lang for lang in languages if lang in supported]
        
        if not valid_languages:
            # 如果没有有效语言，使用默认语言
            valid_languages = cls.DEFAULT_LANGUAGES
        
        return valid_languages
    
    @classmethod
    def get_model_info(cls) -> Dict[str, Any]:
        """
        获取模型信息
        
        Returns:
            Dict[str, Any]: 模型信息字典
        """
        return {
            "model_name": "Qualcomm/EasyOCR",
            "version": "1.0.0",
            "description": "基于EasyOCR的文本识别模型",
            "supported_languages": cls.get_supported_languages(),
            "default_languages": cls.DEFAULT_LANGUAGES,
            "gpu_enabled": cls.GPU_ENABLED,
            "model_storage_dir": cls.MODEL_STORAGE_DIR,
            "max_file_size": cls.MAX_FILE_SIZE,
            "max_files_per_request": cls.MAX_FILES_PER_REQUEST
        }

# 开发环境配置
class DevelopmentConfig(Config):
    """开发环境配置"""
    DEBUG = True
    LOG_LEVEL = "DEBUG"
    CORS_ORIGINS = ["*"]

# 生产环境配置
class ProductionConfig(Config):
    """生产环境配置"""
    DEBUG = False
    LOG_LEVEL = "WARNING"
    CORS_ORIGINS = [
        "https://yourdomain.com",
        "https://api.yourdomain.com"
    ]
    API_KEY_ENABLED = True

# 测试环境配置
class TestingConfig(Config):
    """测试环境配置"""
    DEBUG = True
    LOG_LEVEL = "DEBUG"
    PORT = 8001
    CACHE_ENABLED = False

# 根据环境变量选择配置
def get_config():
    """根据环境变量获取配置"""
    env = os.getenv("ENV", "development").lower()
    
    if env == "production":
        return ProductionConfig
    elif env == "testing":
        return TestingConfig
    else:
        return DevelopmentConfig 