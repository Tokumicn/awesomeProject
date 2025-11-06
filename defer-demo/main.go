package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//method1()
	//method2()

	ctx := context.WithValue(context.Background(), "YY", "YYValue")

	go func() {
		go goFunc(ctx)
	}()
	go func() {
		go goFunc2(ctx)
	}()
	go func() {
		go goFunc3(ctx)
	}()

	time.Sleep(time.Second * 3)
}

func getYY(ctx context.Context) string {
	return ctx.Value("YY").(string)
}

func goFunc(ctx context.Context) {
	go func() {
		go func() {
			fmt.Println("1 goroutine getYY:", getYY(ctx))
		}()
	}()
}

func goFunc2(ctx context.Context) {
	go func() {
		go func() {
			go goFunc(ctx)
		}()
	}()
}

func goFunc3(ctx context.Context) {
	go func() {
		go goFunc2(ctx)
	}()
}

func method1() {
	defer sinceTime("method1")()
	time.Sleep(1 * time.Second)
}

func method2() {
	defer sinceTime("method2")()
	time.Sleep(2 * time.Second)
}

func sinceTime(tag string) func() {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		fmt.Printf("tag [%s] 耗时：%v\n", tag, duration)
	}
}
