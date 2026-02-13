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
	err = c.crawl(path, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c *CrawlerService) crawl(dir string, repositoryData *entity.RepositoryData) (err error) {
	readDir, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range readDir {
		if file.IsDir() {
			err := c.crawl(dir+"/"+file.Name(), repositoryData)
			if err != nil {
				return err
			}
		} else {
			repositoryFile := &entity.RepositoryFile{}
			fileBinary, err := os.Open(dir + "/" + file.Name())
			if err != nil {
				return err
			}
			data, err := io.ReadAll(fileBinary)
			_ = fileBinary.Close()
			if err != nil {
				return err
			}
			repositoryFile.Data = string(data)
			repositoryFile.Path = dir + "/" + file.Name()

			repositoryData.Files = append(repositoryData.Files, *repositoryFile)
		}
	}

	return nil
}
