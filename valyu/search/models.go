package search

import "github.com/Veri5ied/valyu-go/valyu/common"

type Options struct {
	Query              string                `json:"q"`
	SearchType         common.SearchType     `json:"search_type,omitempty"`
	MaxNumResults      int                   `json:"max_num_results,omitempty"`
	MaxPrice           float64               `json:"max_price,omitempty"`
	IsToolCall         *bool                 `json:"is_tool_call,omitempty"`
	RelevanceThreshold float64               `json:"relevance_threshold,omitempty"`
	IncludedSources    []string              `json:"included_sources,omitempty"`
	ExcludeSources     []string              `json:"exclude_sources,omitempty"`
	Category           string                `json:"category,omitempty"`
	StartDate          string                `json:"start_date,omitempty"`
	EndDate            string                `json:"end_date,omitempty"`
	CountryCode        common.CountryCode    `json:"country_code,omitempty"`
	ResponseLength     common.ResponseLength `json:"response_length,omitempty"`
	FastMode           bool                  `json:"fast_mode,omitempty"`
	URLOnly            bool                  `json:"url_only,omitempty"`
}

type Result struct {
	Title           string            `json:"title"`
	URL             string            `json:"url"`
	Content         interface{}       `json:"content"`
	Description     string            `json:"description,omitempty"`
	Source          string            `json:"source"`
	SourceType      string            `json:"source_type,omitempty"`
	DataType        string            `json:"data_type,omitempty"`
	Date            string            `json:"date,omitempty"`
	Length          int               `json:"length"`
	RelevanceScore  float64           `json:"relevance_score,omitempty"`
	PublicationDate string            `json:"publication_date,omitempty"`
	ID              string            `json:"id,omitempty"`
	ImageURL        map[string]string `json:"image_url,omitempty"`
}

type ResultsBySource struct {
	Web         int `json:"web"`
	Proprietary int `json:"proprietary"`
}

type Response struct {
	Success               bool            `json:"success"`
	Error                 string          `json:"error,omitempty"`
	TxID                  string          `json:"tx_id"`
	Query                 string          `json:"query"`
	Results               []Result        `json:"results"`
	ResultsBySource       ResultsBySource `json:"results_by_source"`
	TotalDeductionPCM     float64         `json:"total_deduction_pcm"`
	TotalDeductionDollars float64         `json:"total_deduction_dollars"`
	TotalCharacters       int             `json:"total_characters"`
}
