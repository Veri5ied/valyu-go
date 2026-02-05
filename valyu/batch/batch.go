package batch

import (
	"context"
	"fmt"

	"github.com/Veri5ied/valyu-go/valyu/internal/api"
)

type Service struct {
	client *api.Client
}

func New(client *api.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Create(ctx context.Context, opts *CreateOptions) (*CreateResponse, error) {
	var resp CreateResponse
	if err := s.client.Post(ctx, "/batch", opts, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	var resp ListResponse
	if err := s.client.Get(ctx, "/batch", &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Get(ctx context.Context, batchID string) (*StatusResponse, error) {
	var resp StatusResponse
	if err := s.client.Get(ctx, fmt.Sprintf("/batch/%s", batchID), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
