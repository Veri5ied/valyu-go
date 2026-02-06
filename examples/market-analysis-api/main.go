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
)

type AnalysisRequest struct {
	Query string `json:"query"`
}

type AnalysisResponse struct {
	Company  string      `json:"company"`
	Analysis interface{} `json:"analysis"`
	Sources  interface{} `json:"sources"`
}

func main() {
	defaultApiKey := os.Getenv("VALYU_API_KEY")

	http.HandleFunc("/analyze", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			apiKey = defaultApiKey
		}

		if apiKey == "" {
			http.Error(w, "API Key required (use X-API-Key header or set VALYU_API_KEY env var)", http.StatusUnauthorized)
			return
		}

		var req AnalysisRequest
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

		log.Printf("Analyzing: %s", req.Query)

		// Prompt for market analysis
		prompt := fmt.Sprintf("Perform a comprehensive market analysis for %s. Include current market position, recent financial performance, future outlook, and key risks/opportunities.", req.Query)

		ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
		defer cancel()

		ans, err := client.Answer.Answer(ctx, prompt, nil)
		if err != nil {
			log.Printf("SDK error: %v", err)
			http.Error(w, "Failed to analyze market", http.StatusInternalServerError)
			return
		}

		if !ans.Success {
			log.Printf("API error: %s", ans.Error)
			http.Error(w, fmt.Sprintf("API error: %s", ans.Error), http.StatusBadGateway)
			return
		}

		resp := AnalysisResponse{
			Company:  req.Query,
			Analysis: ans.Contents,
			Sources:  ans.SearchResults,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Market Analysis API server starting on :%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
