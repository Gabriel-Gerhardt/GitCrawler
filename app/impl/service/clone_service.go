package service

import (
	"os"
	"os/exec"
)

type CloneService struct {
}

func NewCloneService() *CloneService {
	return &CloneService{}
}

func (c *CloneService) CloneRepository(repositoryUrl string) (string, error) {
	path, err := c.createRepositoryDirectory()
	if err != nil {
		return "", err
	}
	cmd := exec.Command("git", "clone", repositoryUrl, ".")
	cmd.Dir = path
	err = cmd.Run()

	if err != nil {
		return "", err
	}
	return path, nil
}

func (c *CloneService) createRepositoryDirectory() (string, error) {
	path, err := os.Getwd()
	path, err = os.MkdirTemp(path, "temp")

	if err != nil {
		return "", err
	}
	return path, nil
}
