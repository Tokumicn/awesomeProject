package main

import (
	"fmt"
	"strings"
)

func main() {
	//trimSuffix := strings.TrimSuffix("2025-12-100#", "1#")
	//println(trimSuffix)
	//
	//trimSuffix = strings.TrimSuffix("2025-12-100#", "0#")
	//println(trimSuffix)

	str := strings.TrimSpace("123456\r\n")
	fmt.Println(str)

	str = strings.TrimSpace("123\r\n456\r\n")
	fmt.Println(str)

	str = strings.TrimSpace("123\n456\r\n")
	fmt.Println(str)
}
