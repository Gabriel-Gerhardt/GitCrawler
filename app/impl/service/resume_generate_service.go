package service

import (
	"fmt"
	"os"
)

type ResumeGenerateService struct {
	key string
}

func NewResumeGenerateService() *ResumeGenerateService {
	return &ResumeGenerateService{key: os.Getenv("API_KEY")}
}

func (s *ResumeGenerateService) GenerateBusinessResume(data string) {
	fullPrompt := os.Getenv("AI_RESUME_PROMPT") + data
	fmt.Println(fullPrompt)

}
