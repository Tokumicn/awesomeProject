package main

import "fmt"

func main() {

	test := []string{"1", "2", "3"}

	test2 := test[3:]

	fmt.Println(len(test), test2)
}
