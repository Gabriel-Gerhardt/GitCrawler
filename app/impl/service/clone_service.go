package service

import (
	"net/http"
	"os"
	"os/exec"
)

const BASE_PATH = "../../resources"

type CloneService struct {
}

func NewCloneService() *CloneService {
	return &CloneService{}
}

func (c *CloneService) CloneRepository(repositoryUrl string) (string, error, http.Response) {
	path, err := c.createRepositoryDirectory()
	if err != nil {
		return "", err, http.Response{StatusCode: http.StatusInternalServerError, Body: nil}
	}
	cmd := exec.Command("git", "clone", repositoryUrl, ".")
	cmd.Dir = path
	err = cmd.Run()

	if err != nil {
		return "", err, http.Response{StatusCode: http.StatusBadRequest, Body: nil}
	}
	return path, nil, http.Response{StatusCode: http.StatusOK, Body: nil}
}

func (c *CloneService) createRepositoryDirectory() (string, error) {
	path, err := os.MkdirTemp(BASE_PATH, "temp")
	if err != nil {
		return "", err
	}
	return path, nil
}
