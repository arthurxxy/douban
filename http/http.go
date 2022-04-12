package douban

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//* 引用传递-指针到函数内，函数中参数修改，将影响实际函数
type CollectorOption func(*Collector)

type Collector struct {
	UserAgent string
	Headers   *http.Header
}

func NewCollector(options ...CollectorOption) *Collector {
	//&用法，指向Collector的指针，Collector的地址
	c := &Collector{}
	c.Init()

	for _, f := range options {
		//f的用法
		f(c)
	}

	//c.parseSettingsFromEnv()

	return c
}
func (c *Collector) Init() {
	c.UserAgent = "douban - user agent"
	c.Headers = nil
}

func UserAgent(ua string) CollectorOption {
	return func(c *Collector) {
		c.UserAgent = ua
	}
}

func Headers(headers map[string]string) CollectorOption {
	return func(c *Collector) {
		custom_headers := make(http.Header)
		for header, value := range headers {
			custom_headers.Add(header, value)
		}
		c.Headers = &custom_headers
	}
}

func (c *Collector) Visit(URL string) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return getbook(resp)
}
func getbook(resp *http.Response) error {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("div.card-body").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("p").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
	return err
}
