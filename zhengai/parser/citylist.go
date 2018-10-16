package parser

import (
	"HqDistributedCrawler/engine"
	"regexp"
	"HqDistributedCrawler/distribute/config"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`


func ParserCityList(contents []byte,_ string) engine.ParserResult  {

	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(contents,-1)
	result := engine.ParserResult{}

	for _,m := range match{

		/*
		//这里就不保存城市链接解析的数据了
		url := string(m[1])
		name := string(m[2])
		cpinyin := url[len("http://www.zhenai.com/zhenghun/"):]
		item := engine.Item{
			Url: url,
			Type:"zhenai",
			Id:cpinyin,
			Payload:name,

		}
		result.Items = append(result.Items,item)
		*/
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			Parser:engine.NewFuncParser(ParserCity,config.ParserCity),
		})
		//fmt.Printf("City: %s, URL: %s",m[2],m[1])
		//fmt.Println()

	}
	return result
}

