package main

import (
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
)

const htmlTemplate = `<!DOCTYPE html><html><head><title>Welcome</title></head><body><h1>Hello {{.Name}},</h1><p>Thank you for signing up! Your verification code is: <strong>{{.Code}}</strong></p><p>Click <a href="https://example.com/verify?code={{.Code}}">here</a> to verify your account.</p></body></html>`

func main() {
	// 定义模板数据
	data := struct {
		Name string
		Code string
	}{
		Name: "Alice",
		Code: "123456",
	}
	// 解析模板
	tmpl, err := template.New("email").Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}
	// 渲染模板
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		panic(err)
	}
	// 发送邮件
	sendEmail("461694377@qq.com", "Welcome to Our Service", body.String())
}

// 发送邮件的函数
func sendEmail(to, subject, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "461694377@qq.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body) // 发送 HTML 邮件
	d := gomail.NewDialer("smtp.qq.com", 465, "461694377@qq.com", "8bvjiefvpsseebhee77")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
