package douban

import (
	"errors"
	"net/http"
)

type CollectorOption func(*Collector)

type Collector struct {
	UserAgent string
	Headers   *http.Header
}

func NewCollector(options ...CollectorOption) *Collector {
	//&用法
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
	c.UserAgent = "douban"
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
	return errors.New("imcomplete")
}
