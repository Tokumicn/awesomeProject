package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
)

func main() {
	//startTime := time.Now().Add(time.Millisecond * -11)
	//duration := time.Now().Sub(startTime)
	//fmt.Printf("%v", duration)

	fmt.Println((time.Second * 90).String())

	t := time.Now()
	startOfDay := now.With(t).BeginningOfDay() // 当天零点
	endOfMonth := now.With(t).EndOfMonth()     // 当月最后一天的23:59:59.999999999
	fmt.Println("startOfDay: ", startOfDay, "endOfMonth: ", endOfMonth)

	// 时区转换
	beijing := time.FixedZone("Beijing Time", 8*3600) // 东八区
	tInBeijing := now.With(time.Now()).In(beijing)    // 转换为北京时间
	fmt.Println("tInBeijing: ", tInBeijing)

	// 获取当前时间的周开始和结束（可自定义周起始日）
	now.WeekStartDay = time.Monday
	weekStart := now.BeginningOfWeek()
	weekEnd := now.EndOfWeek()
	fmt.Printf("本周范围: %v 至 %v\n", weekStart, weekEnd)
	// 解析多种格式的时间字符串
	ti, _ := now.Parse("2023-10-05 14:30")
	fmt.Println("解析时间:", ti)

}
