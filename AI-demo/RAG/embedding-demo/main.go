package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:11434/api/embed"
	method := "POST"

	payload := strings.NewReader(`{
  "model": "qwen3:0.6b",  
  "input": "我爱北京天安门"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
