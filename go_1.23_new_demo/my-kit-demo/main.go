package main

import (
	"awesomeProject/go_1.23_new_demo/my-kit-demo/log"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cast"
)

const (
	LogCfgKeyLevel  = "level"
	LogCfgKeyWriter = "writer"
	LogCfgKeyStore  = "store"

	QuitKey = "quit"
	ExitKey = "exit"
)

func main() {

	conf := log.LoggerConfig{}
	log.NewLogger(conf)

	log.SLog.Info("开始")

	reader := bufio.NewReader(os.Stdin)
	for {
		// 读取用户名
		fmt.Println("enter log config change: ")
		key, _ := reader.ReadString('\n')
		key = strings.TrimSpace(key)

		switch key {

		case LogCfgKeyLevel:
			fmt.Printf("输入对配置[%s]的更改: ", LogCfgKeyLevel)
			level, _ := reader.ReadString('\n')
			level = strings.TrimSpace(level)
			intLevel, err := cast.ToIntE(level)
			if err != nil {
				log.SLog.Error("cast.ToInt[%s] err: %v", level, err)
			}
			log.SetLoggerLevel(intLevel)

		case LogCfgKeyWriter:
			fmt.Printf("输入对配置[%s]的更改(Text|Json): ", LogCfgKeyWriter)
			writerTyp, _ := reader.ReadString('\n')
			conf.WriterType = writerTyp
			log.NewLogger(conf)

		case QuitKey, ExitKey:
			goto gotoFlag

		default:
			fmt.Printf("输入配置项有误[%s]\n", key)
		}
	}
gotoFlag:
	log.SLog.Info("结束了")
	log.SLog.Debug("[DEBUG] 结束了")

	fmt.Println("结束!!!")
}
