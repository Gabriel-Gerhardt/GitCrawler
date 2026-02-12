package rest

import (
	"gitcrawler/app/impl/facade"
)

type CrawlerController struct {
	repositoryFacade *facade.RepositoryFacade
}

func NewCrawlerController() *CrawlerController {
	return &CrawlerController{facade.NewRepositoryFacade()}
}

func (c *CrawlerController) GetAllRepositoryFiles(url string) {
	if url == "" {
		return
	}
	c.repositoryFacade.GetAllRepositoryFiles(url)
}
