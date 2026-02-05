package datasources

import "github.com/Veri5ied/valyu-go/valyu/common"

type Pricing struct {
	CPM float64 `json:"cpm"`
}

type Coverage struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type Datasource struct {
	ID              string                      `json:"id"`
	Name            string                      `json:"name"`
	Description     string                      `json:"description"`
	Category        common.DatasourceCategoryID `json:"category"`
	Type            string                      `json:"type"`
	Modality        []string                    `json:"modality"`
	Topics          []string                    `json:"topics"`
	Languages       []string                    `json:"languages,omitempty"`
	Source          string                      `json:"source,omitempty"`
	ExampleQueries  []string                    `json:"example_queries"`
	Pricing         Pricing                     `json:"pricing"`
	ResponseSchema  map[string]interface{}      `json:"response_schema,omitempty"`
	UpdateFrequency string                      `json:"update_frequency,omitempty"`
	Size            int                         `json:"size,omitempty"`
	Coverage        *Coverage                   `json:"coverage,omitempty"`
}

type ListResponse struct {
	Success     bool         `json:"success"`
	Error       string       `json:"error,omitempty"`
	Datasources []Datasource `json:"datasources,omitempty"`
}
