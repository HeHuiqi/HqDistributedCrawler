package main

import (
	"HqDistributedCrawler/distribute/rpcsupport"
	"HqDistributedCrawler/distribute/persist"
	"github.com/olivere/elastic"
	"fmt"
	"HqDistributedCrawler/distribute/config"
	"flag"
)

var port  = flag.Int("port",0,
	"the port for me to listen on")
//开启saver
//go run itemsaver.go --port=1234
func main() {

	flag.Parse()
	if *port == 0{
		fmt.Println("must sepcify a port")
		return
	}

	serveRpc(fmt.Sprintf(":%d",*port),config.ElasticIndex)
}

func serveRpc(host,index string) error {
	client,err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host,&persist.ItemSaverService{
		Client:client,
		Index:index,
	})
}
