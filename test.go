package main

import (
	"fmt"
	"math/rand"
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
	rand.Seed(time.Now().UnixNano())
	slice := []int{0, 2}
	rm := rand.Intn(len(slice))
	//slice = append(slice[:rm], slice[rm+1:]...)
	fmt.Println(rm, slice[rm-1])

	// var proxys = config.Conf.GetStringSlice("ini.proxy")
	// c := douban.NewCollector(
	// 	douban.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"),
	// 	douban.ProxyTrans(proxys),
	// )
	// c.OnResponse(func(r *http.Response) {
	// 	book.Getbook(r)
	// 	fmt.Print(c.ProxyTrans)
	// })
	// c.Visit("http://www.ip111.cn/")
}
