package answer

import (
	"github.com/Veri5ied/valyu-go/valyu/common"
	"github.com/Veri5ied/valyu-go/valyu/search"
)

type Options struct {
	StructuredOutput   interface{}        `json:"structured_output,omitempty"`
	SystemInstructions string             `json:"system_instructions,omitempty"`
	SearchType         common.SearchType  `json:"search_type,omitempty"`
	DataMaxPrice       float64            `json:"data_max_price,omitempty"`
	CountryCode        common.CountryCode `json:"country_code,omitempty"`
	IncludedSources    []string           `json:"included_sources,omitempty"`
	ExcludedSources    []string           `json:"excluded_sources,omitempty"`
	StartDate          string             `json:"start_date,omitempty"`
	EndDate            string             `json:"end_date,omitempty"`
	FastMode           bool               `json:"fast_mode,omitempty"`
}

type SearchMetadata struct {
	TxIDs           []string `json:"tx_ids"`
	NumberOfResults int      `json:"number_of_results"`
	TotalCharacters int      `json:"total_characters"`
}

type AIUsage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type Cost struct {
	TotalDeductionDollars    float64 `json:"total_deduction_dollars"`
	SearchDeductionDollars   float64 `json:"search_deduction_dollars"`
	ContentsDeductionDollars float64 `json:"contents_deduction_dollars,omitempty"`
	AIDeductionDollars       float64 `json:"ai_deduction_dollars"`
}

type Response struct {
	Success        bool            `json:"success"`
	Error          string          `json:"error,omitempty"`
	TxID           string          `json:"tx_id,omitempty"`
	OriginalQuery  string          `json:"original_query,omitempty"`
	Contents       interface{}     `json:"contents,omitempty"`
	DataType       string          `json:"data_type,omitempty"`
	SearchResults  []search.Result `json:"search_results,omitempty"`
	SearchMetadata SearchMetadata  `json:"search_metadata,omitempty"`
	AIUsage        AIUsage         `json:"ai_usage,omitempty"`
	Cost           Cost            `json:"cost,omitempty"`
}

type StreamChunk struct {
	Type           string          `json:"type"`
	SearchResults  []search.Result `json:"search_results,omitempty"`
	Content        string          `json:"content,omitempty"`
	FinishReason   string          `json:"finish_reason,omitempty"`
	TxID           string          `json:"tx_id,omitempty"`
	OriginalQuery  string          `json:"original_query,omitempty"`
	DataType       string          `json:"data_type,omitempty"`
	SearchMetadata *SearchMetadata `json:"search_metadata,omitempty"`
	AIUsage        *AIUsage        `json:"ai_usage,omitempty"`
	Cost           *Cost           `json:"cost,omitempty"`
	Error          string          `json:"error,omitempty"`
}
