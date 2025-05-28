package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {

	g, ctx := errgroup.WithContext(context.Background())

	// 启动多个 Goroutine
	g.Go(func() error {
		// 任务 1
		select {
		case <-ctx.Done():
			return ctx.Err() // 若其他任务出错，这里会收到取消信号
		default:
			fmt.Println("Task 1")
			return nil
		}
	})

	g.Go(func() error {
		// 任务 2（模拟出错）
		return fmt.Errorf("task 2 failed")
	})

	// 等待所有任务完成，并返回首个错误
	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err) // 输出: Error: task 2 failed
	}

}

func errorDemo() {
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
