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
			err = c.crawl(dir+"/"+file.Name(), repositoryData)
			if err != nil {
				return err
			}
		} else {
			fileData, path, openFileErr := c.openFile(dir, file)
			if openFileErr != nil {
				return err
			}
			c.appendFileData(repositoryData, path, fileData)

		}
	}
	return nil
}

func (c *CrawlerService) openFile(dir string, file os.DirEntry) (data string, path string, err error) {
	path = dir + "/" + file.Name()
	fileBinary, err := os.Open(path)

	if err != nil {
		return "", path, err
	}
	defer fileBinary.Close()

	binaryData, err := io.ReadAll(fileBinary)
	if err != nil {
		return "", path, err
	}
	data = string(binaryData)
	return data, path, nil
}

func (c *CrawlerService) appendFileData(repositoryData *entity.RepositoryData, path string, data string) {
	repositoryFile := &entity.RepositoryFile{}
	repositoryFile.Data = data
	repositoryFile.Path = path
	repositoryData.Files = append(repositoryData.Files, *repositoryFile)
}
