package main

import (
	"net/http"
	"HqDistributedCrawler/web/controller"
	"log"
)

func main() {
	http.Handle("/search",controller.NewSearchResultHandler("./web/view/template.html"))
	log.Printf("SearchWeb:http://localhost:8888/search?q=1709891452")
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic(err)
	}
}
