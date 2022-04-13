package douban

import (
	"log"
	"net/http"
	"net/url"
)

//* 引用传递-指针到函数内，函数中参数修改，将影响实际函数
type CollectorOption func(*Collector)

type Collector struct {
	UserAgent  string
	Proxy      string
	Headers    *http.Header
	callback   func(*http.Response)
	ProxyTrans *http.Transport
}

func NewCollector(options ...CollectorOption) *Collector {
	//&用法，指向Collector的指针，Collector的地址
	c := &Collector{}
	c.Init()

	for _, f := range options {
		//f的用法
		f(c)
	}

	return c
}
func (c *Collector) Init() {
	c.UserAgent = "douban - user agent"
	c.Headers = nil
	c.ProxyTrans = nil
}

func UserAgent(ua string) CollectorOption {
	return func(c *Collector) {
		c.UserAgent = ua
	}
}
func ProxyTrans(p []string) CollectorOption {
	return func(c *Collector) {
		pp := p[0]
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(pp)
		}
		c.ProxyTrans = &http.Transport{
			Proxy: proxy,
		}
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
	client := &http.Client{
		Transport: c.ProxyTrans,
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	c.handleresponse(resp)

	return err
}

func (c *Collector) OnResponse(f func(r *http.Response)) {
	c.callback = f
}

func (c *Collector) handleresponse(r *http.Response) {
	c.callback(r)
}
