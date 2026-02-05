package contents

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

func (s *Service) Get(ctx context.Context, urls []string, opts *Options) (*Response, error) {
	req := struct {
		URLs []string `json:"urls"`
		*Options
	}{
		URLs:    urls,
		Options: opts,
	}

	var resp Response
	if err := s.client.Post(ctx, "/contents", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
