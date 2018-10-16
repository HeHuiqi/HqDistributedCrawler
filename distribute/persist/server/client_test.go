package main

import (
	"testing"
	"HqDistributedCrawler/distribute/rpcsupport"
	"HqDistributedCrawler/model"
	"HqDistributedCrawler/engine"
	"fmt"
	"time"
	"HqDistributedCrawler/distribute/config"
)

func TestItemSaver(t *testing.T)  {
	const host  = ":1234"
	go serveRpc(host,"test")
	time.Sleep(time.Second)

	client,err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url:"http://album.zhenai.com/u/108906739",
		Type:"zhenai",
		Id:"108906739",
		Payload:model.Profile{
			Age:34,
			Height:162,
			Weight:162,
			Income:"3001-5000元",
			Gender:"女",
			Name:"安静的雪",
			Xinzuo:"牡羊座",
			Occupation:"人事/行政",
			Marriage:"离异",
			House:"已购房",
			Hukou:"山东菏泽",
			Education:"大学本科",
			Car:"未购车",
		},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc,item,&result)
	if err != nil || result != "ok"{
		t.Errorf("result:%v; err: %v",result,err)
	}
	fmt.Println(result)
}
