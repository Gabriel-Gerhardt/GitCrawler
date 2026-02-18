package contract

import "gitcrawler/app/impl/entity"

type CrawlerServiceContract interface {
	CrawlRepository(path string, repoName string) (data *entity.RepositoryData, err error)
}
