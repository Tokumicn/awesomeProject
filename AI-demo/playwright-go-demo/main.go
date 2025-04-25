package main

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://mp.weixin.qq.com/s/m4dpXYoCedVVqaJJmvToVQ"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	content, err := page.Content()
	if err != nil {
		log.Fatalf("could not get content: %v", err)
	}

	fmt.Println(content)

	//entries, err := page.Locator("span leaf").All()
	//if err != nil {
	//	log.Fatalf("could not get entries: %v", err)
	//}

	entries, err := page.Locator(".athing").All()
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	for i, entry := range entries {
		title, err := entry.Locator("td.title > span > a").TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}
		fmt.Printf("%d: %s\n", i+1, title)
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
