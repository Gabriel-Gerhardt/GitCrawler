package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

const openRouterUrl = "https://openrouter.ai/api/v1/chat/completions"

type ResumeGenerateService struct {
	key string
}

func NewResumeGenerateService() *ResumeGenerateService {
	return &ResumeGenerateService{key: os.Getenv("API_KEY")}
}

func (s *ResumeGenerateService) GenerateBusinessResume(data string) (text string, err error) {
	if s.key == "" {
		return "", errors.New("API_KEY not set")
	}
	request, err := s.buildRequest(data)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("could not read response body of AI requisition")
	}
	defer resp.Body.Close()

	aiResponse, err := s.getResponse(respBody)
	if err != nil {
		return "", err
	}
	fmt.Println(aiResponse)
	return aiResponse, nil
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
		return nil, errors.New("unable to mount request")
	}
	request.Header.Set("Authorization", "Bearer "+s.key)

	return request, nil
}

func escapeJSON(str string) string {
	b := []byte(str)
	b = bytes.ReplaceAll(b, []byte(`\`), []byte(`\\`))
	b = bytes.ReplaceAll(b, []byte(`"`), []byte(`\"`))
	b = bytes.ReplaceAll(b, []byte("\n"), []byte(`\n`))
	return string(b)
}

func (s *ResumeGenerateService) getResponse(respBody []byte) (choice string, err error) {
	var aiResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err = json.Unmarshal(respBody, &aiResp); err != nil {
		return "", errors.New("could not parse AI response JSON")
	}

	if len(aiResp.Choices) == 0 {
		return "", errors.New("AI has no response")
	}

	return aiResp.Choices[0].Message.Content, nil
}
