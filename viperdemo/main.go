package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var cfg = flag.String("c", "", "config file")

type Config struct {
	Username string
	Password string
	// Viper 支持嵌套结构体
	Server struct {
		IP   string
		Port int
	}
}

func main() {
	//// 设置默认配置
	//viper.SetDefault("username", "jianghushinian")
	//viper.SetDefault("server", map[string]string{"ip": "127.0.0.1", "port": "8080"})
	//
	//// 读取配置值
	//fmt.Printf("username: %s\n", viper.Get("Username")) // key 不区分大小写
	//fmt.Printf("server: %+v\n", viper.Get("server"))

	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config file error: %s \n", err.Error()))
	}

	// 将配置信息反序列化到结构体中
	var cfg *Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("unmarshal config error: %s \n", err.Error()))
	}

	// 注册每次配置文件发生变更后都会调用的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 每次配置文件发生变化，需要重新将其反序列化到结构体中
		if err := viper.Unmarshal(&cfg); err != nil {
			panic(fmt.Errorf("unmarshal config error: %s \n", err.Error()))
		}
	})

	// 监控配置文件变化
	viper.WatchConfig()

	// use config...
	fmt.Println(cfg.Username)
}

func demo() {
	flag.Parse()

	if *cfg != "" {
		viper.SetConfigFile(*cfg)   // 指定配置文件（路径 + 配置文件名）
		viper.SetConfigType("yaml") // 如果配置文件名中没有扩展名，则需要显式指定配置文件的格式
	} else {
		viper.AddConfigPath(".")             // 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath("$HOME/.config") // 可以多次调用 AddConfigPath 来设置多个配置文件搜索路径
		viper.SetConfigName("cfg")           // 指定配置文件名（没有扩展名）
	}

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(errors.New("config file not found"))
		} else {
			fmt.Println(errors.New("config file was found but another error was produced"))
		}
		return
	}

	fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())

	// 读取配置值
	fmt.Printf("username: %s\n", viper.Get("username"))
}
