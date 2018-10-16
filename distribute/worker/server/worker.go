package main

import (
	"HqDistributedCrawler/distribute/rpcsupport"
	"fmt"
	"HqDistributedCrawler/distribute/config"
	"HqDistributedCrawler/distribute/worker"
)

func main() {

	err := rpcsupport.ServeRpc(fmt.Sprintf(":%d",config.WorkerPort0),
		worker.CrawlService{})
	if err != nil {
		panic(err)
	}
	
}
