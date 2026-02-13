package service

import (
	"gitcrawler/app/impl/entity"
	"io"
	"os"
	"path/filepath"
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
		path := filepath.Join(dir, file.Name())
		if file.IsDir() {
			err = c.crawl(path, repositoryData)
			if err != nil {
				return err
			}
		} else {
			fileData, openFileErr := c.openFile(path)
			if openFileErr != nil {
				return openFileErr
			}
			c.appendFileData(repositoryData, path, fileData)

		}
	}
	return nil
}

func (c *CrawlerService) openFile(path string) (data string, err error) {
	fileBinary, err := os.Open(path)

	if err != nil {
		return "", err
	}
	defer fileBinary.Close()
	binaryData, err := io.ReadAll(fileBinary)

	if err != nil {
		return "", err
	}
	data = string(binaryData)
	return data, nil
}

func (c *CrawlerService) appendFileData(repositoryData *entity.RepositoryData, path string, data string) {
	repositoryFile := &entity.RepositoryFile{}
	repositoryFile.Data = data
	repositoryFile.Path = path
	repositoryData.Files = append(repositoryData.Files, *repositoryFile)
}
