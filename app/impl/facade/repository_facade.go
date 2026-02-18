package facade

import (
	"gitcrawler/app/impl/contract"
	"gitcrawler/app/impl/service"
	"gitcrawler/app/impl/strategy"
	"os"
	"path/filepath"
	"strings"
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
	repoName := strings.TrimSuffix(filepath.Base(url), ".git")

	extensions := []string{
		".java",
		".go",
	}
	data, err := c.crawlerService.CrawlRepository(path, repoName, extensions)
	if err != nil {
		return err
	}
	converter, err := c.converterStrategy()

	if err != nil {
		return err
	}

	err = converter.Convert(data)

	if err != nil {
		return err
	}
	return nil
}

func (c *RepositoryFacade) converterStrategy() (converter strategy.DataConverter, err error) {
	return strategy.NewConverterCsv(), nil
}
