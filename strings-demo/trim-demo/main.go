package main

import "strings"

func main() {
	trimSuffix := strings.TrimSuffix("2025-12-100#", "1#")
	println(trimSuffix)

	trimSuffix = strings.TrimSuffix("2025-12-100#", "0#")
	println(trimSuffix)
}
