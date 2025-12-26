package main

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

func main() {

	t1, err := dateparse.ParseLocal("2025-01-10")
	if err != nil {
		fmt.Println(err)
	}

	t2 := time.Now()

	sub := t2.Sub(t1)

	fmt.Println(int(sub.Hours() / 24))

	//startTime := time.Now().Add(time.Millisecond * -11)
	//duration := time.Now().Sub(startTime)
	//fmt.Printf("%v", duration)
	//msInt := 115428
	//
	//duration := time.Duration(msInt * int(time.Millisecond))
	//fmt.Println(duration)
	//
	//time.Now().Format("2006-01-02 15:04:05")

	//fmt.Println((time.Second * 90).String())
	//
	//t := time.Now()
	//startOfDay := now.With(t).BeginningOfDay() // 当天零点
	//endOfMonth := now.With(t).EndOfMonth()     // 当月最后一天的23:59:59.999999999
	//fmt.Println("startOfDay: ", startOfDay, "endOfMonth: ", endOfMonth)
	//
	//// 时区转换
	//beijing := time.FixedZone("Beijing Time", 8*3600) // 东八区
	//tInBeijing := now.With(time.Now()).In(beijing)    // 转换为北京时间
	//fmt.Println("tInBeijing: ", tInBeijing)

	// 获取当前时间的周开始和结束（可自定义周起始日）
	//now.WeekStartDay = time.Monday
	//weekStart := now.BeginningOfWeek()
	//weekEnd := now.EndOfWeek()
	//fmt.Printf("本周范围: %v 至 %v\n", weekStart, weekEnd)
	//// 解析多种格式的时间字符串
	//ti, _ := now.Parse("2023-10-05 14:30")
	//fmt.Println("解析时间:", ti)

}
