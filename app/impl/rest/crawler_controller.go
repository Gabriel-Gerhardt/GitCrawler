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

func (c *CrawlerController) GetAllRepositoryFiles(url string) (err error, response http.Response) {
	if url == "" {
		return nil, http.Response{StatusCode: http.StatusNotFound, Body: nil}
	}
	return c.repositoryFacade.GetAllRepositoryFiles(url)
}
