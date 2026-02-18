package rest

import (
	"gitcrawler/app/impl/facade"
	"net/http"
)

type CrawlerController struct {
	repositoryFacade *facade.RepositoryFacade
}

func NewCrawlerController() *CrawlerController {
	return &CrawlerController{facade.NewRepositoryFacade()}
}

func (c *CrawlerController) GetAllRepositoryFiles(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Url must contain something", http.StatusBadRequest)
	}
	err := c.repositoryFacade.GetAllRepositoryFiles("https://github.com/Gabriel-Gerhardt/Webhook-Manager.git")
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("fuck you"))

}
