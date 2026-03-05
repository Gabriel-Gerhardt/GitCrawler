package register

import (
	"gitcrawler/app/impl/external/rest"
	"net/http"
)

func GetHandlers(crawlerController *rest.CrawlerController) {
	http.Handle("/getRepoData", http.HandlerFunc(crawlerController.GetRepositoryFiles))
	http.Handle("/getBusinessRepoResume", http.HandlerFunc(crawlerController.GetBusinessRepoResume))
}
