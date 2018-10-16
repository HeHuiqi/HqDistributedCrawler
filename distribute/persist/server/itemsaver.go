package main

import (
	"HqDistributedCrawler/distribute/rpcsupport"
	"HqDistributedCrawler/distribute/persist"
	"github.com/olivere/elastic"
	"fmt"
	"HqDistributedCrawler/distribute/config"
)

//存储server
func main() {

	serveRpc(fmt.Sprintf(":%d",config.ItemSaverPort),config.ElasticIndex)
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
