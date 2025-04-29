package main

import (
	"fmt"
	"github.com/tsenart/vegeta/v12/lib"
	"net/http"
	"os"
	"time"
)

func main() {
	// 定义目标
	target := "http://localhost:8080"

	// 创建 Vegeta 负载测试器
	attacker := vegeta.NewAttacker()

	// 自定义头信息
	headers := http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer " + os.Getenv("TOKEN")},
		"Cookie":        []string{os.Getenv("COOKIE")},
	}

	// 定义请求参数
	rate := vegeta.Rate{Freq: 10, Per: time.Second} // 每秒10个请求
	duration := 10 * time.Second                    // 测试持续10秒

	// 开始负载测试
	trGet := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target,
		Header: headers,
	})

	//trPost := vegeta.NewStaticTargeter(vegeta.Target{
	//	Method: "POST",
	//	URL:    target,
	//	Header: headers,
	//	Body:   []byte{}, // TODO body参数
	//})

	for res := range attacker.Attack(trGet, rate, duration, "demo-test") {
		fmt.Printf("请求 %s 状态码: %d\n", res.URL, res.Code)
	}

	// 完成后输出结果
	fmt.Println("负载测试完成")
}
