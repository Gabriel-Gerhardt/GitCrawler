package rest

import (
	"encoding/json"
	"gitcrawler/app/impl/adapters/dto/request"
	"gitcrawler/app/impl/adapters/dto/response"
	"gitcrawler/app/impl/adapters/facade"
	"net/http"
)

type CrawlerController struct {
	repositoryFacade *facade.RepositoryFacade
}

func NewCrawlerController() *CrawlerController {
	return &CrawlerController{facade.NewRepositoryFacade()}
}

func (c *CrawlerController) GetRepositoryFiles(w http.ResponseWriter, r *http.Request) {
	var req request.RepositoryFilesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if req.Url == "" {
		http.Error(w, "Url must contain something", http.StatusBadRequest)
		return
	}
	err := c.repositoryFacade.GetRepositoryFiles(req.Url, req.Extensions, req.Dirs, req.Option)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("The csv is ready"))

}

func (c *CrawlerController) GetBusinessRepoResume(w http.ResponseWriter, r *http.Request) {
	var resp response.ResumeResponse
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Url must contain something", http.StatusBadRequest)
		return
	}

	aiResponse, err := c.repositoryFacade.GenerateBusinessResume(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.AiResponse = aiResponse

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp.AiResponse))
}
