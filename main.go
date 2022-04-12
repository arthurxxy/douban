package main

import (
	"douban/config"
	douban "douban/http"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

//显示内存调用，及回收
func printMemStats() {

	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)

}

func main() {
	//调试内存使用
	go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()

	fmt.Println(config.Conf.GetInt("ini.delay"))
	fmt.Println(config.Conf.Get("ini.proxy"))
	c := douban.NewCollector(
		douban.UserAgent("123"),
	)
	var err error
	for i := 0; i < 1000; i++ {
		m := make(map[int]string)
		m[i] = "123"
		err = c.Visit("http://ip111.cn")
		time.Sleep(100 * time.Millisecond)
		//runtime.GC()
		printMemStats()
	}

	fmt.Println(err)
}
