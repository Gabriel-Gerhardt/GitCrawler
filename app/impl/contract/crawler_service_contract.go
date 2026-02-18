package contract

import "gitcrawler/app/impl/entity"

type CrawlerServiceContract interface {
	CrawlRepository(path string) (data *entity.RepositoryData, err error)
}
