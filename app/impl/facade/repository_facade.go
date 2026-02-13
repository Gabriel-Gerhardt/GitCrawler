package facade

import (
	"fmt"
	"gitcrawler/app/impl/service"
	"os"
)

type RepositoryFacade struct {
	cloneService   *service.CloneService
	crawlerService *service.CrawlerService
}

func NewRepositoryFacade() *RepositoryFacade {
	return &RepositoryFacade{service.NewCloneService(), service.NewCrawlerService()}
}

func (c *RepositoryFacade) GetAllRepositoryFiles(url string) (err error) {
	path, err := c.cloneService.CloneRepository(url)
	if err != nil {
		return err
	}
	data, err := c.crawlerService.CrawlRepository(path)
	if err != nil {
		return err
	}
	fmt.Println(data)

	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			return
		}
	}(path)

	return nil
}
