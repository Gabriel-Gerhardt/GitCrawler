package service

import (
	"fmt"
	"gitcrawler/app/impl/entity"
)

type CrawlerService struct {
}

func NewCrawlerService() *CrawlerService {
	return &CrawlerService{}
}

func (c *CrawlerService) CrawlRepository(path string) (data *entity.RepositoryData, err error) {
	fmt.Println(data)

	return nil, nil
}
func (c *CrawlerService) crawl(url string) (string, error) {
	return "", nil
}
