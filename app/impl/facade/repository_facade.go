package facade

import (
	"errors"
	"gitcrawler/app/impl/contract"
	"gitcrawler/app/impl/entity"
	"gitcrawler/app/impl/enum"
	"gitcrawler/app/impl/service"
	"gitcrawler/app/impl/strategy"
	"os"
	"path/filepath"
	"strings"
)

type RepositoryFacade struct {
	cloneService           contract.CloneServiceContract
	crawlerService         contract.CrawlerServiceContract
	resumeGeneratorService contract.ResumeGeneratorServiceContract
}

func NewRepositoryFacade() *RepositoryFacade {
	return &RepositoryFacade{service.NewCloneService(), service.NewCrawlerService(), service.NewResumeGenerateService()}
}

func (c *RepositoryFacade) GetRepositoryFiles(url string, extensions []string, dirs []string, option enum.ConversionOption) (err error) {
	err = c.isUrlValid(url)
	if err != nil {
		return err
	}
	data, err := c.createAndCrawl(url, extensions, dirs)

	converter, err := c.converterStrategy(option)

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
	err = c.isUrlValid(url)
	if err != nil {
		return err
	}

	data, err := c.createAndCrawl(url, extensions, dirs)
	if err != nil {
		return err
	}
	if data == nil {
		return errors.New("repository business data is empty")
	}
	_, err = c.resumeGeneratorService.GenerateBusinessResume(data.String())
	if err != nil {
		return err
	}
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

func (c *RepositoryFacade) converterStrategy(option enum.ConversionOption) (converter strategy.DataConverter, err error) {
	switch option {
	case "csv":
		return strategy.NewConverterCsv(), nil
	default:
		return nil, errors.New("unknown conversion option")
	}
}

func (c *RepositoryFacade) isUrlValid(url string) error {
	if !strings.Contains(url, ".git") {
		return errors.New("repository url must contain .git")
	}
	return nil
}
