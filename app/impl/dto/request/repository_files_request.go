package request

import "gitcrawler/app/impl/enum"

type RepositoryFilesRequest struct {
	Url        string                `json:"url"`
	Dirs       []string              `json:"dirs"`
	Extensions []string              `json:"extensions"`
	Option     enum.ConversionOption `json:"option"`
}
