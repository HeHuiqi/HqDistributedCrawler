package config

const (
	// Parser names
	ParserCity = "ParserCity"
	ParserCityList = "ParserCityList"
	ParserProfile = "ParserProfile"
	NilParser = "NilParser"

	//Service ports
	//ItemSaverPort = 1234
	//WorkerPort0 = 9000

	// ElasticSearch
	ElasticIndex = "dating_profile"

	//RPC Endpoints
	ItemSaverRpc = "ItemSaverService.Save"

	CrawlServiceRpc = "CrawlService.Process"


	// Rate limiting
	Qps = 20
)
