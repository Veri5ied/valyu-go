package contents

import "github.com/Veri5ied/valyu-go/valyu/common"

type Options struct {
	Summary         interface{}           `json:"summary,omitempty"`
	ExtractEffort   common.ExtractEffort  `json:"extract_effort,omitempty"`
	ResponseLength  common.ResponseLength `json:"response_length,omitempty"`
	MaxPriceDollars float64               `json:"max_price_dollars,omitempty"`
	Screenshot      bool                  `json:"screenshot,omitempty"`
}

type Result struct {
	URL            string            `json:"url"`
	Title          string            `json:"title"`
	Content        interface{}       `json:"content"`
	Length         int               `json:"length"`
	Source         string            `json:"source"`
	Price          float64           `json:"price"`
	Description    string            `json:"description,omitempty"`
	Summary        interface{}       `json:"summary,omitempty"`
	SummarySuccess bool              `json:"summary_success,omitempty"`
	DataType       string            `json:"data_type,omitempty"`
	ImageURL       map[string]string `json:"image_url,omitempty"`
	ScreenshotURL  string            `json:"screenshot_url,omitempty"`
	Citation       string            `json:"citation,omitempty"`
}

type Response struct {
	Success          bool     `json:"success"`
	Error            string   `json:"error,omitempty"`
	TxID             string   `json:"tx_id,omitempty"`
	URLsRequested    int      `json:"urls_requested,omitempty"`
	URLsProcessed    int      `json:"urls_processed,omitempty"`
	URLsFailed       int      `json:"urls_failed,omitempty"`
	Results          []Result `json:"results,omitempty"`
	TotalCostDollars float64  `json:"total_cost_dollars,omitempty"`
	TotalCharacters  int      `json:"total_characters,omitempty"`
}
