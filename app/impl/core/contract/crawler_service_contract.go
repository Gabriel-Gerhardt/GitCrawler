package contract

import (
	"gitcrawler/app/impl/core/entity"
)

type CrawlerServiceContract interface {
	CrawlRepository(path string, repoName string, validExtensions []string, validDirs []string) (data *entity.RepositoryData, err error)
}
