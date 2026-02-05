package batch

import (
	"github.com/Veri5ied/valyu-go/valyu/common"
)

type SearchParams struct {
	SearchType      string             `json:"search_type,omitempty"`
	IncludedSources []string           `json:"included_sources,omitempty"`
	ExcludedSources []string           `json:"excluded_sources,omitempty"`
	StartDate       string             `json:"start_date,omitempty"`
	EndDate         string             `json:"end_date,omitempty"`
	Category        string             `json:"category,omitempty"`
	CountryCode     common.CountryCode `json:"country_code,omitempty"`
}

type Counts struct {
	Total     int `json:"total"`
	Queued    int `json:"queued"`
	Running   int `json:"running"`
	Completed int `json:"completed"`
	Failed    int `json:"failed"`
	Cancelled int `json:"cancelled"`
}

type Batch struct {
	BatchID        string                  `json:"batch_id"`
	OrganisationID string                  `json:"organisation_id"`
	APIKeyID       string                  `json:"api_key_id"`
	CreditID       string                  `json:"credit_id"`
	Status         common.BatchStatus      `json:"status"`
	Mode           common.DeepResearchMode `json:"mode"`
	Name           string                  `json:"name,omitempty"`
	OutputFormats  []string                `json:"output_formats,omitempty"`
	SearchParams   *SearchParams           `json:"search_params,omitempty"`
	CreatedAt      string                  `json:"created_at"`
	CompletedAt    string                  `json:"completed_at,omitempty"`
	Counts         Counts                  `json:"counts"`
	Cost           float64                 `json:"cost"`
	WebhookSecret  string                  `json:"webhook_secret,omitempty"`
	Metadata       map[string]interface{}  `json:"metadata,omitempty"`
}

type CreateOptions struct {
	Name          string                  `json:"name,omitempty"`
	Mode          common.DeepResearchMode `json:"mode,omitempty"`
	OutputFormats []string                `json:"output_formats,omitempty"`
	Search        *SearchParams           `json:"search,omitempty"`
	WebhookURL    string                  `json:"webhook_url,omitempty"`
	Metadata      map[string]interface{}  `json:"metadata,omitempty"`
}

type CreateResponse struct {
	Success       bool                    `json:"success"`
	Error         string                  `json:"error,omitempty"`
	BatchID       string                  `json:"batch_id,omitempty"`
	Status        common.BatchStatus      `json:"status,omitempty"`
	Mode          common.DeepResearchMode `json:"mode,omitempty"`
	Name          string                  `json:"name,omitempty"`
	OutputFormats []string                `json:"output_formats,omitempty"`
	CreatedAt     string                  `json:"created_at,omitempty"`
	Counts        *Counts                 `json:"counts,omitempty"`
	Cost          float64                 `json:"cost,omitempty"`
	WebhookSecret string                  `json:"webhook_secret,omitempty"`
}

type ListResponse struct {
	Success bool    `json:"success"`
	Error   string  `json:"error,omitempty"`
	Batches []Batch `json:"batches,omitempty"`
}

type StatusResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Batch   *Batch `json:"batch,omitempty"`
}
