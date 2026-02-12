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

func (c *RepositoryFacade) GetAllRepositoryFiles(url string) (err error, response http.Response) {
	_, err, response = c.cloneService.CloneRepository(url)
	if err != nil {
		return err, response
	}
	/*defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			return
		}
	}(path) */
	return nil, response
}
