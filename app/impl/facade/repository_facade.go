package facade

import (
	"fmt"
	"gitcrawler/app/impl/contract"
	"gitcrawler/app/impl/service"
	"os"
)

type RepositoryFacade struct {
	cloneService   contract.CloneServiceContract
	crawlerService contract.CrawlerServiceContract
}

func NewRepositoryFacade() *RepositoryFacade {
	return &RepositoryFacade{service.NewCloneService(), service.NewCrawlerService()}
}

func (c *RepositoryFacade) GetAllRepositoryFiles(url string) (err error) {
	path, err := c.cloneService.CloneRepository(url)

	if path != "" {
		defer os.RemoveAll(path)
	}

	if err != nil {
		return err
	}

	data, err := c.crawlerService.CrawlRepository(path)
	if err != nil {
		return err
	}
	fmt.Println(data)

	return nil
}
