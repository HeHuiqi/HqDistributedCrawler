package client

import (
	"github.com/olivere/elastic"
	"log"
	"reflect"
	"HqDistributedCrawler/engine"
	"errors"
	"encoding/json"
	"context"
	"HqDistributedCrawler/distribute/rpcsupport"
	"HqDistributedCrawler/distribute/config"
)

func ItemSaver(host string) (chan engine.Item,error){

	client,err := rpcsupport.NewClient(host)
	if err != nil {
		return nil,err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("client/isv Item Saver Got Item #%d: %v",itemCount,item)
			itemCount++
			// call RPC to save Item
			result := ""
			err = client.Call(config.ItemSaverRpc,item,&result)
			if err != nil {
				log.Printf("client/isv Item Saver:  saving item %v; error:%v",item,err)
			}
		}
	}()
	return out,nil
}
func Save(client *elastic.Client,index string,item engine.Item) error  {

	//这里不存储城市列表名字和某个城市下用户名
	if reflect.TypeOf(item.Payload).Name() == "string" {
		return nil
	}

	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	//测试先不存储
	return nil

	itemJsonStr,_ := json.Marshal(item)
	//fmt.Println("Item",string(itemJsonStr))

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyString(string(itemJsonStr))

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_,err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}