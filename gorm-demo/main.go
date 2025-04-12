package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {

	//newLogger := logger.New(
	//	log.New(os.Stdout, "/r/n", log.LstdFlags),
	//	logger.Config{
	//		SlowThreshold: 5 * time.Second,
	//		LogLevel:      logger.Info,
	//		Colorful:      true,
	//	},
	//)

	// 连接到数据库，这里以 SQLite 为例
	db, err := gorm.Open(sqlite.Open("/Users/zhangrui/GolandProjects/awesomeProject/gorm-old-demo/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 设置数据库连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(60 * time.Second)
	sqlDB.SetConnMaxIdleTime(60 * time.Second)

	// 数据库连接池实时状态查询
	sqlDB.Stats()

	// 定义模型
	type User struct {
		Name string
		Age  int
		Desc string
	}

	// 创建记录
	name := fmt.Sprintf("John Doe%d", time.Now().Unix())
	newUser := User{Name: name, Age: 30}

	db.DryRun = true // 注意：DryRun 不会返回错误，它只返回将要执行的 SQL
	// 使用 DryRun 生成并打印 SQL
	//sql := db.ToSQL(func(localDB *gorm.DB) *gorm.DB {
	//	return localDB.Where("userId IN (?) AND centerId = ?", []int64{111, 12, 32, 44}, 11).Find(&newUser)
	//})
	//fmt.Println("Generated SQL:")
	//fmt.Println(sql)

	err = db.Debug().Table("User").Create(newUser).Error
	if err != nil {
		log.Fatal(err)
		return
	}

	users := make([]User, 0)
	err = db.Debug().Table("User").Find(&users).Error
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(users)

}
