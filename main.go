package main

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery" //github packages required for crawling and scraping
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector( //declaring colly variable function
		colly.AllowedDomains("gobyexample.com"), //target website
	)

	c.OnHTML("article", func(e *colly.HTMLElement) { //used to discriminate between meta tags (snippets of text)
		metaTags := e.DOM.ParentsUntil("`").Find("meta")
		metaTags.Each(func(_ int, s *goquery.Selection) {

		})
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) { //specifies the URL
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Limit(&colly.LimitRule{ //delay between crawl
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Crawl on Page", r.URL.String()) //prints URLs from the website
	})

	c.Visit("https://gobyexample.com/") //target site

}
