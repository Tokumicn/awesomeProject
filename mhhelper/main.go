package main

import (
	"awesomeProject/mhhelper/mhjlreq"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	answerHtml, err := mhjlreq.GetMHJLAnswer(ctx, "力量套装")
	if err != nil {
		panic(err)
	}
	fmt.Println("返回结果: ", answerHtml)
}

func demo() {
	//ctx, cancel := chromedp.NewContext(
	//	context.Background(),
	//	// chromedp.WithDebugf(log.Printf),
	//)
	//defer cancel()
	//
	//ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	//defer cancel()
	//
	//var example string
	//err := chromedp.Run(ctx,
	//	chromedp.Navigate(`https://xyq.gm.163.com/sprite.html`),
	//	// wait for footer element is visible (ie, page is loaded)
	//	chromedp.WaitVisible(`body > footer`),
	//	chromedp.SetValue("ques", "力量套装")
	//	//// find and click "Example" link
	//	//chromedp.Click(`#example-After`, chromedp.NodeVisible),
	//	//// retrieve the text of the textarea
	//	//chromedp.Value(`#example-After textarea`, &example),
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Printf("Go's time.After example:\n%s", example)
}
