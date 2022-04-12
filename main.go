package main

import (
	"douban/config"
	douban "douban/http"
	"fmt"
)

func main() {
	fmt.Println(config.Conf.GetInt("ini.delay"))
	fmt.Println(config.Conf.Get("ini.proxy"))
	c := douban.NewCollector(
		douban.UserAgent("123"),
	)
	err := c.Visit("123")

	fmt.Println(err)
}
