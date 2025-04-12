package main

import (
	"database/sql"
	_ "embed"
	_ "github.com/mattn/go-sqlite3" // 使用SQLite3的驱动
	"os"
)

//go:embed mydatabase.db
var db []byte

func openDatabase() (*sql.DB, error) {
	// 创建一个临时文件
	tempFile, err := os.CreateTemp("", "mydatabase-temp.db")
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	// 将数据库文件内容写入临时文件
	if _, err := tempFile.Write(db); err != nil {
		os.Remove(tempFile.Name())
		return nil, err
	}
	if err := tempFile.Close(); err != nil {
		os.Remove(tempFile.Name())
		return nil, err
	}

	// 打开临时数据库文件
	return sql.Open("sqlite3", tempFile.Name())
}
