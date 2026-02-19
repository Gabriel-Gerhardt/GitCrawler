package service

import (
	"fmt"
	"net/http"
	"os"
)

const openRouterUrl = "https://openrouter.ai/api/v1/chat/completions"

type ResumeGenerateService struct {
	key string
}

func NewResumeGenerateService() *ResumeGenerateService {
	return &ResumeGenerateService{key: os.Getenv("API_KEY")}
}

func (s *ResumeGenerateService) GenerateBusinessResume(data string) {
	fullPrompt := os.Getenv("AI_RESUME_PROMPT") + data
	fmt.Println(fullPrompt)
	http.Post(openRouterUrl, "", nil)
}
