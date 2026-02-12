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

func (c *CrawlerService) crawlRepository() {
	data := entity.RepositoryData{}
	fmt.Println(data)
}
