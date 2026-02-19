package register

import (
	"gitcrawler/app/impl/rest"
	"net/http"
)

func GetHandlers() {
	crawlerController := rest.NewCrawlerController()
	http.Handle("/getAllRepoData", http.HandlerFunc(crawlerController.GetAllRepositoryFiles))
	http.Handle("/getBusinessRepoResume", http.HandlerFunc(crawlerController.GetBusinessRepoResume))
}
