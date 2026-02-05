package search

import (
	"context"

	"github.com/Veri5ied/valyu-go/valyu/internal/api"
)

type Service struct {
	client *api.Client
}

func New(client *api.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Search(ctx context.Context, query string, opts *Options) (*Response, error) {
	req := struct {
		Query string `json:"query"`
		*Options
	}{
		Query:   query,
		Options: opts,
	}

	var resp Response
	if err := s.client.Post(ctx, "/search", req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
