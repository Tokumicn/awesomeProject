package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "dataSourceName")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
