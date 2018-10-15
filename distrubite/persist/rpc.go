package persist

import (
	"HqDistributedCrawler/engine"
	"github.com/olivere/elastic"
	"HqDistributedCrawler/persist"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item,result *string) error {

	err := persist.Save(s.Client,s.Index,item)
	log.Printf("clientItemServcie %v saved.",item)
	if err == nil {
		*result = "ok"
	}else {
		log.Printf("clientItemServcie error saving  saved %v;item; eror:%v",item,err)

	}
	return err
}