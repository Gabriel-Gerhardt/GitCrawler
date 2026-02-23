package strategy

import (
	"gitcrawler/app/impl/core/entity"
)

type DataConverter interface {
	Convert(data *entity.RepositoryData) (err error)
}
