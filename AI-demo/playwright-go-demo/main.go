package main

import (
	"github.com/playwright-community/playwright-go"
	"log"
)

func main() {
	err := playwright.Install()
	if err != nil {
		log.Fatal(err)
	}
}
