package valyu

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type BatchService struct {
	client *Client
}

func (s *BatchService) Create(ctx context.Context, opts *CreateBatchOptions) (*CreateBatchResponse, error) {
	if opts == nil {
		opts = &CreateBatchOptions{}
	}

	payload := make(map[string]interface{})

	if opts.Name != "" {
		payload["name"] = opts.Name
	}
	if opts.Mode != "" {
		payload["mode"] = opts.Mode
	}
	if len(opts.OutputFormats) > 0 {
		payload["output_formats"] = opts.OutputFormats
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
	if opts.WebhookURL != "" {
		payload["webhook_url"] = opts.WebhookURL
	}
	if opts.Metadata != nil {
		payload["metadata"] = opts.Metadata
	}

	var response CreateBatchResponse
	if err := s.client.doRequest(ctx, "POST", "/deepresearch/batches", payload, &response); err != nil {
		return &CreateBatchResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *BatchService) Status(ctx context.Context, batchID string) (*BatchStatusResponse, error) {
	var response BatchStatusResponse
	if err := s.client.doRequest(ctx, "GET", fmt.Sprintf("/deepresearch/batches/%s", batchID), nil, &response); err != nil {
		return &BatchStatusResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *BatchService) AddTasks(ctx context.Context, batchID string, opts *AddBatchTasksOptions) (*AddBatchTasksResponse, error) {
	if opts == nil || len(opts.Tasks) == 0 {
		return &AddBatchTasksResponse{
			Success: false,
			Error:   "tasks array cannot be empty",
		}, nil
	}

	if len(opts.Tasks) > 100 {
		return &AddBatchTasksResponse{
			Success: false,
			Error:   "Maximum 100 tasks allowed per request",
		}, nil
	}

	for _, task := range opts.Tasks {
		if task.Query == "" {
			return &AddBatchTasksResponse{
				Success: false,
				Error:   "Each task must have a 'query' field",
			}, nil
		}
	}

	tasks := make([]map[string]interface{}, len(opts.Tasks))
	for i, task := range opts.Tasks {
		taskPayload := map[string]interface{}{
			"query": task.Query,
		}
		if task.ID != "" {
			taskPayload["id"] = task.ID
		}
		if task.Strategy != "" {
			taskPayload["strategy"] = task.Strategy
		}
		if len(task.URLs) > 0 {
			taskPayload["urls"] = task.URLs
		}
		if task.Metadata != nil {
			taskPayload["metadata"] = task.Metadata
		}
		tasks[i] = taskPayload
	}

	payload := map[string]interface{}{
		"tasks": tasks,
	}

	var response AddBatchTasksResponse
	if err := s.client.doRequest(ctx, "POST", fmt.Sprintf("/deepresearch/batches/%s/tasks", batchID), payload, &response); err != nil {
		return &AddBatchTasksResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *BatchService) ListTasks(ctx context.Context, batchID string, opts *ListBatchTasksOptions) (*ListBatchTasksResponse, error) {
	params := url.Values{}

	if opts != nil {
		if opts.Status != "" {
			params.Set("status", string(opts.Status))
		}
		if opts.Limit > 0 {
			params.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.LastKey != "" {
			params.Set("last_key", opts.LastKey)
		}
	}

	path := fmt.Sprintf("/deepresearch/batches/%s/tasks", batchID)
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var response ListBatchTasksResponse
	if err := s.client.doRequest(ctx, "GET", path, nil, &response); err != nil {
		return &ListBatchTasksResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *BatchService) Cancel(ctx context.Context, batchID string) (*CancelBatchResponse, error) {
	var response CancelBatchResponse
	if err := s.client.doRequest(ctx, "POST", fmt.Sprintf("/deepresearch/batches/%s/cancel", batchID), map[string]interface{}{}, &response); err != nil {
		return &CancelBatchResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *BatchService) List(ctx context.Context, opts *ListBatchesOptions) (*ListBatchesResponse, error) {
	params := url.Values{}

	if opts != nil && opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}

	path := "/deepresearch/batches"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var response ListBatchesResponse
	if err := s.client.doRequest(ctx, "GET", path, nil, &response); err != nil {
		return &ListBatchesResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *BatchService) WaitForCompletion(ctx context.Context, batchID string, opts *BatchWaitOptions) (*Batch, error) {
	if opts == nil {
		opts = &BatchWaitOptions{}
	}

	pollInterval := opts.PollInterval
	if pollInterval == 0 {
		pollInterval = 10000
	}

	maxWaitTime := opts.MaxWaitTime
	if maxWaitTime == 0 {
		maxWaitTime = 7200000
	}

	startTime := time.Now()

	for {
		statusResponse, err := s.Status(ctx, batchID)
		if err != nil {
			return nil, err
		}

		if !statusResponse.Success || statusResponse.Batch == nil {
			return nil, fmt.Errorf(statusResponse.Error)
		}

		batch := statusResponse.Batch

		if opts.OnProgress != nil {
			opts.OnProgress(batch)
		}

		if batch.Status == BatchStatusCompleted ||
			batch.Status == BatchStatusCompletedWithErrors ||
			batch.Status == BatchStatusCancelled {
			return batch, nil
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
