package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Veri5ied/valyu-go/valyu"
	"github.com/Veri5ied/valyu-go/valyu/answer"
	"github.com/Veri5ied/valyu-go/valyu/common"
	"github.com/Veri5ied/valyu-go/valyu/search"
)

type SearchRequest struct {
	Query      string `json:"query"`
	MaxResults int    `json:"max_results,omitempty"`
}

type SummarizeRequest struct {
	Topic string `json:"topic"`
}

type PaperRequest struct {
	URL string `json:"url"`
}

type SearchResponse struct {
	Query   string         `json:"query"`
	Results []SearchResult `json:"results"`
	Count   int            `json:"count"`
}

type SearchResult struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	Abstract string `json:"abstract"`
	Source   string `json:"source"`
}

type SummarizeResponse struct {
	Topic   string      `json:"topic"`
	Summary interface{} `json:"summary"`
	Sources interface{} `json:"sources"`
}

type PaperResponse struct {
	URL      string      `json:"url"`
	Analysis interface{} `json:"analysis"`
	Sources  interface{} `json:"sources"`
}

func main() {
	defaultApiKey := os.Getenv("VALYU_API_KEY")

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		apiKey := getAPIKey(r, defaultApiKey)
		if apiKey == "" {
			http.Error(w, "API Key required (use X-API-Key header or set VALYU_API_KEY env var)", http.StatusUnauthorized)
			return
		}

		var req SearchRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.Query == "" {
			http.Error(w, "Query is required", http.StatusBadRequest)
			return
		}

		client, err := valyu.New(apiKey)
		if err != nil {
			log.Printf("failed to create valyu client: %v", err)
			http.Error(w, "Failed to initialize SDK", http.StatusInternalServerError)
			return
		}

		log.Printf("Searching academic papers: %s", req.Query)

		maxResults := req.MaxResults
		if maxResults == 0 {
			maxResults = 10
		}

		ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
		defer cancel()

		searchResp, err := client.Search.Search(ctx, req.Query, &search.Options{
			SearchType:      common.SearchTypeProprietary,
			MaxNumResults:   maxResults,
			IncludedSources: []string{"valyu/valyu-arxiv"},
		})
		if err != nil {
			log.Printf("SDK error: %v", err)
			http.Error(w, "Failed to search papers", http.StatusInternalServerError)
			return
		}

		if !searchResp.Success {
			log.Printf("API error: %s", searchResp.Error)
			http.Error(w, fmt.Sprintf("API error: %s", searchResp.Error), http.StatusBadGateway)
			return
		}

		results := make([]SearchResult, 0, len(searchResp.Results))
		for _, r := range searchResp.Results {
			results = append(results, SearchResult{
				Title:    r.Title,
				URL:      r.URL,
				Abstract: r.Description,
				Source:   r.Source,
			})
		}

		resp := SearchResponse{
			Query:   req.Query,
			Results: results,
			Count:   len(results),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/summarize", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		apiKey := getAPIKey(r, defaultApiKey)
		if apiKey == "" {
			http.Error(w, "API Key required (use X-API-Key header or set VALYU_API_KEY env var)", http.StatusUnauthorized)
			return
		}

		var req SummarizeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.Topic == "" {
			http.Error(w, "Topic is required", http.StatusBadRequest)
			return
		}

		client, err := valyu.New(apiKey)
		if err != nil {
			log.Printf("failed to create valyu client: %v", err)
			http.Error(w, "Failed to initialize SDK", http.StatusInternalServerError)
			return
		}

		log.Printf("Summarizing research topic: %s", req.Topic)

		prompt := fmt.Sprintf("Provide a comprehensive academic summary of recent research on '%s'. Include key findings, major contributors, current debates, and future research directions. Focus on peer-reviewed sources.", req.Topic)

		ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
		defer cancel()

		ans, err := client.Answer.Answer(ctx, prompt, &answer.Options{
			SearchType: common.SearchTypeProprietary,
		})
		if err != nil {
			log.Printf("SDK error: %v", err)
			http.Error(w, "Failed to generate summary", http.StatusInternalServerError)
			return
		}

		if !ans.Success {
			log.Printf("API error: %s", ans.Error)
			http.Error(w, fmt.Sprintf("API error: %s", ans.Error), http.StatusBadGateway)
			return
		}

		resp := SummarizeResponse{
			Topic:   req.Topic,
			Summary: ans.Contents,
			Sources: ans.SearchResults,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/papers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		apiKey := getAPIKey(r, defaultApiKey)
		if apiKey == "" {
			http.Error(w, "API Key required (use X-API-Key header or set VALYU_API_KEY env var)", http.StatusUnauthorized)
			return
		}

		var req PaperRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.URL == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		client, err := valyu.New(apiKey)
		if err != nil {
			log.Printf("failed to create valyu client: %v", err)
			http.Error(w, "Failed to initialize SDK", http.StatusInternalServerError)
			return
		}

		log.Printf("Analyzing paper: %s", req.URL)

		prompt := fmt.Sprintf("Analyze this research paper: %s. Provide: 1) Main hypothesis and research question, 2) Methodology used, 3) Key findings and contributions, 4) Limitations and critiques, 5) Implications for future research.", req.URL)

		ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
		defer cancel()

		ans, err := client.Answer.Answer(ctx, prompt, &answer.Options{
			SearchType: common.SearchTypeAll,
		})
		if err != nil {
			log.Printf("SDK error: %v", err)
			http.Error(w, "Failed to analyze paper", http.StatusInternalServerError)
			return
		}

		if !ans.Success {
			log.Printf("API error: %s", ans.Error)
			http.Error(w, fmt.Sprintf("API error: %s", ans.Error), http.StatusBadGateway)
			return
		}

		resp := PaperResponse{
			URL:      req.URL,
			Analysis: ans.Contents,
			Sources:  ans.SearchResults,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Academic Research API server starting on :%s", port)
	log.Printf("Available endpoints:")
	log.Printf("  POST /search - Search academic papers")
	log.Printf("  POST /summarize - Summarize research topics")
	log.Printf("  POST /papers - Analyze specific papers")
	log.Printf("  GET  /health - Health check")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func getAPIKey(r *http.Request, defaultKey string) string {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		apiKey = defaultKey
	}
	return apiKey
}
