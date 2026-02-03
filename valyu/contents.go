package valyu

import (
	"context"
	"fmt"
)

type contentsRequest struct {
	URLs            []string    `json:"urls"`
	Summary         interface{} `json:"summary,omitempty"`
	ExtractEffort   string      `json:"extract_effort,omitempty"`
	ResponseLength  string      `json:"response_length,omitempty"`
	MaxPriceDollars float64     `json:"max_price_dollars,omitempty"`
	Screenshot      bool        `json:"screenshot,omitempty"`
}

func (c *Client) Contents(ctx context.Context, urls []string, opts *ContentsOptions) (*ContentsResponse, error) {
	if urls == nil || len(urls) == 0 {
		return &ContentsResponse{
			Success: false,
			Error:   "urls array cannot be empty",
		}, nil
	}

	if len(urls) > 10 {
		return &ContentsResponse{
			Success:       false,
			Error:         "Maximum 10 URLs allowed per request",
			URLsRequested: len(urls),
			URLsFailed:    len(urls),
		}, nil
	}

	if opts == nil {
		opts = &ContentsOptions{}
	}

	if opts.ExtractEffort != "" {
		validEfforts := map[ExtractEffort]bool{
			ExtractEffortNormal: true,
			ExtractEffortHigh:   true,
			ExtractEffortAuto:   true,
		}
		if !validEfforts[opts.ExtractEffort] {
			return &ContentsResponse{
				Success:       false,
				Error:         "extractEffort must be 'normal', 'high', or 'auto'",
				URLsRequested: len(urls),
				URLsFailed:    len(urls),
			}, nil
		}
	}

	if opts.ResponseLength != "" {
		validLengths := map[ResponseLength]bool{
			ResponseLengthShort:  true,
			ResponseLengthMedium: true,
			ResponseLengthLarge:  true,
			ResponseLengthMax:    true,
		}
		if !validLengths[opts.ResponseLength] {
			return &ContentsResponse{
				Success:       false,
				Error:         "responseLength must be 'short', 'medium', 'large', or 'max'",
				URLsRequested: len(urls),
				URLsFailed:    len(urls),
			}, nil
		}
	}

	req := contentsRequest{
		URLs: urls,
	}

	if opts.Summary != nil {
		req.Summary = opts.Summary
	}
	if opts.ExtractEffort != "" {
		req.ExtractEffort = string(opts.ExtractEffort)
	}
	if opts.ResponseLength != "" {
		req.ResponseLength = string(opts.ResponseLength)
	}
	if opts.MaxPriceDollars > 0 {
		req.MaxPriceDollars = opts.MaxPriceDollars
	}
	if opts.Screenshot {
		req.Screenshot = opts.Screenshot
	}

	var response ContentsResponse
	if err := c.doRequest(ctx, "POST", "/contents", req, &response); err != nil {
		return &ContentsResponse{
			Success:       false,
			Error:         fmt.Sprintf("request failed: %v", err),
			URLsRequested: len(urls),
			URLsFailed:    len(urls),
		}, nil
	}

	return &response, nil
}
