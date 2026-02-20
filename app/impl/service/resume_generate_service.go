package service

import (
	"bytes"
	"fmt"
	"io"
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

func (s *ResumeGenerateService) GenerateBusinessResume(data string) (text []byte, err error) {

	request, err := s.buildRequest(data)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(string(respBody))
	
	return respBody, nil
}

func (s *ResumeGenerateService) buildRequest(data string) (*http.Request, error) {
	fullPrompt := os.Getenv("AI_RESUME_PROMPT") + " build a resume of the data of this repo. data:" + data

	body := []byte(`{
		"model": "openai/gpt-4o-mini",
		"messages": [{"role": "user", "content": "` + escapeJSON(fullPrompt) + `"}]
	}`)

	request, err := http.NewRequest("POST", openRouterUrl, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	request.Header.Set("Authorization: Bearer", s.key)
	return request, nil
}

func escapeJSON(str string) string {
	b := []byte(str)
	b = bytes.ReplaceAll(b, []byte(`\`), []byte(`\\`))
	b = bytes.ReplaceAll(b, []byte(`"`), []byte(`\"`))
	b = bytes.ReplaceAll(b, []byte("\n"), []byte(`\n`))
	return string(b)
}
