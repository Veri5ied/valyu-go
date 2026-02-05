package deepresearch

import "github.com/Veri5ied/valyu-go/valyu/common"

type SearchConfig struct {
	SearchType      string             `json:"search_type,omitempty"`
	IncludedSources []string           `json:"included_sources,omitempty"`
	ExcludedSources []string           `json:"excluded_sources,omitempty"`
	StartDate       string             `json:"start_date,omitempty"`
	EndDate         string             `json:"end_date,omitempty"`
	Category        string             `json:"category,omitempty"`
	CountryCode     common.CountryCode `json:"country_code,omitempty"`
}

type FileAttachment struct {
	Data      string `json:"data"`
	Filename  string `json:"filename"`
	MediaType string `json:"mediaType"`
	Context   string `json:"context,omitempty"`
}

type MCPAuth struct {
	Type    string            `json:"type"`
	Token   string            `json:"token,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

type MCPServerConfig struct {
	URL          string   `json:"url"`
	Name         string   `json:"name,omitempty"`
	ToolPrefix   string   `json:"toolPrefix,omitempty"`
	Auth         *MCPAuth `json:"auth,omitempty"`
	AllowedTools []string `json:"allowedTools,omitempty"`
}

type CreateOptions struct {
	Query             string                  `json:"query"`
	Mode              common.DeepResearchMode `json:"mode,omitempty"`
	OutputFormats     []string                `json:"output_formats,omitempty"`
	Strategy          string                  `json:"strategy,omitempty"`
	Search            *SearchConfig           `json:"search,omitempty"`
	URLs              []string                `json:"urls,omitempty"`
	Files             []FileAttachment        `json:"files,omitempty"`
	Deliverables      []interface{}           `json:"deliverables,omitempty"`
	MCPServers        []MCPServerConfig       `json:"mcp_servers,omitempty"`
	CodeExecution     *bool                   `json:"code_execution,omitempty"`
	PreviousReports   []string                `json:"previous_reports,omitempty"`
	WebhookURL        string                  `json:"webhook_url,omitempty"`
	BrandCollectionID string                  `json:"brand_collection_id,omitempty"`
	Metadata          map[string]interface{}  `json:"metadata,omitempty"`
}

type CreateResponse struct {
	Success        bool                      `json:"success"`
	Error          string                    `json:"error,omitempty"`
	DeepResearchID string                    `json:"deepresearch_id,omitempty"`
	Status         common.DeepResearchStatus `json:"status,omitempty"`
	Mode           common.DeepResearchMode   `json:"mode,omitempty"`
	CreatedAt      string                    `json:"created_at,omitempty"`
	Metadata       map[string]interface{}    `json:"metadata,omitempty"`
	Public         bool                      `json:"public,omitempty"`
	WebhookSecret  string                    `json:"webhook_secret,omitempty"`
	Message        string                    `json:"message,omitempty"`
}

type Progress struct {
	CurrentStep int `json:"current_step"`
	TotalSteps  int `json:"total_steps"`
}

type Source struct {
	Title       string  `json:"title"`
	URL         string  `json:"url"`
	Snippet     string  `json:"snippet,omitempty"`
	Description string  `json:"description,omitempty"`
	Source      string  `json:"source,omitempty"`
	Price       float64 `json:"price,omitempty"`
	ID          string  `json:"id,omitempty"`
	DOI         string  `json:"doi,omitempty"`
	Category    string  `json:"category,omitempty"`
}

type Usage struct {
	SearchCost   float64 `json:"search_cost"`
	ContentsCost float64 `json:"contents_cost"`
	AICost       float64 `json:"ai_cost"`
	ComputeCost  float64 `json:"compute_cost"`
	TotalCost    float64 `json:"total_cost"`
}

type ImageMetadata struct {
	ImageID        string `json:"image_id"`
	ImageType      string `json:"image_type"`
	DeepResearchID string `json:"deepresearch_id"`
	Title          string `json:"title"`
	Description    string `json:"description,omitempty"`
	ImageURL       string `json:"image_url"`
	S3Key          string `json:"s3_key"`
	CreatedAt      int64  `json:"created_at"`
	ChartType      string `json:"chart_type,omitempty"`
}

type DeliverableResult struct {
	ID          string `json:"id"`
	Request     string `json:"request"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
	S3Key       string `json:"s3_key"`
	RowCount    int    `json:"row_count,omitempty"`
	ColumnCount int    `json:"column_count,omitempty"`
	Error       string `json:"error,omitempty"`
	CreatedAt   int64  `json:"created_at"`
}

type StatusResponse struct {
	Success        bool                      `json:"success"`
	Error          string                    `json:"error,omitempty"`
	DeepResearchID string                    `json:"deepresearch_id,omitempty"`
	Status         common.DeepResearchStatus `json:"status,omitempty"`
	Query          string                    `json:"query,omitempty"`
	Mode           common.DeepResearchMode   `json:"mode,omitempty"`
	OutputFormats  []string                  `json:"output_formats,omitempty"`
	CreatedAt      string                    `json:"created_at,omitempty"`
	Public         bool                      `json:"public,omitempty"`
	Progress       *Progress                 `json:"progress,omitempty"`
	Messages       []interface{}             `json:"messages,omitempty"`
	CompletedAt    string                    `json:"completed_at,omitempty"`
	Output         interface{}               `json:"output,omitempty"`
	OutputType     string                    `json:"output_type,omitempty"`
	PDFURL         string                    `json:"pdf_url,omitempty"`
	Images         []ImageMetadata           `json:"images,omitempty"`
	Deliverables   []DeliverableResult       `json:"deliverables,omitempty"`
	Sources        []Source                  `json:"sources,omitempty"`
	Cost           float64                   `json:"cost,omitempty"`
	Usage          *Usage                    `json:"usage,omitempty"`
	BatchID        string                    `json:"batch_id,omitempty"`
	BatchTaskID    string                    `json:"batch_task_id,omitempty"`
}

type ListItem struct {
	DeepResearchID string                    `json:"deepresearch_id"`
	Query          string                    `json:"query"`
	Status         common.DeepResearchStatus `json:"status"`
	CreatedAt      int64                     `json:"created_at"`
	Public         bool                      `json:"public,omitempty"`
}

type ListResponse struct {
	Success bool       `json:"success"`
	Error   string     `json:"error,omitempty"`
	Data    []ListItem `json:"data,omitempty"`
}
