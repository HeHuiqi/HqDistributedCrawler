package worker

import "HqDistributedCrawler/engine"

type CrawlService struct {

}

func (CrawlService) Process(req Request,
	result *ParserResult) error {

	engineReq,err := DeserializeRequst(req)
	if err != nil {
		return  err
	}
	engineResult,err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeParserResult(engineResult)

	return nil
}
