package service

import (
	"net/http"
	"strings"
)

const GITHUB_API_URL = "https://api.github.com"

type CloneService struct {
}

func NewCloneService() *CloneService {
	return &CloneService{}
}

func (c *CloneService) CloneRepository(repositoryUrl string) (response *http.Response) {

	request, err := http.NewRequest("POST", GITHUB_API_URL, strings.NewReader(repositoryUrl))
	if err != nil {
		panic(err)
	}
	client := http.Client{}
	response, err = client.Do(request)
	if err != nil {
		return nil
	}
	return response
}
