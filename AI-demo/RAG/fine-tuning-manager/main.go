package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

// WebAPIManager 定义一个结构体来管理 Web API 进程
type WebAPIManager struct {
	process *os.Process
	cmd     *exec.Cmd
}

// 启动 Web API 服务
func (m *WebAPIManager) start() error {
	m.cmd = exec.Command("python", "app.py") // 假设你的 Python Web API 脚本名为 app.py
	m.process = m.cmd.Process
	fmt.Println("Starting Web API service...")
	return m.cmd.Start()
}

// 停止 Web API 服务
func (m *WebAPIManager) stop() error {
	if m.process != nil {
		fmt.Println("Stopping Web API service...")
		return m.process.Kill()
	}
	return nil
}

// 重启 Web API 服务
func (m *WebAPIManager) reload() error {
	if err := m.stop(); err != nil {
		return fmt.Errorf("failed to stop the process: %v", err)
	}
	time.Sleep(1 * time.Second) // 等待进程完全停止
	if err := m.start(); err != nil {
		return fmt.Errorf("failed to start the process: %v", err)
	}
	return nil
}

func main() {
	manager := WebAPIManager{}

	// 启动 Web API 服务
	if err := manager.start(); err != nil {
		fmt.Printf("Failed to start Web API service: %v\n", err)
		return
	}

	// 设置信号处理，监听 reload 指令
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP) // 使用 SIGHUP 信号作为 reload 指令

	for {
		select {
		case <-signalChan:
			fmt.Println("Received reload signal, restarting Web API service...")
			if err := manager.reload(); err != nil {
				fmt.Printf("Failed to reload Web API service: %v\n", err)
			}
		case <-time.After(1 * time.Hour): // 每小时检查一次，确保进程正常运行
			if manager.process == nil || manager.process.Pid == 0 {
				fmt.Println("Web API service is not running, restarting...")
				if err := manager.start(); err != nil {
					fmt.Printf("Failed to start Web API service: %v\n", err)
				}
			}
		}
	}
}
