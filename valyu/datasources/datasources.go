package datasources

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

func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	var resp ListResponse
	if err := s.client.Get(ctx, "/datasources", &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
