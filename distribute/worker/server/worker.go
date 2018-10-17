package main

import (
	"HqDistributedCrawler/distribute/rpcsupport"
	"fmt"
	"HqDistributedCrawler/distribute/worker"
	"flag"
)

var port  = flag.Int("port",0,
	"the port for me to listen on")
//可以开多个worker
/*
go run worker.go --port=9000
go run  worker.go --port=9001
*/
func main() {

	flag.Parse()
	if *port == 0{
		fmt.Println("must sepcify a port")
		return
	}
	err := rpcsupport.ServeRpc(fmt.Sprintf(":%d",*port),
		worker.CrawlService{})
	if err != nil {
		panic(err)
	}
	
}
