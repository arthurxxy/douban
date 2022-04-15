package main

import (
	"douban/book"
	"douban/config"
	douban "douban/http"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	//var a []string //slice ,而数组需要指定元素及个数 var balance [10] float32
	// a := config.Conf.GetStringSlice("ini.proxy")
	// typeOfA := reflect.TypeOf(a)
	// if a != "" {
	// 	fmt.Println(typeOfA.Name(), typeOfA.Kind(), len(a))
	// }
	//fmt.Print(a)
	//rand.Seed(time.Now().UnixNano())
	//slice := []int{0, 2}
	//rm := rand.Intn(len(slice))
	//slice = append(slice[:rm], slice[rm+1:]...)
	//fmt.Println(rm, slice[rm-1])
	// var a int
	// var r uint32 = 3
	// var s = []string{"a", "b", "c"}
	// for a < 10 {

	// 	index := atomic.AddUint32(&r, 1) - 1
	// 	i := index % uint32(len(s))
	// 	a++
	// 	fmt.Println(index, i)
	// }

	var proxys = config.Conf.GetStringSlice("ini.proxy")
	var delay = config.Conf.GetInt("ini.delay")
	var url = config.Conf.GetString("ini.url")
	var fromid = config.Conf.GetInt("ini.fromid")
	// c := douban.NewCollector(
	// 	douban.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"),
	// 	douban.DBProxy(proxys),
	// )
	// c.OnResponse(func(r *http.Response) {
	// 	book.Getbook(r)
	// 	//fmt.Print(c.)

	// })
	// for i := 0; i < 10; i++ {
	// 	c.Visit("http://www.ip111.cn/")
	// 	fmt.Println(i)
	// }
	c := douban.NewCollector(
		douban.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"),
		douban.DBProxy(proxys),
	)

	c.OnResponse(func(r *http.Response) {
		books := book.Getdouban(r)
		if books["书名"] != "" {
			log.Println("get books detail success:", books["ID"], "-", books["书名"])

		} else {
			log.Println("get books detail error!")
		}
		log.Println("use proxy:", c.CurrentProxy)
		time.Sleep(time.Duration(delay) * time.Second)
		fromid++
		c.Visit(fmt.Sprintf(url, fromid))
	})

	c.Visit(fmt.Sprintf(url, fromid))
}
