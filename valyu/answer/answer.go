package answer

import (
	"bufio"
	"context"
	"encoding/json"
	"strings"

	"github.com/Veri5ied/valyu-go/valyu/internal/api"
	"github.com/Veri5ied/valyu-go/valyu/search"
)

type Service struct {
	client *api.Client
}

func New(client *api.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Answer(ctx context.Context, query string, opts *Options) (*Response, error) {
	streamCh, err := s.Stream(ctx, query, opts)
	if err != nil {
		return nil, err
	}

	finalResp := &Response{
		Success: true,
	}
	var fullContent strings.Builder

	for chunk := range streamCh {
		if chunk.Error != "" {
			finalResp.Error = chunk.Error
			finalResp.Success = false
		}
		if chunk.Type == "search_results" {
			finalResp.SearchResults = append(finalResp.SearchResults, chunk.SearchResults...)
		}
		if chunk.Type == "content" {
			fullContent.WriteString(chunk.Content)
		}
		if chunk.Type == "metadata" {
			if chunk.TxID != "" {
				finalResp.TxID = chunk.TxID
			}
			if chunk.OriginalQuery != "" {
				finalResp.OriginalQuery = chunk.OriginalQuery
			}
			if chunk.SearchMetadata != nil {
				finalResp.SearchMetadata = *chunk.SearchMetadata
			}
			if chunk.AIUsage != nil {
				finalResp.AIUsage = *chunk.AIUsage
			}
			if chunk.Cost != nil {
				finalResp.Cost = *chunk.Cost
			}
			if chunk.DataType != "" {
				finalResp.DataType = chunk.DataType
			}
		}
	}

	finalResp.Contents = fullContent.String()
	if finalResp.Error == "" {
		finalResp.Success = true
	}

	return finalResp, nil
}

func (s *Service) Stream(ctx context.Context, query string, opts *Options) (<-chan StreamChunk, error) {
	var reqOpts Options
	if opts != nil {
		reqOpts = *opts
	}
	if reqOpts.SearchType == "" {
		reqOpts.SearchType = "all"
	}

	req := struct {
		Query string `json:"query"`
		Options
	}{
		Query:   query,
		Options: reqOpts,
	}

	resp, err := s.client.PostStream(ctx, "/answer", req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		var errResp struct {
			Error string `json:"error"`
		}
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		if errResp.Error != "" {
			return nil, &api.APIError{StatusCode: resp.StatusCode, Message: errResp.Error}
		}
		return nil, &api.APIError{StatusCode: resp.StatusCode, Message: "Request failed"}
	}

	ch := make(chan StreamChunk, 100)
	go func() {
		defer close(ch)
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.HasPrefix(line, "data: ") {
				continue
			}

			dataStr := strings.TrimPrefix(line, "data: ")
			if dataStr == "[DONE]" {
				ch <- StreamChunk{Type: "done"}
				return
			}

			var parsed map[string]interface{}
			if err := json.Unmarshal([]byte(dataStr), &parsed); err != nil {
				continue
			}

			var chunk StreamChunk
			_ = json.Unmarshal([]byte(dataStr), &chunk)

			if errMsg, ok := parsed["error"].(string); ok && errMsg != "" {
				chunk.Type = "error"
				chunk.Error = errMsg
			} else if results, ok := parsed["search_results"]; ok {
				chunk.Type = "search_results"
				if len(chunk.SearchResults) == 0 {
					if b, err := json.Marshal(results); err == nil {
						var sr []search.Result
						if err := json.Unmarshal(b, &sr); err == nil {
							chunk.SearchResults = sr
						}
					}
				}
			} else if choices, ok := parsed["choices"].([]interface{}); ok && len(choices) > 0 {
				chunk.Type = "content"
				if choice, ok := choices[0].(map[string]interface{}); ok {
					if delta, ok := choice["delta"].(map[string]interface{}); ok {
						if content, ok := delta["content"].(string); ok {
							chunk.Content = content
						}
					}
					if fr, ok := choice["finish_reason"].(string); ok {
						chunk.FinishReason = fr
					}
				}
			} else if success, ok := parsed["success"]; ok {
				chunk.Type = "metadata"
				if successBool, ok := success.(bool); ok && !successBool {
					chunk.Type = "error"
					if errMsg, ok := parsed["error"].(string); ok {
						chunk.Error = errMsg
					}
				}
			}

			if chunk.Type != "" {
				ch <- chunk
			}
		}
	}()

	return ch, nil
}
