package main

import (
	"HqDistributedCrawler/distrubite/rpcsupport"
	"HqDistributedCrawler/distrubite/persist"
	"github.com/olivere/elastic"
)

func main() {

	serveRpc(":1234","dating_profile")
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
