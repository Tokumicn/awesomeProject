package main

import (
	"fmt"
	"regexp"
)

func main() {

	text1 := "^(?:你好|大家好|哈喽).*(?:今天|明天|后天).*(?:吗|哈哈|嘿)$"
	//text2 := "^(你好|大家好|哈喽).*(今天|明天|后天).*(吗|哈哈|嘿)$"
	compile, err := regexp.Compile(text1)
	if err != nil {
		panic(err)
	}

	if matched := compile.MatchString("你好今天吗"); matched {
		fmt.Println("通过")
	}

	if matched := compile.MatchString("今天吗"); matched {
		fmt.Println("通过")
	}

	if matched := compile.MatchString("大家好明天哈哈"); matched {
		fmt.Println("通过")
	}

	if matched := compile.MatchString("哈喽后天嘿"); matched {
		fmt.Println("通过")
	}

	if matched := compile.MatchString("你好今天天气好吗"); matched {
		fmt.Println("通过")
	}

}
