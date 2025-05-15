package main

import (
	"fmt"
	"time"
)

func main() {
	method1()
	method2()
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
