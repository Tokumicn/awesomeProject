package main

import (
	"fmt"
	"html/template"
)

func main() {
	tmpl := `<html><body><h1>{{.Title}}</h1><p>{{.Content}}</p></body></html>`
	t, _ := template.New("demo").Parse(tmpl)
	fmt.Println(t)
}
