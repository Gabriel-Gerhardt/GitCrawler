package register

import (
	"gitcrawler/app/impl/rest"
	"net/http"
)

func GetHandlers() {
	crawlerController := rest.NewCrawlerController()
	http.Handle("/getRepoData", http.HandlerFunc(crawlerController.GetRepositoryFiles))
	http.Handle("/getBusinessRepoResume", http.HandlerFunc(crawlerController.GetBusinessRepoResume))
}
