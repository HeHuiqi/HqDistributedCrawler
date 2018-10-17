package main

import (
	"HqDistributedCrawler/engine"
	"HqDistributedCrawler/scheduler"
	"HqDistributedCrawler/zhengai/parser"
	itemsaver "HqDistributedCrawler/distribute/persist/client"
	"HqDistributedCrawler/distribute/config"
	worker "HqDistributedCrawler/distribute/worker/client"
	"net/rpc"
	"HqDistributedCrawler/distribute/rpcsupport"
	"log"
	"flag"
	"strings"
)
var (
	itemSaverHost = flag.String("itemsaver_host","",
		"itemsaver host")
	workerHosts = flag.String("worker_hosts","",
		"worker hosts (comma separated)")

)
/*
注意这里的host要和itemsaver的host以及所有worker的端口保持一致
 go run main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"
*/
//分布式版
func main() {

	flag.Parse()

	itemChan,err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
	pool := createClientPool(
		strings.Split(*workerHosts,","),
	)
	processor := worker.CreateProcessor(pool)

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

func createClientPool(hosts [] string) chan *rpc.Client{
	var clients []*rpc.Client
	for _,h := range hosts {
		client,err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients,client)
			log.Printf("connected to %s",h)

		}else {
			log.Printf("error connecting to %s: %v",h,err)
		}
	}

	out := make(chan  *rpc.Client)
	//轮流分发
	go func() {
		for {
			for _,client := range clients {
				out <- client
			}
		}

	}()
	return out
}