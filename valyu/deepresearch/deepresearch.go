package deepresearch

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
	if err := s.client.Post(ctx, "/deepresearch", opts, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Get(ctx context.Context, id string) (*StatusResponse, error) {
	var resp StatusResponse
	if err := s.client.Get(ctx, fmt.Sprintf("/deepresearch/%s", id), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) List(ctx context.Context, opts ...interface{}) (*ListResponse, error) {
	var resp ListResponse
	if err := s.client.Get(ctx, "/deepresearch", &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
