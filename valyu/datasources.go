package valyu

import (
	"context"
	"net/url"
)

type DatasourcesService struct {
	client *Client
}

func (s *DatasourcesService) List(ctx context.Context, opts *DatasourcesListOptions) (*DatasourcesListResponse, error) {
	params := url.Values{}

	if opts != nil && opts.Category != "" {
		params.Set("category", string(opts.Category))
	}

	path := "/datasources"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var response DatasourcesListResponse
	if err := s.client.doRequest(ctx, "GET", path, nil, &response); err != nil {
		return &DatasourcesListResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}

func (s *DatasourcesService) Categories(ctx context.Context) (*DatasourcesCategoriesResponse, error) {
	var response DatasourcesCategoriesResponse
	if err := s.client.doRequest(ctx, "GET", "/datasources/categories", nil, &response); err != nil {
		return &DatasourcesCategoriesResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &response, nil
}
