package service

import (
	"errors"
	"fmt"
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
		return path, err
	}
	cmd := exec.Command("git", "clone", repositoryUrl, ".")
	cmd.Dir = path
	err = cmd.Run()
	if err != nil {
		return path, errors.New(fmt.Sprintf("Repository not found, project may be private: %s", err.Error()))
	}
	return path, nil
}

func (c *CloneService) createRepositoryDirectory() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path, err = os.MkdirTemp(path, "temp")
	return path, nil
}
