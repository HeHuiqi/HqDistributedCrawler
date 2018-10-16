package worker

import (
	"HqDistributedCrawler/engine"
	"HqDistributedCrawler/distribute/config"
	"HqDistributedCrawler/zhengai/parser"
	"errors"
	"fmt"
	"log"
)

type SerializeParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url string
	Parser SerializeParser
}

type ParserResult struct {

	Items []engine.Item
	Requests []Request
}

func SerializeRequst(r engine.Request) Request  {
	name,args := r.Parser.Serialize()
	return Request{
		Url:r.Url,
		Parser:SerializeParser{
			Name:name,
			Args:args,
		},
	}
}
func SerializeParserResult(r engine.ParserResult) ParserResult {

	result := ParserResult{
		Items:r.Items,
	}
	for _,req := range r.Requests {
		result.Requests = append(result.Requests,SerializeRequst(req))
	}
	return result
}

func DeserializeRequst(r Request) (engine.Request,error) {

	parser,err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{},err
	}
	return engine.Request{
		Url:r.Url,
		Parser:parser,
	},nil
}

func DeserializeParserResult(r ParserResult) engine.ParserResult {

	result := engine.ParserResult{
		Items:r.Items,
	}
	for _,req := range r.Requests {
		dreq,err := DeserializeRequst(req)
		if err != nil {
			log.Printf("error desrialializing request: %v",err)
			continue
		}
		result.Requests = append(result.Requests,dreq)

	}
	return result
}
func deserializeParser(p SerializeParser) (engine.Parser,error)  {
	switch p.Name {
	case config.ParserCityList:
		return engine.NewFuncParser(parser.ParserCityList,
			config.ParserCityList),nil

	case config.ParserCity:
		return engine.NewFuncParser(parser.ParserCity,
			config.ParserCity),nil

	case config.NilParser:
		return engine.NilParser{},nil

	case config.ParserProfile:
		if userName,ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName),nil

		}else {
			return nil,fmt.Errorf("Invalid arg: %v",p.Args)
		}


	default:
		fmt.Println("unknown parser name==",p.Name)
		return nil,errors.New("unknown parser name")
	}
}
