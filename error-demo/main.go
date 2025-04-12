package main

import (
	"errors"
	"fmt"
)

func main() {

	// 假设我们有一个错误
	var err error = errors.New("error 11111.")

	// 将错误封装到 interface{} 类型
	var boxedInterface interface{} = err

	// 使用类型断言来检查 interface{} 变量中是否包含 error
	if typedErr, ok := boxedInterface.(error); ok {
		fmt.Println("Interface contains an error:", typedErr.Error())
	} else {
		fmt.Println("Interface does not contain an error")
	}

	if typedErr, ok := boxedInterface.(string); ok {
		fmt.Println("Interface contains an error:", typedErr)
	} else {
		fmt.Println("Interface does not contain an error")
	}
}
