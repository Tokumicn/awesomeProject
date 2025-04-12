package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"time"
)

// 定义模型
type User struct {
	gorm.Model
	Name string
	Age  int
	Desc string
}

func main() {

	db, err := gorm.Open("sqlite3", "/Users/zhangrui/GolandProjects/awesomeProject/gorm-old-demo/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// 创建记录
	newUser := User{Name: "John Doe", Age: 30}

	err = db.Create(newUser).Error
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	db.Callback().Query().Before("gorm:before_query").Register("beforeQuery", func(scope *gorm.Scope) {
		scope.Set("startTime", time.Now())
	})
	// 在 db.Callback() 注册 After 回调
	db.Callback().Query().After("gorm:after_query").Register("afterQuery", func(scope *gorm.Scope) {
		// 记录查询结束时间
		endTime := time.Now()
		// 获取开始时间
		startTimeInterface, _ := scope.Get("startTime")
		startTime, ok := startTimeInterface.(time.Time)
		if !ok {
			panic("startTimeInterface is not time.Time.")
		}
		// 计算查询执行时长
		duration := endTime.Sub(startTime)
		// 记录查询时长
		log.Printf("Query took %v", duration)
	})

	db.DB().Stats()

}

// beforeQuery 是 Query 操作的 Before 回调函数
func beforeQuery_Old(db *gorm.DB) {
	// 记录查询开始时间
	db.InstantSet("startTime", time.Now())
}

// 你可以添加一个 After 回调来记录查询结束时间和执行时长
func afterQuery_Old(db *gorm.DB) {
	// 记录查询结束时间
	endTime := time.Now()
	// 获取开始时间
	startTimeInterface, _ := db.Get("startTime")
	startTime, ok := startTimeInterface.(time.Time)
	if !ok {
		panic("startTimeInterface is not time.Time.")
	}
	// 计算查询执行时长
	duration := endTime.Sub(startTime)
	// 记录查询时长
	log.Printf("Query took %v", duration)
}
