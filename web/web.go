package main

import (
	"net/http"
	"HqDistributedCrawler/web/controller"
)

func main() {
	http.Handle("/search",controller.NewSearchResultHandler("./web/view/template.html"))
	http.ListenAndServe(":8888",nil)
}
