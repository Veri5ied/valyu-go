package api

import (
	"fmt"
)

type APIError struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message"`
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("valyu: %s (status: %d) - %s", e.Code, e.StatusCode, e.Message)
	}
	return fmt.Sprintf("valyu: error (status: %d) - %s", e.StatusCode, e.Message)
}
