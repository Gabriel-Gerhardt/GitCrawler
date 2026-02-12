package facade

import (
	"gitcrawler/app/impl/service"
	"net/http"
)

type RepositoryFacade struct {
	cloneService   *service.CloneService
	crawlerService *service.CrawlerService
}

func NewRepositoryFacade() *RepositoryFacade {
	return &RepositoryFacade{service.NewCloneService(), service.NewCrawlerService()}
}

func (c *RepositoryFacade) GetAllRepositoryFiles(url string) *http.Response {
	return c.cloneService.CloneRepository(url)
}
func (c *RepositoryFacade) ReturnUrl(url string) string {
	return url
}
