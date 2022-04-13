package book

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Getbook(resp *http.Response) {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("div.card-deck.mb-3text-center").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("div").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}

func Getdouban(resp *http.Response) map[string]string {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	mapbooks := make(map[string]string)

	// Find the review items
	doc.Find("div[id=wrapper]").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title

		mapbooks["书名"] = s.Find("h1 span").Text()
		mapbooks["评分"] = s.Find("strong.rating_num").Text()
		mapbooks["封面"], _ = s.Find("a.nbg").Attr("href")
		mapbooks["ID"] = resp.Request.URL.Path[9:16]
		mapbooks["简介"] = s.Find("div.intro").Text()
		mapbooks["created"] = time.Now().Format("20060102150405")
		bookinfo := s.Find("div#info").Text()
		bookinfo = strings.Replace(bookinfo, " ", "", -1)
		bookinfo = strings.Replace(bookinfo, "\n", "", -1)
		bookinfo = strings.Replace(bookinfo, "|", "", -1)
		bookinfo = strings.Replace(bookinfo, "出版社:", "|出版社:", 1)
		bookinfo = strings.Replace(bookinfo, "出版年:", "|出版年:", 1)
		bookinfo = strings.Replace(bookinfo, "页数:", "|页数:", 1)
		bookinfo = strings.Replace(bookinfo, "定价:", "|定价:", 1)
		bookinfo = strings.Replace(bookinfo, "装帧:", "|装帧:", 1)
		bookinfo = strings.Replace(bookinfo, "丛书:", "|丛书:", 1)
		bookinfo = strings.Replace(bookinfo, "副标题:", "|副标题:", 1)
		bookinfo = strings.Replace(bookinfo, "ISBN:", "|ISBN:", 1)
		bookinfo = strings.Replace(bookinfo, "译者:", "|译者:", 1)
		bookinfo = strings.Replace(bookinfo, "原作名:", "|原作名:", 1)
		bookinfo = strings.Replace(bookinfo, "出品方:", "|出品方:", 1)
		info := strings.Split(bookinfo, "|")
		for _, value := range info {
			if value != "" {
				mapbooks[strings.Split(value, ":")[0]] = strings.Split(value, ":")[1]
			}
		}
		//title := s.Find("div").Text()

	})
	return mapbooks
}
