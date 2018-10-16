package client

import (
	"HqDistributedCrawler/engine"
	"HqDistributedCrawler/distribute/rpcsupport"
	"fmt"
	"HqDistributedCrawler/distribute/config"
	"HqDistributedCrawler/distribute/worker"
)

func CreateProcessor() (engine.Processor,error)  {
	client,err := rpcsupport.NewClient(fmt.Sprintf(":%d",config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParserResult,error)  {

		sReq:=worker.SerializeRequst(req)

		var sResult worker.ParserResult

		err := client.Call(config.CrawlServiceRpc,sReq,&sResult)

		if err != nil {
			return engine.ParserResult{},err
		}

		return  worker.DeserializeParserResult(sResult),nil
	},nil
}