package register

import (
	"gitcrawler/app/impl/rest"
	"net/http"
)

func GetHandlers() {
	crawlerController := rest.NewCrawlerController()
	http.Handle("/getAll", http.HandlerFunc(crawlerController.GetAllRepositoryFiles))
}
