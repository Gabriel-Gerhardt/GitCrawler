package request

import (
	"gitcrawler/app/impl/core/enum"
)

type RepositoryFilesRequest struct {
	Url        string                `json:"url"`
	Dirs       []string              `json:"dirs"`
	Extensions []string              `json:"extensions"`
	Option     enum.ConversionOption `json:"option"`
}
