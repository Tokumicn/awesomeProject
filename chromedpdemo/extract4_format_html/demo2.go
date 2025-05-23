package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"golang.org/x/net/html"
	"log/slog"
	"strings"
)

func main() {
	htmlStr := `<a name=\"ask\" onclick=\"return reask('160级武器种类');\" href=\"#\" class=\"sprite_color_link\">160级武器种类</a><br>　　<span class=\"sprite_color_text\">属性范围</span><span class=\"sprite_color_text\">：<br>　　</span><span class=\"sprite_color_text\">强化打造</span><span class=\"sprite_color_text\">：命中(571-777)；伤害(490-667)<br>　　</span><span class=\"sprite_color_text\">获得</span><span class=\"sprite_color_text\">：<br>　　1、用<a name=\"ask\" onclick=\"return reask('陨铁');\" href=\"#\" class=\"sprite_color_link\">陨铁</a></span>幻化150级的强化武器(有几率失败)，得到<a name=\"ask\" onclick=\"return reask('元身');\" href=\"#\" class=\"sprite_color_link\">元身</a><br>　　2、用元身与<a name=\"ask\" onclick=\"return reask('战魄');\" href=\"#\" class=\"sprite_color_link\">战魄</a>进行打造，触发任务，完成后找袁天罡(长安城357，245)领取未鉴定武器<br>　　<span class=\"sprite_color_text\">其他说明</span><span class=\"sprite_color_text\">：<br>　　1、幻化150级武器的属性不影响元身属性(建议用</span><span class=\"sprite_color_text\">150级白板国标</span><span class=\"sprite_color_text\">武器幻化)<br>　　2、失败3次的武器无法幻化<br>　　3、角色≥120级才能使用元身和战魄打造160装备</span><span class=\"sprite_color_text\"></span>`

	ctx := context.TODO()
	parseCleanHtml, err := parseAndCleanAttributes(ctx, htmlStr)
	if err != nil {
		panic(err)
	}

	formattingText := extractFormattingText(ctx, parseCleanHtml)

	fmt.Println(formattingText)
}

// 对原始的 html 进行解析并清理无用属性
func parseAndCleanAttributes(ctx context.Context, htmlStr string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
	if err != nil {
		return "", err
	}

	// 清理属性
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		for _, node := range s.Nodes {
			cleanAttributes(node)
		}
	})

	cleanHtml, err := doc.Html()
	if err != nil {
		return "", err
	}

	slog.DebugContext(ctx, "clean html success. cleanHtml: ", cleanHtml)

	// 转换为纯文本
	return cleanHtml, nil
}

func cleanAttributes(n *html.Node) {
	if n.Type == html.ElementNode {
		n.Attr = nil // 清空所有属性
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cleanAttributes(c)
	}
}

// 保留 html 格式提取文本
func extractFormattingText(ctx context.Context, html string) string {
	browser := rod.New().ControlURL(launcher.New().
		Headless(true).
		Set("default-encoding", "utf-8"). // 关键设置
		MustLaunch()).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("data:text/html;charset=UTF-8," + html).MustWaitLoad()
	// 获取可视区域文本（自动处理 CSS 样式）
	element := page.MustElement("body")
	text := element.MustText()
	slog.InfoContext(ctx, "build html success. html2text: ", text)
	return text
}
