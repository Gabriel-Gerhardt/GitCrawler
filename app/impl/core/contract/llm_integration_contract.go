package contract

type LlmIntegrationContract interface {
	ReturnAIResponse(prompt string) (aiResponse string, err error)
}
