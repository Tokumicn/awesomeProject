package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// 1. 读取本地图片
	imageBytes, err := os.ReadFile("./file-demo/.bak/1.png")
	if err != nil {
		log.Fatal("读取文件失败:", err)
	}

	// 2. 检测MIME类型（自动识别PNG/JPEG等）
	mimeType := http.DetectContentType(imageBytes)

	// 3. 构建数据URI
	dataURI := fmt.Sprintf("data:%s;base64,%s",
		mimeType,
		base64.StdEncoding.EncodeToString(imageBytes),
	)

	tempPrintStr := dataURI[:256]
	fmt.Println(tempPrintStr) // 输出结果示例：data:image/png;base64,iVBORw0KG...
}
