package valyu

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type DeepResearchService struct {
	client *Client
}

func (s *DeepResearchService) Create(ctx context.Context, opts *DeepResearchCreateOptions) (*DeepResearchCreateResponse, error) {
	if opts == nil || opts.Query == "" {
		return &DeepResearchCreateResponse{
			Success: false,
			Error:   "query is required and cannot be empty",
		}, nil
	}

	payload := map[string]interface{}{
		"query":          opts.Query,
		"mode":           opts.Mode,
		"output_formats": opts.OutputFormats,
		"code_execution": true,
	}

	if opts.Mode == "" {
		payload["mode"] = DeepResearchModeFast
	}
	if opts.OutputFormats == nil {
		payload["output_formats"] = []string{"markdown"}
	}
	if opts.CodeExecution != nil {
		payload["code_execution"] = *opts.CodeExecution
	}

	if opts.Strategy != "" {
		payload["strategy"] = opts.Strategy
	}
	if opts.Search != nil {
		search := map[string]interface{}{}
		if opts.Search.SearchType != "" {
			search["search_type"] = opts.Search.SearchType
		}
		if len(opts.Search.IncludedSources) > 0 {
			search["included_sources"] = opts.Search.IncludedSources
		}
		if len(opts.Search.ExcludedSources) > 0 {
			search["excluded_sources"] = opts.Search.ExcludedSources
		}
		if opts.Search.StartDate != "" {
			search["start_date"] = opts.Search.StartDate
		}
		if opts.Search.EndDate != "" {
			search["end_date"] = opts.Search.EndDate
		}
		if opts.Search.Category != "" {
			search["category"] = opts.Search.Category
		}
		payload["search"] = search
	}
	if len(opts.URLs) > 0 {
		payload["urls"] = opts.URLs
	}
	if len(opts.Files) > 0 {
		payload["files"] = opts.Files
	}
	if len(opts.Deliverables) > 0 {
		payload["deliverables"] = opts.Deliverables
	}
	if len(opts.MCPServers) > 0 {
		payload["mcp_servers"] = opts.MCPServers
	}
	if len(opts.PreviousReports) > 0 {
		payload["previous_reports"] = opts.PreviousReports
	}
	if opts.WebhookURL != "" {
		payload["webhook_url"] = opts.WebhookURL
	}
	if opts.BrandCollectionID != "" {
		payload["brand_collection_id"] = opts.BrandCollectionID
	}
	if opts.Metadata != nil {
		payload["metadata"] = opts.Metadata
	}

	var response DeepResearchCreateResponse
	if err := s.client.doRequest(ctx, "POST", "/deepresearch/tasks", payload, &response); err != nil {
		return &DeepResearchCreateResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *DeepResearchService) Status(ctx context.Context, taskID string) (*DeepResearchStatusResponse, error) {
	var response DeepResearchStatusResponse
	if err := s.client.doRequest(ctx, "GET", fmt.Sprintf("/deepresearch/tasks/%s/status", taskID), nil, &response); err != nil {
		return &DeepResearchStatusResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *DeepResearchService) Wait(ctx context.Context, taskID string, opts *WaitOptions) (*DeepResearchStatusResponse, error) {
	if opts == nil {
		opts = &WaitOptions{}
	}

	pollInterval := opts.PollInterval
	if pollInterval == 0 {
		pollInterval = 5000
	}

	maxWaitTime := opts.MaxWaitTime
	if maxWaitTime == 0 {
		maxWaitTime = 7200000
	}

	startTime := time.Now()

	for {
		status, err := s.Status(ctx, taskID)
		if err != nil {
			return nil, err
		}

		if !status.Success {
			return nil, fmt.Errorf(status.Error)
		}

		if opts.OnProgress != nil {
			opts.OnProgress(status)
		}

		if status.Status == DeepResearchStatusCompleted ||
			status.Status == DeepResearchStatusFailed ||
			status.Status == DeepResearchStatusCancelled {
			return status, nil
		}

		if time.Since(startTime).Milliseconds() > maxWaitTime {
			return nil, fmt.Errorf("maximum wait time exceeded")
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Duration(pollInterval) * time.Millisecond):
		}
	}
}

func (s *DeepResearchService) Stream(ctx context.Context, taskID string, callback StreamCallback) error {
	lastMessageCount := 0

	for {
		status, err := s.Status(ctx, taskID)
		if err != nil {
			if callback.OnError != nil {
				callback.OnError(err)
			}
			return err
		}

		if !status.Success {
			err := fmt.Errorf(status.Error)
			if callback.OnError != nil {
				callback.OnError(err)
			}
			return err
		}

		if status.Progress != nil && callback.OnProgress != nil {
			callback.OnProgress(status.Progress.CurrentStep, status.Progress.TotalSteps)
		}

		if len(status.Messages) > lastMessageCount && callback.OnMessage != nil {
			for i := lastMessageCount; i < len(status.Messages); i++ {
				callback.OnMessage(status.Messages[i])
			}
			lastMessageCount = len(status.Messages)
		}

		if status.Status == DeepResearchStatusCompleted {
			if callback.OnComplete != nil {
				callback.OnComplete(status)
			}
			return nil
		} else if status.Status == DeepResearchStatusFailed || status.Status == DeepResearchStatusCancelled {
			err := fmt.Errorf("task %s", status.Status)
			if callback.OnError != nil {
				callback.OnError(err)
			}
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(5 * time.Second):
		}
	}
}

func (s *DeepResearchService) List(ctx context.Context, opts *ListOptions) (*DeepResearchListResponse, error) {
	if opts == nil {
		return &DeepResearchListResponse{
			Success: false,
			Error:   "apiKeyId is required",
		}, nil
	}

	limit := opts.Limit
	if limit == 0 {
		limit = 10
	}

	path := fmt.Sprintf("/deepresearch/list?api_key_id=%s&limit=%d", url.QueryEscape(opts.APIKeyID), limit)

	var response DeepResearchListResponse
	if err := s.client.doRequest(ctx, "GET", path, nil, &response); err != nil {
		return &DeepResearchListResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *DeepResearchService) Update(ctx context.Context, taskID string, instruction string) (*DeepResearchUpdateResponse, error) {
	if instruction == "" {
		return &DeepResearchUpdateResponse{
			Success: false,
			Error:   "instruction is required and cannot be empty",
		}, nil
	}

	payload := map[string]string{"instruction": instruction}

	var response DeepResearchUpdateResponse
	if err := s.client.doRequest(ctx, "POST", fmt.Sprintf("/deepresearch/tasks/%s/update", taskID), payload, &response); err != nil {
		return &DeepResearchUpdateResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *DeepResearchService) Cancel(ctx context.Context, taskID string) (*DeepResearchCancelResponse, error) {
	var response DeepResearchCancelResponse
	if err := s.client.doRequest(ctx, "POST", fmt.Sprintf("/deepresearch/tasks/%s/cancel", taskID), map[string]interface{}{}, &response); err != nil {
		return &DeepResearchCancelResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *DeepResearchService) Delete(ctx context.Context, taskID string) (*DeepResearchDeleteResponse, error) {
	var response DeepResearchDeleteResponse
	if err := s.client.doRequest(ctx, "DELETE", fmt.Sprintf("/deepresearch/tasks/%s/delete", taskID), nil, &response); err != nil {
		return &DeepResearchDeleteResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *DeepResearchService) TogglePublic(ctx context.Context, taskID string, isPublic bool) (*DeepResearchTogglePublicResponse, error) {
	payload := map[string]bool{"public": isPublic}

	var response DeepResearchTogglePublicResponse
	if err := s.client.doRequest(ctx, "POST", fmt.Sprintf("/deepresearch/tasks/%s/public", taskID), payload, &response); err != nil {
		return &DeepResearchTogglePublicResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *DeepResearchService) GetAssets(ctx context.Context, taskID string, assetID string, opts *GetAssetsOptions) (*DeepResearchAssetsResponse, error) {
	path := fmt.Sprintf("/deepresearch/tasks/%s/assets/%s", taskID, assetID)
	if opts != nil && opts.Token != "" {
		path += "?token=" + url.QueryEscape(opts.Token)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", s.client.baseURL+path, nil)
	if err != nil {
		return &DeepResearchAssetsResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	if opts == nil || opts.Token == "" {
		req.Header.Set("x-api-key", s.client.apiKey)
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return &DeepResearchAssetsResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &DeepResearchAssetsResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &DeepResearchAssetsResponse{
		Success:     true,
		Data:        data,
		ContentType: resp.Header.Get("Content-Type"),
	}, nil
}
