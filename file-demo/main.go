package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("当前目录：", dir)

	// 假设我们要检查的图片文件名是 "image.jpg"
	imagePath := "1.png"

	// 检查文件是否存在
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		fmt.Printf("文件 %s 不存在\n", imagePath)
		return
	}

	// 创建备份目录
	backupDir := ".bak"
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		err := os.Mkdir(backupDir, 0755)
		if err != nil {
			fmt.Printf("创建备份目录失败: %v\n", err)
			return
		}
	}

	// 构建备份文件的路径
	backupPath := filepath.Join(backupDir, filepath.Base(imagePath))

	// 复制文件到备份目录
	err = copyFile(imagePath, backupPath)
	if err != nil {
		fmt.Printf("备份文件失败: %v\n", err)
		return
	}

	// 删除原文件
	err = os.Remove(imagePath)
	if err != nil {
		fmt.Printf("删除原文件失败: %v\n", err)
		return
	}

	fmt.Println("文件备份并删除成功")
}

// copyFile 复制文件内容
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Close()
}
