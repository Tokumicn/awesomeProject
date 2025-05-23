package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	//decodeBytesDemo()
	//decodeStringDemo()
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("当前目录：", dir)

	// 假设我们要检查的图片文件名是 "image.jpg"
	imagePath := fmt.Sprintf("%s/1.jpeg", dir)

	// 检查文件是否存在
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		fmt.Printf("文件 %s 不存在\n", imagePath)
		return
	}

	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("file size: ", float64(info.Size())/1024)

	// 读取本地文件
	bytes, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// 将字节切片转换为 base64 编码的字符串
	base64String := base64.StdEncoding.EncodeToString(bytes)

	fmt.Println(((len(base64String) * 3) / 4) / 1024)
	fmt.Println(len(base64String) / 1024)
}

func decodeBytesDemo() {
	srcBytes := []byte("hello world")
	// 根据源文件 bytes 的大小通过 EncodedLen 方法 返回目标bytes大小
	dstBytes := make([]byte, base64.StdEncoding.EncodedLen(len(srcBytes)))
	base64.StdEncoding.Encode(dstBytes, srcBytes)
	fmt.Printf("encode(`hello world`) = %s\n", string(dstBytes))

	decodeDstBytes := make([]byte, base64.StdEncoding.DecodedLen(len(dstBytes)))
	base64.StdEncoding.Decode(decodeDstBytes, dstBytes)
	fmt.Printf("decode(`%s`) = %s\n", string(dstBytes), string(decodeDstBytes))
}

func decodeStringDemo() {
	s := "hello world"
	sEncode := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Printf("encode(`hello world`) = %s\n", sEncode)

	sDecode, err := base64.StdEncoding.DecodeString(sEncode)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("decode(`%s`) = %s\n", sEncode, sDecode)
	}
}
