package main

import (
	"HqDistributedCrawler/engine"
	"HqDistributedCrawler/scheduler"
	"HqDistributedCrawler/zhengai/parser"
	"HqDistributedCrawler/distrubite/persist/client"
)

func main() {

	itemSaver,err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan:itemSaver,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParserCityList,
	})
}
