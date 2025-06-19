"""
EasyOCR使用示例
展示如何使用模型加载器和API进行文本识别
"""

import requests
import json
import os
from PIL import Image
import numpy as np
from model_loader import get_model_loader

def example_direct_usage():
    """
    直接使用模型加载器的示例
    """
    print("=== 直接使用模型加载器示例 ===")
    
    try:
        # 初始化模型加载器
        model_loader = get_model_loader(languages=['ch_sim', 'en'], gpu=True)
        
        # 获取模型信息
        model_info = model_loader.get_model_info()
        print(f"模型信息: {json.dumps(model_info, indent=2, ensure_ascii=False)}")
        
        # 获取支持的语言
        languages = model_loader.get_supported_languages()
        print(f"支持的语言: {languages}")
        
        # 如果有测试图片，可以进行识别
        test_image_path = "test_image.jpg"
        if os.path.exists(test_image_path):
            print(f"\n识别图片: {test_image_path}")
            results = model_loader.recognize_text(test_image_path)
            
            print(f"识别结果:")
            for i, result in enumerate(results):
                print(f"  文本 {i+1}: {result['text']}")
                print(f"  置信度: {result['confidence']:.3f}")
                print(f"  位置: {result['bbox']}")
                print()
        else:
            print(f"\n测试图片 {test_image_path} 不存在，跳过识别测试")
            
    except Exception as e:
        print(f"直接使用示例失败: {str(e)}")

def example_api_usage():
    """
    使用API的示例
    """
    print("\n=== API使用示例 ===")
    
    # API基础URL
    base_url = "http://localhost:8000"
    
    try:
        # 1. 检查服务健康状态
        print("1. 检查服务健康状态")
        response = requests.get(f"{base_url}/health")
        print(f"健康检查结果: {response.json()}")
        
        # 2. 获取模型信息
        print("\n2. 获取模型信息")
        response = requests.get(f"{base_url}/model/info")
        print(f"模型信息: {json.dumps(response.json(), indent=2, ensure_ascii=False)}")
        
        # 3. 获取支持的语言
        print("\n3. 获取支持的语言")
        response = requests.get(f"{base_url}/model/languages")
        print(f"支持的语言: {response.json()}")
        
        # 4. 如果有测试图片，进行OCR识别
        test_image_path = "test_image.jpg"
        if os.path.exists(test_image_path):
            print(f"\n4. OCR识别图片: {test_image_path}")
            
            with open(test_image_path, 'rb') as f:
                files = {'file': f}
                data = {
                    'languages': 'ch_sim,en',
                    'detail': 1,
                    'paragraph': False
                }
                
                response = requests.post(f"{base_url}/ocr/upload", files=files, data=data)
                result = response.json()
                
                if result['success']:
                    print(f"识别成功，共识别到 {len(result['data'])} 个文本区域")
                    for i, item in enumerate(result['data']):
                        print(f"  文本 {i+1}: {item['text']}")
                        print(f"  置信度: {item['confidence']:.3f}")
                        print(f"  位置: {item['bbox']}")
                        print()
                else:
                    print(f"识别失败: {result['message']}")
        else:
            print(f"\n4. 测试图片 {test_image_path} 不存在，跳过API识别测试")
            
    except requests.exceptions.ConnectionError:
        print("无法连接到API服务器，请确保服务器正在运行")
    except Exception as e:
        print(f"API使用示例失败: {str(e)}")

def create_test_image():
    """
    创建一个简单的测试图片
    """
    print("\n=== 创建测试图片 ===")
    
    try:
        # 创建一个简单的测试图片
        from PIL import Image, ImageDraw, ImageFont
        
        # 创建白色背景图片
        img = Image.new('RGB', (400, 200), color='white')
        draw = ImageDraw.Draw(img)
        
        # 尝试使用系统字体，如果没有则使用默认字体
        try:
            # 在Mac上尝试使用系统字体
            font = ImageFont.truetype("/System/Library/Fonts/PingFang.ttc", 24)
        except:
            try:
                font = ImageFont.truetype("/System/Library/Fonts/Arial.ttf", 24)
            except:
                font = ImageFont.load_default()
        
        # 绘制文本
        draw.text((20, 20), "Hello World", fill='black', font=font)
        draw.text((20, 60), "你好世界", fill='black', font=font)
        draw.text((20, 100), "OCR Test", fill='black', font=font)
        draw.text((20, 140), "文本识别测试", fill='black', font=font)
        
        # 保存图片
        img.save("test_image.jpg")
        print("测试图片已创建: test_image.jpg")
        
    except Exception as e:
        print(f"创建测试图片失败: {str(e)}")

def main():
    """
    主函数
    """
    print("EasyOCR使用示例")
    print("=" * 50)
    
    # 创建测试图片
    create_test_image()
    
    # 直接使用模型加载器
    example_direct_usage()
    
    # 使用API
    example_api_usage()
    
    print("\n=== 使用说明 ===")
    print("1. 确保已安装所有依赖: pip install -r requirements.txt")
    print("2. 启动API服务器: python api_server.py")
    print("3. 访问API文档: http://localhost:8000/docs")
    print("4. 运行此示例: python example_usage.py")

if __name__ == "__main__":
    main() 