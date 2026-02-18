package strategy

import "gitcrawler/app/impl/entity"

type DataConverter interface {
	Convert(data *entity.RepositoryData) (err error)
}
