package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	decodeBytesDemo()
	decodeStringDemo()
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
