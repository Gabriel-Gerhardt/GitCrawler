package service

import (
	"gitcrawler/app/impl/entity"
	"io"
	"os"
)

type CrawlerService struct {
}

func NewCrawlerService() *CrawlerService {
	return &CrawlerService{}
}

func (c *CrawlerService) CrawlRepository(path string) (data *entity.RepositoryData, err error) {
	data = &entity.RepositoryData{}
	_, err = c.crawl(path, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (c *CrawlerService) crawl(dir string, repositoryData *entity.RepositoryData) (path string, err error) {
	readDir, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, file := range readDir {
		if file.IsDir() {
			_, err := c.crawl(dir+"/"+file.Name(), repositoryData)
			if err != nil {
				return "", err
			}
		} else {
			repositoryFile := &entity.RepositoryFile{}
			fileBinary, err := os.OpenFile(dir+"/"+file.Name(), os.O_RDONLY, os.ModePerm)
			if err != nil {
				return "", err
			}
			data, _ := io.ReadAll(fileBinary)

			repositoryFile.Data = string(data)
			repositoryFile.Path = dir + "/" + file.Name()

			repositoryData.Files = append(repositoryData.Files, *repositoryFile)
		}
	}

	return "", nil
}
