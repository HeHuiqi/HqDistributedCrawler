package main

import (
	"HqDistributedCrawler/engine"
	"HqDistributedCrawler/scheduler"
	"HqDistributedCrawler/zhengai/parser"
	"HqDistributedCrawler/persist"
)
//单机版
func main() {

	itemSaver,err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan:itemSaver,
		RequestProcessor:engine.Worker,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser:engine.NewFuncParser(parser.ParserCityList,"ParserCityList"),
	})

}

