package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/sagikazarmark/slog-shim"
	"strconv"
	"strings"
)

type Repository struct {
	Author  string
	Name    string
	Link    string
	Desc    string
	Lang    string
	Stars   int
	Forks   int
	Add     int
	BuiltBy []string
}

func main() {
	GetGitHubTreading()
}

func GetGitHubTreading() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	extensions.RandomUserAgent(c)

	repos := make([]*Repository, 0, 15)
	c.OnHTML(".Box .Box-row", func(e *colly.HTMLElement) {
		repo := &Repository{}

		// author & repository name
		linkTexts := e.ChildTexts("a.Link")
		var parts []string
		if len(linkTexts) > 0 {
			authorRepoName := linkTexts[0]
			parts = strings.Split(authorRepoName, "/")
			repo.Author = strings.TrimSpace(parts[0])
			repo.Name = strings.TrimSpace(parts[1])
		}

		// link
		repo.Link = e.Request.AbsoluteURL(e.ChildAttr("a.Link", "href"))

		// description
		repo.Desc = e.ChildText("p.pr-4")

		// language
		repo.Lang = strings.TrimSpace(e.ChildText("div.mt-2 > span.mr-3 > span[itemprop]"))

		// star & fork
		starForkStr := e.ChildText("div.mt-2 > a.mr-3")
		starForkStr = strings.Replace(strings.TrimSpace(starForkStr), ",", "", -1)
		parts = strings.Split(starForkStr, "\n")
		repo.Stars, _ = strconv.Atoi(strings.TrimSpace(parts[0]))
		repo.Forks, _ = strconv.Atoi(strings.TrimSpace(parts[len(parts)-1]))

		// add
		addStr := e.ChildText("div.mt-2 > span.float-sm-right")
		parts = strings.Split(addStr, " ")
		repo.Add, _ = strconv.Atoi(parts[0])

		// built by
		e.ForEach("div.mt-2 > span.mr-3  img[src]", func(index int, img *colly.HTMLElement) {
			repo.BuiltBy = append(repo.BuiltBy, img.Attr("src"))
		})

		repos = append(repos, repo)
	})

	err := c.Visit("https://github.com/trending")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	fmt.Printf("%d repositories\n", len(repos))
	fmt.Println("first repository:")
	for _, repo := range repos {
		fmt.Printf("%+v \n", repo)
	}
}

func demo() {
	c := colly.NewCollector(colly.AllowedDomains("www.baidu.com"))

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		err := c.Visit(e.Request.AbsoluteURL(link))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response %s: %d bytes\n", r.Request.URL, len(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
	})

	c.Visit("http://www.baidu.com/")
}
