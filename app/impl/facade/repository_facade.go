package facade

import (
	"fmt"
	"gitcrawler/app/impl/contract"
	"gitcrawler/app/impl/entity"
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

func (c *RepositoryFacade) GetAllRepositoryFiles(url string, extensions []string, dirs []string) (err error) {
	data, err := c.createAndCrawl(url, extensions, dirs)

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

func (c *RepositoryFacade) GenerateBusinessResume(url string) (err error) {
	extensions := []string{
		".java",
		".kt",
		".kts",
		".groovy",
		".go",
		".py",
		".ts",
		".js",
		".cs",
		".rb",
		".php",
		".rs",
		".cpp",
		".cc",
		".cxx",
		".c",
		".h",
		".hpp",
		".scala",
		".ex",
		".exs",
		".dart",
		".swift",
	}

	dirs := []string{
		"rest",
		"api",
		"controller",
		"handler",
		"service",
		"usecase",
		"application",
		"domain",
		"model",
		"entity",
		"aggregate",
		"facade",
		"client",
		"gateway",
		"integration",
		"repository",
		"event",
	}
	data, err := c.createAndCrawl(url, extensions, dirs)
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}

func (c *RepositoryFacade) createAndCrawl(url string, extensions []string, dirs []string) (data *entity.RepositoryData, err error) {
	path, err := c.cloneService.CloneRepository(url)
	if path != "" {
		defer os.RemoveAll(path)
	}

	if err != nil {
		return nil, err
	}
	repoName := strings.TrimSuffix(filepath.Base(url), ".git")

	data, err = c.crawlerService.CrawlRepository(path, repoName, extensions, dirs)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *RepositoryFacade) converterStrategy() (converter strategy.DataConverter, err error) {
	return strategy.NewConverterCsv(), nil
}
