package main

import (
	"context"
	"douban/book"
	"douban/config"
	douban "douban/http"
	"douban/mongo"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

//显示内存调用，及回收
// func printMemStats() {
// 	var m runtime.MemStats
// 	runtime.ReadMemStats(&m)
// 	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
// }

func main() {
	//调试内存使用
	// go func() {
	// 	http.ListenAndServe("0.0.0.0:8080", nil)
	// }()

	var delay = config.Conf.GetInt("ini.delay")
	var url = config.Conf.GetString("ini.url")
	var fromid = config.Conf.GetInt("ini.fromid")
	var mongodb = config.Conf.GetString("db.mongodb")
	var database = config.Conf.GetString("db.database")
	var collection = config.Conf.GetString("db.collection")
	var proxys = config.Conf.GetStringSlice("ini.proxy")

	//connect to mongodb
	coll, err := mongo.ConnectMongo(mongodb, database, collection)
	if err != nil {
		log.Println("mongo connect error:" + err.Error())
	}

	c := douban.NewCollector(
		douban.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"),
		douban.DBProxy(proxys),
	)

	c.OnResponse(func(r *http.Response) {
		books := book.Getdouban(r)
		if books["书名"] != "" {
			log.Println("get books detail success:", books["ID"], "-", books["书名"])
			result, err := coll.InsertOne(context.Background(), books)
			if err == nil {
				log.Println("mongo insert success,record is", result.InsertedID)
			} else {
				log.Println("mongo insert error!")
			}
		} else {
			log.Println("get books detail error!")
		}
		time.Sleep(time.Duration(delay) * time.Second)
		fromid++
		c.Visit(fmt.Sprintf(url, fromid))
	})
	if fromid == 0 {
		fromid = mongo.GetMaxID(mongodb, database, collection) + 1
	}

	c.Visit(fmt.Sprintf(url, fromid))

}
