#!/usr/bin/env python3
"""
EasyOCR API 服务器启动脚本
自动检查环境并启动服务
"""

import os
import sys
import subprocess
import time
import requests
from pathlib import Path

def check_python_version():
    """检查Python版本"""
    print("检查Python版本...")
    if sys.version_info < (3, 8):
        print("❌ Python版本过低，需要Python 3.8或更高版本")
        print(f"当前版本: {sys.version}")
        return False
    print(f"✅ Python版本: {sys.version.split()[0]}")
    return True

def check_dependencies():
    """检查依赖包"""
    print("\n检查依赖包...")
    required_packages = [
        'torch', 'transformers', 'Pillow', 'fastapi', 
        'uvicorn', 'easyocr', 'numpy', 'opencv-python'
    ]
    
    missing_packages = []
    for package in required_packages:
        try:
            __import__(package)
            print(f"✅ {package}")
        except ImportError:
            print(f"❌ {package} - 未安装")
            missing_packages.append(package)
    
    if missing_packages:
        print(f"\n缺少以下依赖包: {', '.join(missing_packages)}")
        print("请运行以下命令安装:")
        print("pip install -r requirements.txt")
        return False
    
    return True

def check_torch_mps():
    """检查PyTorch MPS支持"""
    print("\n检查PyTorch MPS支持...")
    try:
        import torch
        print(f"PyTorch版本: {torch.__version__}")
        
        if torch.backends.mps.is_available():
            print("✅ MPS (Metal Performance Shaders) 可用")
            print("✅ 将使用Mac M3芯片GPU加速")
            return True
        else:
            print("⚠️  MPS不可用，将使用CPU模式")
            return True
    except Exception as e:
        print(f"❌ PyTorch检查失败: {str(e)}")
        return False

def check_network():
    """检查网络连接"""
    print("\n检查网络连接...")
    try:
        response = requests.get("https://huggingface.co", timeout=10)
        if response.status_code == 200:
            print("✅ 网络连接正常")
            return True
        else:
            print("⚠️  网络连接异常")
            return False
    except Exception as e:
        print(f"❌ 网络连接失败: {str(e)}")
        print("首次运行需要下载模型，请确保网络连接正常")
        return False

def create_models_directory():
    """创建模型存储目录"""
    print("\n检查模型存储目录...")
    models_dir = Path("./models")
    if not models_dir.exists():
        models_dir.mkdir(exist_ok=True)
        print("✅ 创建模型存储目录: ./models")
    else:
        print("✅ 模型存储目录已存在")
    return True

def start_server():
    """启动API服务器"""
    print("\n启动EasyOCR API服务器...")
    print("=" * 50)
    
    try:
        # 启动服务器
        cmd = [sys.executable, "api_server.py"]
        process = subprocess.Popen(cmd)
        
        print("🚀 服务器正在启动...")
        print("📝 日志信息:")
        
        # 等待服务器启动
        time.sleep(3)
        
        # 检查服务器是否成功启动
        try:
            response = requests.get("http://localhost:8000/health", timeout=10)
            if response.status_code == 200:
                print("\n✅ 服务器启动成功！")
                print("\n📋 服务信息:")
                print("   🌐 服务地址: http://localhost:8000")
                print("   📚 API文档: http://localhost:8000/docs")
                print("   📖 交互文档: http://localhost:8000/redoc")
                print("   🔍 健康检查: http://localhost:8000/health")
                print("\n💡 使用提示:")
                print("   - 首次运行会自动下载模型，请耐心等待")
                print("   - 按 Ctrl+C 停止服务器")
                print("   - 运行 python test_api.py 测试API功能")
                
                # 等待用户中断
                try:
                    process.wait()
                except KeyboardInterrupt:
                    print("\n🛑 正在停止服务器...")
                    process.terminate()
                    process.wait()
                    print("✅ 服务器已停止")
                    
            else:
                print("❌ 服务器启动失败")
                process.terminate()
                return False
                
        except requests.exceptions.ConnectionError:
            print("❌ 无法连接到服务器")
            process.terminate()
            return False
            
    except Exception as e:
        print(f"❌ 启动服务器失败: {str(e)}")
        return False

def main():
    """主函数"""
    print("EasyOCR API 服务器启动器")
    print("=" * 50)
    
    # 环境检查
    checks = [
        ("Python版本", check_python_version),
        ("依赖包", check_dependencies),
        ("PyTorch MPS", check_torch_mps),
        ("网络连接", check_network),
        ("模型目录", create_models_directory)
    ]
    
    failed_checks = []
    for check_name, check_func in checks:
        if not check_func():
            failed_checks.append(check_name)
    
    if failed_checks:
        print(f"\n❌ 环境检查失败: {', '.join(failed_checks)}")
        print("请解决上述问题后重新运行")
        return False
    
    print("\n✅ 所有环境检查通过")
    
    # 启动服务器
    return start_server()

if __name__ == "__main__":
    try:
        success = main()
        if not success:
            sys.exit(1)
    except KeyboardInterrupt:
        print("\n👋 用户中断，退出程序")
        sys.exit(0)
    except Exception as e:
        print(f"\n❌ 程序异常: {str(e)}")
        sys.exit(1) 