package service

import (
	"os"
	"os/exec"
)

const BASE_PATH = "../../resources"

type CloneService struct {
}

func NewCloneService() *CloneService {
	return &CloneService{}
}

func (c *CloneService) CloneRepository(repositoryUrl string) error {
	path, err := c.createRepositoryDirectory()
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "clone", repositoryUrl, ".")
	cmd.Dir = path
	err = cmd.Run()
	
	if err != nil {
		return err
	}
	return nil
}

func (c *CloneService) createRepositoryDirectory() (string, error) {
	path, err := os.MkdirTemp(BASE_PATH, "temp")
	if err != nil {
		return "", err
	}
	return path, nil
}
