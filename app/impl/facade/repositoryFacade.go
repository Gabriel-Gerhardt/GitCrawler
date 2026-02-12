package facade

import (
	"gitcrawler/app/impl/service"
)

type RepositoryFacade struct {
	cloneService   *service.CloneService
	crawlerService *service.CrawlerService
}

func NewRepositoryFacade() *RepositoryFacade {
	return &RepositoryFacade{service.NewCloneService(), service.NewCrawlerService()}
}

func (c *RepositoryFacade) GetAllRepositoryFiles(url string) {
	c.cloneService.CloneRepository(url)
}
