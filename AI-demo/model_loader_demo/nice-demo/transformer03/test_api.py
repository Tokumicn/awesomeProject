"""
API测试脚本
用于测试EasyOCR API的各项功能
"""

import requests
import json
import time
import os
from PIL import Image, ImageDraw, ImageFont

class APITester:
    """API测试类"""
    
    def __init__(self, base_url="http://localhost:8000"):
        self.base_url = base_url
        self.session = requests.Session()
    
    def test_health_check(self):
        """测试健康检查接口"""
        print("=== 测试健康检查接口 ===")
        try:
            response = self.session.get(f"{self.base_url}/health")
            print(f"状态码: {response.status_code}")
            print(f"响应: {response.json()}")
            return response.status_code == 200
        except Exception as e:
            print(f"健康检查失败: {str(e)}")
            return False
    
    def test_root_endpoint(self):
        """测试根路径接口"""
        print("\n=== 测试根路径接口 ===")
        try:
            response = self.session.get(f"{self.base_url}/")
            print(f"状态码: {response.status_code}")
            print(f"响应: {response.json()}")
            return response.status_code == 200
        except Exception as e:
            print(f"根路径测试失败: {str(e)}")
            return False
    
    def test_model_info(self):
        """测试模型信息接口"""
        print("\n=== 测试模型信息接口 ===")
        try:
            response = self.session.get(f"{self.base_url}/model/info")
            print(f"状态码: {response.status_code}")
            if response.status_code == 200:
                data = response.json()
                print(f"模型名称: {data['model_info']['model_name']}")
                print(f"支持语言: {data['model_info']['languages']}")
                print(f"GPU启用: {data['model_info']['gpu_enabled']}")
                print(f"设备: {data['model_info']['device']}")
            else:
                print(f"响应: {response.json()}")
            return response.status_code == 200
        except Exception as e:
            print(f"模型信息测试失败: {str(e)}")
            return False
    
    def test_languages(self):
        """测试支持语言接口"""
        print("\n=== 测试支持语言接口 ===")
        try:
            response = self.session.get(f"{self.base_url}/model/languages")
            print(f"状态码: {response.status_code}")
            if response.status_code == 200:
                data = response.json()
                print(f"支持的语言: {data['languages']}")
            else:
                print(f"响应: {response.json()}")
            return response.status_code == 200
        except Exception as e:
            print(f"语言列表测试失败: {str(e)}")
            return False
    
    def create_test_image(self, filename="test_api_image.jpg"):
        """创建测试图片"""
        print(f"\n=== 创建测试图片: {filename} ===")
        try:
            # 创建白色背景图片
            img = Image.new('RGB', (500, 300), color='white')
            draw = ImageDraw.Draw(img)
            
            # 尝试使用系统字体
            try:
                font = ImageFont.truetype("/System/Library/Fonts/PingFang.ttc", 20)
            except:
                try:
                    font = ImageFont.truetype("/System/Library/Fonts/Arial.ttf", 20)
                except:
                    font = ImageFont.load_default()
            
            # 绘制多种语言的文本
            texts = [
                (20, 20, "Hello World - English"),
                (20, 60, "你好世界 - 中文"),
                (20, 100, "こんにちは世界 - 日本語"),
                (20, 140, "안녕하세요 세계 - 한국어"),
                (20, 180, "OCR API Test"),
                (20, 220, "文本识别测试"),
                (20, 260, "API テスト")
            ]
            
            for x, y, text in texts:
                draw.text((x, y), text, fill='black', font=font)
            
            # 保存图片
            img.save(filename)
            print(f"测试图片已创建: {filename}")
            return filename
            
        except Exception as e:
            print(f"创建测试图片失败: {str(e)}")
            return None
    
    def test_ocr_upload(self, image_path):
        """测试OCR上传接口"""
        print(f"\n=== 测试OCR上传接口 ===")
        try:
            if not os.path.exists(image_path):
                print(f"测试图片不存在: {image_path}")
                return False
            
            with open(image_path, 'rb') as f:
                files = {'file': f}
                data = {
                    'languages': 'ch_sim,en,ja,ko',
                    'detail': 1,
                    'paragraph': False,
                    'text_threshold': 0.7
                }
                
                start_time = time.time()
                response = self.session.post(f"{self.base_url}/ocr/upload", files=files, data=data)
                end_time = time.time()
                
                print(f"状态码: {response.status_code}")
                print(f"响应时间: {end_time - start_time:.2f}秒")
                
                if response.status_code == 200:
                    result = response.json()
                    print(f"识别成功: {result['success']}")
                    print(f"消息: {result['message']}")
                    
                    if result['data']:
                        print(f"识别到 {len(result['data'])} 个文本区域:")
                        for i, item in enumerate(result['data']):
                            print(f"  {i+1}. 文本: {item['text']}")
                            print(f"     置信度: {item['confidence']:.3f}")
                            print(f"     位置: {item['bbox']}")
                    
                    if result['model_info']:
                        print(f"模型信息: {result['model_info']}")
                else:
                    print(f"响应: {response.json()}")
                
                return response.status_code == 200
                
        except Exception as e:
            print(f"OCR上传测试失败: {str(e)}")
            return False
    
    def test_ocr_batch(self, image_paths):
        """测试批量OCR接口"""
        print(f"\n=== 测试批量OCR接口 ===")
        try:
            files = []
            for i, path in enumerate(image_paths):
                if os.path.exists(path):
                    files.append(('files', open(path, 'rb')))
            
            if not files:
                print("没有可用的测试图片")
                return False
            
            start_time = time.time()
            response = self.session.post(f"{self.base_url}/ocr/batch", files=files)
            end_time = time.time()
            
            # 关闭文件
            for _, f in files:
                f.close()
            
            print(f"状态码: {response.status_code}")
            print(f"响应时间: {end_time - start_time:.2f}秒")
            
            if response.status_code == 200:
                result = response.json()
                print(f"批量识别成功: {result['success']}")
                print(f"消息: {result['message']}")
                
                if result['results']:
                    for item in result['results']:
                        print(f"文件 {item['file_index']}: {item['filename']}")
                        print(f"  成功: {item['success']}")
                        if item['success']:
                            print(f"  文本数量: {item['text_count']}")
                        else:
                            print(f"  错误: {item['error']}")
            else:
                print(f"响应: {response.json()}")
            
            return response.status_code == 200
            
        except Exception as e:
            print(f"批量OCR测试失败: {str(e)}")
            return False
    
    def test_error_cases(self):
        """测试错误情况"""
        print(f"\n=== 测试错误情况 ===")
        
        # 测试无效文件类型
        print("1. 测试无效文件类型")
        try:
            with open("test.txt", "w") as f:
                f.write("This is not an image")
            
            with open("test.txt", "rb") as f:
                files = {'file': f}
                response = self.session.post(f"{self.base_url}/ocr/upload", files=files)
                print(f"状态码: {response.status_code}")
                print(f"响应: {response.json()}")
            
            os.remove("test.txt")
        except Exception as e:
            print(f"错误情况测试失败: {str(e)}")
    
    def run_all_tests(self):
        """运行所有测试"""
        print("开始运行API测试")
        print("=" * 50)
        
        results = []
        
        # 基础接口测试
        results.append(("健康检查", self.test_health_check()))
        results.append(("根路径", self.test_root_endpoint()))
        results.append(("模型信息", self.test_model_info()))
        results.append(("支持语言", self.test_languages()))
        
        # 创建测试图片
        test_image = self.create_test_image()
        if test_image:
            # OCR测试
            results.append(("OCR上传", self.test_ocr_upload(test_image)))
            
            # 批量OCR测试
            results.append(("批量OCR", self.test_ocr_batch([test_image, test_image])))
        
        # 错误情况测试
        self.test_error_cases()
        
        # 输出测试结果
        print("\n" + "=" * 50)
        print("测试结果汇总:")
        print("=" * 50)
        
        passed = 0
        total = len(results)
        
        for test_name, result in results:
            status = "✅ 通过" if result else "❌ 失败"
            print(f"{test_name}: {status}")
            if result:
                passed += 1
        
        print(f"\n总计: {passed}/{total} 个测试通过")
        
        if passed == total:
            print("🎉 所有测试通过！")
        else:
            print("⚠️  部分测试失败，请检查API服务状态")

def main():
    """主函数"""
    print("EasyOCR API 测试工具")
    print("=" * 50)
    
    # 检查API服务是否运行
    try:
        response = requests.get("http://localhost:8000/health", timeout=5)
        if response.status_code != 200:
            print("❌ API服务未正常运行，请先启动服务:")
            print("   python api_server.py")
            return
    except requests.exceptions.ConnectionError:
        print("❌ 无法连接到API服务，请先启动服务:")
        print("   python api_server.py")
        return
    
    print("✅ API服务正在运行")
    
    # 运行测试
    tester = APITester()
    tester.run_all_tests()

if __name__ == "__main__":
    main() 