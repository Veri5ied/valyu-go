package valyu

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type answerRequest struct {
	Query              string      `json:"query"`
	SearchType         string      `json:"search_type,omitempty"`
	StructuredOutput   interface{} `json:"structured_output,omitempty"`
	SystemInstructions string      `json:"system_instructions,omitempty"`
	DataMaxPrice       float64     `json:"data_max_price,omitempty"`
	CountryCode        string      `json:"country_code,omitempty"`
	IncludedSources    []string    `json:"included_sources,omitempty"`
	ExcludedSources    []string    `json:"excluded_sources,omitempty"`
	StartDate          string      `json:"start_date,omitempty"`
	EndDate            string      `json:"end_date,omitempty"`
	FastMode           bool        `json:"fast_mode,omitempty"`
}

func (c *Client) Answer(ctx context.Context, query string, opts *AnswerOptions) (*AnswerResponse, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return &AnswerResponse{
			Success: false,
			Error:   "Query is required and must be a non-empty string",
		}, nil
	}

	if opts == nil {
		opts = &AnswerOptions{}
	}

	searchType := opts.SearchType
	if searchType == "" {
		searchType = SearchTypeAll
	}

	validSearchTypes := map[SearchType]bool{
		SearchTypeAll:         true,
		SearchTypeWeb:         true,
		SearchTypeProprietary: true,
		SearchTypeNews:        true,
	}
	if !validSearchTypes[searchType] {
		return &AnswerResponse{
			Success: false,
			Error:   "Invalid searchType provided. Must be one of: all, web, proprietary, news",
		}, nil
	}

	if opts.SystemInstructions != "" {
		trimmed := strings.TrimSpace(opts.SystemInstructions)
		if len(trimmed) == 0 {
			return &AnswerResponse{
				Success: false,
				Error:   "systemInstructions cannot be empty when provided",
			}, nil
		}
		if len(trimmed) > 2000 {
			return &AnswerResponse{
				Success: false,
				Error:   "systemInstructions must be 2000 characters or less",
			}, nil
		}
	}

	if opts.DataMaxPrice < 0 {
		return &AnswerResponse{
			Success: false,
			Error:   "dataMaxPrice must be a positive number",
		}, nil
	}

	if !validateDateFormat(opts.StartDate) {
		return &AnswerResponse{
			Success: false,
			Error:   "Invalid startDate format. Must be YYYY-MM-DD",
		}, nil
	}

	if !validateDateFormat(opts.EndDate) {
		return &AnswerResponse{
			Success: false,
			Error:   "Invalid endDate format. Must be YYYY-MM-DD",
		}, nil
	}

	for _, source := range opts.IncludedSources {
		if !validateSource(source) {
			return &AnswerResponse{
				Success: false,
				Error:   fmt.Sprintf("Invalid includedSources format. Invalid source: %s", source),
			}, nil
		}
	}

	for _, source := range opts.ExcludedSources {
		if !validateSource(source) {
			return &AnswerResponse{
				Success: false,
				Error:   fmt.Sprintf("Invalid excludedSources format. Invalid source: %s", source),
			}, nil
		}
	}

	req := answerRequest{
		Query:      query,
		SearchType: string(searchType),
	}

	if opts.StructuredOutput != nil {
		req.StructuredOutput = opts.StructuredOutput
	}
	if opts.SystemInstructions != "" {
		req.SystemInstructions = strings.TrimSpace(opts.SystemInstructions)
	}
	if opts.DataMaxPrice > 0 {
		req.DataMaxPrice = opts.DataMaxPrice
	}
	if opts.CountryCode != "" {
		req.CountryCode = string(opts.CountryCode)
	}
	if len(opts.IncludedSources) > 0 {
		req.IncludedSources = opts.IncludedSources
	}
	if len(opts.ExcludedSources) > 0 {
		req.ExcludedSources = opts.ExcludedSources
	}
	if opts.StartDate != "" {
		req.StartDate = opts.StartDate
	}
	if opts.EndDate != "" {
		req.EndDate = opts.EndDate
	}
	if opts.FastMode {
		req.FastMode = opts.FastMode
	}

	resp, err := c.doRequestRaw(ctx, "POST", "/answer", req)
	if err != nil {
		return &AnswerResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}
	defer resp.Body.Close()

	var fullContent strings.Builder
	var searchResults []SearchResult
	var finalResponse AnswerResponse

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue
		}

		dataStr := strings.TrimPrefix(line, "data: ")
		if dataStr == "[DONE]" {
			continue
		}

		var parsed map[string]interface{}
		if err := json.Unmarshal([]byte(dataStr), &parsed); err != nil {
			continue
		}

		if results, ok := parsed["search_results"]; ok {
			if _, hasSuccess := parsed["success"]; !hasSuccess {
				if resultsData, err := json.Marshal(results); err == nil {
					var sr []SearchResult
					if json.Unmarshal(resultsData, &sr) == nil {
						searchResults = append(searchResults, sr...)
					}
				}
			}
		}

		if choices, ok := parsed["choices"].([]interface{}); ok && len(choices) > 0 {
			if choice, ok := choices[0].(map[string]interface{}); ok {
				if delta, ok := choice["delta"].(map[string]interface{}); ok {
					if content, ok := delta["content"].(string); ok {
						fullContent.WriteString(content)
					}
				}
			}
		}

		if success, ok := parsed["success"]; ok {
			if successBool, ok := success.(bool); ok && successBool {
				finalResponse.Success = true
				if txID, ok := parsed["tx_id"].(string); ok {
					finalResponse.TxID = txID
				}
				if originalQuery, ok := parsed["original_query"].(string); ok {
					finalResponse.OriginalQuery = originalQuery
				}
				if dataType, ok := parsed["data_type"].(string); ok {
					finalResponse.DataType = dataType
				}
				if contents, ok := parsed["contents"]; ok {
					finalResponse.Contents = contents
				}

				if sm, ok := parsed["search_metadata"].(map[string]interface{}); ok {
					if smData, err := json.Marshal(sm); err == nil {
						json.Unmarshal(smData, &finalResponse.SearchMetadata)
					}
				}

				if au, ok := parsed["ai_usage"].(map[string]interface{}); ok {
					if auData, err := json.Marshal(au); err == nil {
						json.Unmarshal(auData, &finalResponse.AIUsage)
					}
				}

				if cost, ok := parsed["cost"].(map[string]interface{}); ok {
					if costData, err := json.Marshal(cost); err == nil {
						json.Unmarshal(costData, &finalResponse.Cost)
					}
				}

				if sr, ok := parsed["search_results"]; ok {
					if srData, err := json.Marshal(sr); err == nil {
						json.Unmarshal(srData, &finalResponse.SearchResults)
					}
				} else if len(searchResults) > 0 {
					finalResponse.SearchResults = searchResults
				}
			} else {
				finalResponse.Success = false
				if errMsg, ok := parsed["error"].(string); ok {
					finalResponse.Error = errMsg
				}
			}
		}
	}

	if finalResponse.Contents == nil && fullContent.Len() > 0 {
		finalResponse.Contents = fullContent.String()
		finalResponse.Success = true
	}

	if len(searchResults) > 0 && len(finalResponse.SearchResults) == 0 {
		finalResponse.SearchResults = searchResults
	}

	if finalResponse.Contents != nil || fullContent.Len() > 0 {
		finalResponse.Success = true
	}

	if !finalResponse.Success && finalResponse.Error == "" {
		finalResponse.Error = "Unknown error occurred"
	}

	return &finalResponse, nil
}

func (c *Client) AnswerStream(ctx context.Context, query string, opts *AnswerOptions) (<-chan AnswerStreamChunk, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		ch := make(chan AnswerStreamChunk, 1)
		ch <- AnswerStreamChunk{
			Type:  "error",
			Error: "Query is required and must be a non-empty string",
		}
		close(ch)
		return ch, nil
	}

	if opts == nil {
		opts = &AnswerOptions{}
	}

	searchType := opts.SearchType
	if searchType == "" {
		searchType = SearchTypeAll
	}

	req := answerRequest{
		Query:      query,
		SearchType: string(searchType),
	}

	if opts.StructuredOutput != nil {
		req.StructuredOutput = opts.StructuredOutput
	}
	if opts.SystemInstructions != "" {
		req.SystemInstructions = strings.TrimSpace(opts.SystemInstructions)
	}
	if opts.DataMaxPrice > 0 {
		req.DataMaxPrice = opts.DataMaxPrice
	}
	if opts.CountryCode != "" {
		req.CountryCode = string(opts.CountryCode)
	}
	if len(opts.IncludedSources) > 0 {
		req.IncludedSources = opts.IncludedSources
	}
	if len(opts.ExcludedSources) > 0 {
		req.ExcludedSources = opts.ExcludedSources
	}
	if opts.StartDate != "" {
		req.StartDate = opts.StartDate
	}
	if opts.EndDate != "" {
		req.EndDate = opts.EndDate
	}
	if opts.FastMode {
		req.FastMode = opts.FastMode
	}

	ch := make(chan AnswerStreamChunk, 100)

	go func() {
		defer close(ch)

		resp, err := c.doRequestRaw(ctx, "POST", "/answer", req)
		if err != nil {
			ch <- AnswerStreamChunk{
				Type:  "error",
				Error: err.Error(),
			}
			return
		}
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.HasPrefix(line, "data: ") {
				continue
			}

			dataStr := strings.TrimPrefix(line, "data: ")
			if dataStr == "[DONE]" {
				ch <- AnswerStreamChunk{Type: "done"}
				continue
			}

			var parsed map[string]interface{}
			if err := json.Unmarshal([]byte(dataStr), &parsed); err != nil {
				continue
			}

			if results, ok := parsed["search_results"]; ok {
				if _, hasSuccess := parsed["success"]; !hasSuccess {
					if resultsData, err := json.Marshal(results); err == nil {
						var sr []SearchResult
						if json.Unmarshal(resultsData, &sr) == nil {
							ch <- AnswerStreamChunk{
								Type:          "search_results",
								SearchResults: sr,
							}
						}
					}
				}
			}

			if choices, ok := parsed["choices"].([]interface{}); ok && len(choices) > 0 {
				if choice, ok := choices[0].(map[string]interface{}); ok {
					if delta, ok := choice["delta"].(map[string]interface{}); ok {
						content := ""
						if c, ok := delta["content"].(string); ok {
							content = c
						}
						finishReason := ""
						if fr, ok := choice["finish_reason"].(string); ok {
							finishReason = fr
						}
						if content != "" || finishReason != "" {
							ch <- AnswerStreamChunk{
								Type:         "content",
								Content:      content,
								FinishReason: finishReason,
							}
						}
					}
				}
			}

			if success, ok := parsed["success"]; ok {
				chunk := AnswerStreamChunk{Type: "metadata"}
				if txID, ok := parsed["tx_id"].(string); ok {
					chunk.TxID = txID
				}
				if originalQuery, ok := parsed["original_query"].(string); ok {
					chunk.OriginalQuery = originalQuery
				}
				if dataType, ok := parsed["data_type"].(string); ok {
					chunk.DataType = dataType
				}

				if sm, ok := parsed["search_metadata"].(map[string]interface{}); ok {
					smObj := &SearchMetadata{}
					if smData, err := json.Marshal(sm); err == nil {
						json.Unmarshal(smData, smObj)
					}
					chunk.SearchMetadata = smObj
				}
				if au, ok := parsed["ai_usage"].(map[string]interface{}); ok {
					auObj := &AIUsage{}
					if auData, err := json.Marshal(au); err == nil {
						json.Unmarshal(auData, auObj)
					}
					chunk.AIUsage = auObj
				}
				if cost, ok := parsed["cost"].(map[string]interface{}); ok {
					costObj := &Cost{}
					if costData, err := json.Marshal(cost); err == nil {
						json.Unmarshal(costData, costObj)
					}
					chunk.Cost = costObj
				}
				if sr, ok := parsed["search_results"]; ok {
					if srData, err := json.Marshal(sr); err == nil {
						var searchResults []SearchResult
						json.Unmarshal(srData, &searchResults)
						chunk.SearchResults = searchResults
					}
				}

				if successBool, ok := success.(bool); ok && !successBool {
					chunk.Type = "error"
					if errMsg, ok := parsed["error"].(string); ok {
						chunk.Error = errMsg
					}
				}

				ch <- chunk
			}
		}
	}()

	return ch, nil
}
