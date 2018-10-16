package main

import (
	"HqDistributedCrawler/engine"
	"HqDistributedCrawler/scheduler"
	"HqDistributedCrawler/zhengai/parser"
	itemsaver "HqDistributedCrawler/distribute/persist/client"
	"fmt"
	"HqDistributedCrawler/distribute/config"
	worker "HqDistributedCrawler/distribute/worker/client"
)
//分布式版
func main() {

	itemChan,err := itemsaver.ItemSaver(fmt.Sprintf(":%d",config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	processor,err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan:itemChan,
		RequestProcessor:processor,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser:engine.NewFuncParser(parser.ParserCityList,
			config.ParserCityList),
	})
}
