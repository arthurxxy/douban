package douban

import (
	"log"
	"net/http"
	"net/url"
	"sync/atomic"
)

//* 引用传递-指针到函数内，函数中参数修改，将影响实际函数
type CollectorOption func(*Collector)

type Collector struct {
	UserAgent    string
	Headers      *http.Header
	callback     func(*http.Response)
	DBProxy      []string
	proxyindex   uint32
	CurrentProxy string
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
	c.proxyindex = 0
}

func UserAgent(ua string) CollectorOption {
	return func(c *Collector) {
		c.UserAgent = ua
	}
}
func DBProxy(p []string) CollectorOption {
	return func(c *Collector) {
		c.DBProxy = p
	}
}

func (c *Collector) switchProxy() *url.URL {
	if len(c.DBProxy) < 1 {
		return nil
	} else {
		index := atomic.AddUint32(&c.proxyindex, 1) - 1
		sURL := c.DBProxy[index%uint32(len(c.DBProxy))]
		c.CurrentProxy = sURL
		u, _ := url.Parse(sURL)
		return u
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
	sp := c.switchProxy()
	tr := &http.Transport{
		Proxy: http.ProxyURL(sp),
		//DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: tr,
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
