package valyu

type SearchOptions struct {
	SearchType         SearchType     `json:"search_type,omitempty"`
	MaxNumResults      int            `json:"max_num_results,omitempty"`
	MaxPrice           float64        `json:"max_price,omitempty"`
	IsToolCall         *bool          `json:"is_tool_call,omitempty"`
	RelevanceThreshold float64        `json:"relevance_threshold,omitempty"`
	IncludedSources    []string       `json:"included_sources,omitempty"`
	ExcludeSources     []string       `json:"exclude_sources,omitempty"`
	Category           string         `json:"category,omitempty"`
	StartDate          string         `json:"start_date,omitempty"`
	EndDate            string         `json:"end_date,omitempty"`
	CountryCode        CountryCode    `json:"country_code,omitempty"`
	ResponseLength     ResponseLength `json:"response_length,omitempty"`
	FastMode           bool           `json:"fast_mode,omitempty"`
	URLOnly            bool           `json:"url_only,omitempty"`
}

type SearchResult struct {
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

type SearchResponse struct {
	Success               bool            `json:"success"`
	Error                 string          `json:"error,omitempty"`
	TxID                  string          `json:"tx_id"`
	Query                 string          `json:"query"`
	Results               []SearchResult  `json:"results"`
	ResultsBySource       ResultsBySource `json:"results_by_source"`
	TotalDeductionPCM     float64         `json:"total_deduction_pcm"`
	TotalDeductionDollars float64         `json:"total_deduction_dollars"`
	TotalCharacters       int             `json:"total_characters"`
}

type ContentsOptions struct {
	Summary         interface{}    `json:"summary,omitempty"`
	ExtractEffort   ExtractEffort  `json:"extract_effort,omitempty"`
	ResponseLength  ResponseLength `json:"response_length,omitempty"`
	MaxPriceDollars float64        `json:"max_price_dollars,omitempty"`
	Screenshot      bool           `json:"screenshot,omitempty"`
}

type ContentResult struct {
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

type ContentsResponse struct {
	Success          bool            `json:"success"`
	Error            string          `json:"error,omitempty"`
	TxID             string          `json:"tx_id,omitempty"`
	URLsRequested    int             `json:"urls_requested,omitempty"`
	URLsProcessed    int             `json:"urls_processed,omitempty"`
	URLsFailed       int             `json:"urls_failed,omitempty"`
	Results          []ContentResult `json:"results,omitempty"`
	TotalCostDollars float64         `json:"total_cost_dollars,omitempty"`
	TotalCharacters  int             `json:"total_characters,omitempty"`
}

type AnswerOptions struct {
	StructuredOutput   interface{} `json:"structured_output,omitempty"`
	SystemInstructions string      `json:"system_instructions,omitempty"`
	SearchType         SearchType  `json:"search_type,omitempty"`
	DataMaxPrice       float64     `json:"data_max_price,omitempty"`
	CountryCode        CountryCode `json:"country_code,omitempty"`
	IncludedSources    []string    `json:"included_sources,omitempty"`
	ExcludedSources    []string    `json:"excluded_sources,omitempty"`
	StartDate          string      `json:"start_date,omitempty"`
	EndDate            string      `json:"end_date,omitempty"`
	FastMode           bool        `json:"fast_mode,omitempty"`
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

type AnswerResponse struct {
	Success        bool           `json:"success"`
	Error          string         `json:"error,omitempty"`
	TxID           string         `json:"tx_id,omitempty"`
	OriginalQuery  string         `json:"original_query,omitempty"`
	Contents       interface{}    `json:"contents,omitempty"`
	DataType       string         `json:"data_type,omitempty"`
	SearchResults  []SearchResult `json:"search_results,omitempty"`
	SearchMetadata SearchMetadata `json:"search_metadata,omitempty"`
	AIUsage        AIUsage        `json:"ai_usage,omitempty"`
	Cost           Cost           `json:"cost,omitempty"`
}

type AnswerStreamChunk struct {
	Type           string          `json:"type"`
	SearchResults  []SearchResult  `json:"search_results,omitempty"`
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

type DeepResearchSearchConfig struct {
	SearchType      string      `json:"search_type,omitempty"`
	IncludedSources []string    `json:"included_sources,omitempty"`
	ExcludedSources []string    `json:"excluded_sources,omitempty"`
	StartDate       string      `json:"start_date,omitempty"`
	EndDate         string      `json:"end_date,omitempty"`
	Category        string      `json:"category,omitempty"`
	CountryCode     CountryCode `json:"country_code,omitempty"`
}

type DeepResearchCreateOptions struct {
	Query             string                    `json:"query"`
	Mode              DeepResearchMode          `json:"mode,omitempty"`
	OutputFormats     []string                  `json:"output_formats,omitempty"`
	Strategy          string                    `json:"strategy,omitempty"`
	Search            *DeepResearchSearchConfig `json:"search,omitempty"`
	URLs              []string                  `json:"urls,omitempty"`
	Files             []FileAttachment          `json:"files,omitempty"`
	Deliverables      []interface{}             `json:"deliverables,omitempty"`
	MCPServers        []MCPServerConfig         `json:"mcp_servers,omitempty"`
	CodeExecution     *bool                     `json:"code_execution,omitempty"`
	PreviousReports   []string                  `json:"previous_reports,omitempty"`
	WebhookURL        string                    `json:"webhook_url,omitempty"`
	BrandCollectionID string                    `json:"brand_collection_id,omitempty"`
	Metadata          map[string]interface{}    `json:"metadata,omitempty"`
}

type Progress struct {
	CurrentStep int `json:"current_step"`
	TotalSteps  int `json:"total_steps"`
}

type DeepResearchSource struct {
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

type DeepResearchUsage struct {
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

type DeepResearchCreateResponse struct {
	Success        bool                   `json:"success"`
	Error          string                 `json:"error,omitempty"`
	DeepResearchID string                 `json:"deepresearch_id,omitempty"`
	Status         DeepResearchStatus     `json:"status,omitempty"`
	Mode           DeepResearchMode       `json:"mode,omitempty"`
	CreatedAt      string                 `json:"created_at,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	Public         bool                   `json:"public,omitempty"`
	WebhookSecret  string                 `json:"webhook_secret,omitempty"`
	Message        string                 `json:"message,omitempty"`
}

type DeepResearchStatusResponse struct {
	Success        bool                 `json:"success"`
	Error          string               `json:"error,omitempty"`
	DeepResearchID string               `json:"deepresearch_id,omitempty"`
	Status         DeepResearchStatus   `json:"status,omitempty"`
	Query          string               `json:"query,omitempty"`
	Mode           DeepResearchMode     `json:"mode,omitempty"`
	OutputFormats  []string             `json:"output_formats,omitempty"`
	CreatedAt      string               `json:"created_at,omitempty"`
	Public         bool                 `json:"public,omitempty"`
	Progress       *Progress            `json:"progress,omitempty"`
	Messages       []interface{}        `json:"messages,omitempty"`
	CompletedAt    string               `json:"completed_at,omitempty"`
	Output         interface{}          `json:"output,omitempty"`
	OutputType     string               `json:"output_type,omitempty"`
	PDFURL         string               `json:"pdf_url,omitempty"`
	Images         []ImageMetadata      `json:"images,omitempty"`
	Deliverables   []DeliverableResult  `json:"deliverables,omitempty"`
	Sources        []DeepResearchSource `json:"sources,omitempty"`
	Cost           float64              `json:"cost,omitempty"`
	Usage          *DeepResearchUsage   `json:"usage,omitempty"`
	BatchID        string               `json:"batch_id,omitempty"`
	BatchTaskID    string               `json:"batch_task_id,omitempty"`
}

type DeepResearchListItem struct {
	DeepResearchID string             `json:"deepresearch_id"`
	Query          string             `json:"query"`
	Status         DeepResearchStatus `json:"status"`
	CreatedAt      int64              `json:"created_at"`
	Public         bool               `json:"public,omitempty"`
}

type DeepResearchListResponse struct {
	Success bool                   `json:"success"`
	Error   string                 `json:"error,omitempty"`
	Data    []DeepResearchListItem `json:"data,omitempty"`
}

type DeepResearchUpdateResponse struct {
	Success        bool   `json:"success"`
	Error          string `json:"error,omitempty"`
	Message        string `json:"message,omitempty"`
	DeepResearchID string `json:"deepresearch_id,omitempty"`
}

type DeepResearchCancelResponse struct {
	Success        bool   `json:"success"`
	Error          string `json:"error,omitempty"`
	Message        string `json:"message,omitempty"`
	DeepResearchID string `json:"deepresearch_id,omitempty"`
}

type DeepResearchDeleteResponse struct {
	Success        bool   `json:"success"`
	Error          string `json:"error,omitempty"`
	Message        string `json:"message,omitempty"`
	DeepResearchID string `json:"deepresearch_id,omitempty"`
}

type DeepResearchTogglePublicResponse struct {
	Success        bool   `json:"success"`
	Error          string `json:"error,omitempty"`
	Message        string `json:"message,omitempty"`
	DeepResearchID string `json:"deepresearch_id,omitempty"`
	Public         bool   `json:"public,omitempty"`
}

type DeepResearchAssetsResponse struct {
	Success     bool   `json:"success"`
	Error       string `json:"error,omitempty"`
	Data        []byte `json:"-"`
	ContentType string `json:"-"`
}

type WaitOptions struct {
	PollInterval int64
	MaxWaitTime  int64
	OnProgress   func(status *DeepResearchStatusResponse)
}

type StreamCallback struct {
	OnMessage  func(message interface{})
	OnProgress func(current, total int)
	OnComplete func(result *DeepResearchStatusResponse)
	OnError    func(err error)
}

type ListOptions struct {
	APIKeyID string
	Limit    int
}

type BatchCounts struct {
	Total     int `json:"total"`
	Queued    int `json:"queued"`
	Running   int `json:"running"`
	Completed int `json:"completed"`
	Failed    int `json:"failed"`
	Cancelled int `json:"cancelled"`
}

type Batch struct {
	BatchID        string                    `json:"batch_id"`
	OrganisationID string                    `json:"organisation_id"`
	APIKeyID       string                    `json:"api_key_id"`
	CreditID       string                    `json:"credit_id"`
	Status         BatchStatus               `json:"status"`
	Mode           DeepResearchMode          `json:"mode"`
	Name           string                    `json:"name,omitempty"`
	OutputFormats  []string                  `json:"output_formats,omitempty"`
	SearchParams   *DeepResearchSearchConfig `json:"search_params,omitempty"`
	CreatedAt      string                    `json:"created_at"`
	CompletedAt    string                    `json:"completed_at,omitempty"`
	Counts         BatchCounts               `json:"counts"`
	Cost           float64                   `json:"cost"`
	WebhookSecret  string                    `json:"webhook_secret,omitempty"`
	Metadata       map[string]interface{}    `json:"metadata,omitempty"`
}

type CreateBatchOptions struct {
	Name          string                    `json:"name,omitempty"`
	Mode          DeepResearchMode          `json:"mode,omitempty"`
	OutputFormats []string                  `json:"output_formats,omitempty"`
	Search        *DeepResearchSearchConfig `json:"search,omitempty"`
	WebhookURL    string                    `json:"webhook_url,omitempty"`
	Metadata      map[string]interface{}    `json:"metadata,omitempty"`
}

type BatchTaskInput struct {
	ID       string                 `json:"id,omitempty"`
	Query    string                 `json:"query"`
	Strategy string                 `json:"strategy,omitempty"`
	URLs     []string               `json:"urls,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type AddBatchTasksOptions struct {
	Tasks []BatchTaskInput `json:"tasks"`
}

type CreateBatchResponse struct {
	Success       bool             `json:"success"`
	Error         string           `json:"error,omitempty"`
	BatchID       string           `json:"batch_id,omitempty"`
	Status        BatchStatus      `json:"status,omitempty"`
	Mode          DeepResearchMode `json:"mode,omitempty"`
	Name          string           `json:"name,omitempty"`
	OutputFormats []string         `json:"output_formats,omitempty"`
	CreatedAt     string           `json:"created_at,omitempty"`
	Counts        *BatchCounts     `json:"counts,omitempty"`
	Cost          float64          `json:"cost,omitempty"`
	WebhookSecret string           `json:"webhook_secret,omitempty"`
}

type BatchStatusResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Batch   *Batch `json:"batch,omitempty"`
}

type BatchTaskCreated struct {
	TaskID         string `json:"task_id,omitempty"`
	DeepResearchID string `json:"deepresearch_id"`
	Status         string `json:"status"`
}

type AddBatchTasksResponse struct {
	Success bool               `json:"success"`
	Error   string             `json:"error,omitempty"`
	BatchID string             `json:"batch_id,omitempty"`
	Added   int                `json:"added,omitempty"`
	Tasks   []BatchTaskCreated `json:"tasks,omitempty"`
	Counts  *BatchCounts       `json:"counts,omitempty"`
}

type BatchTaskListItem struct {
	DeepResearchID string             `json:"deepresearch_id"`
	TaskID         string             `json:"task_id,omitempty"`
	Query          string             `json:"query"`
	Status         DeepResearchStatus `json:"status"`
	CreatedAt      string             `json:"created_at"`
	CompletedAt    string             `json:"completed_at,omitempty"`
}

type BatchPagination struct {
	Count   int    `json:"count"`
	LastKey string `json:"last_key,omitempty"`
	HasMore bool   `json:"has_more"`
}

type ListBatchTasksOptions struct {
	Status  DeepResearchStatus
	Limit   int
	LastKey string
}

type ListBatchTasksResponse struct {
	Success    bool                `json:"success"`
	Error      string              `json:"error,omitempty"`
	BatchID    string              `json:"batch_id,omitempty"`
	Tasks      []BatchTaskListItem `json:"tasks,omitempty"`
	Pagination *BatchPagination    `json:"pagination,omitempty"`
}

type CancelBatchResponse struct {
	Success        bool        `json:"success"`
	Error          string      `json:"error,omitempty"`
	BatchID        string      `json:"batch_id,omitempty"`
	Status         BatchStatus `json:"status,omitempty"`
	CancelledCount int         `json:"cancelled_count,omitempty"`
	Message        string      `json:"message,omitempty"`
}

type ListBatchesOptions struct {
	Limit int
}

type ListBatchesResponse struct {
	Success bool    `json:"success"`
	Error   string  `json:"error,omitempty"`
	Batches []Batch `json:"batches,omitempty"`
}

type BatchWaitOptions struct {
	PollInterval int64
	MaxWaitTime  int64
	OnProgress   func(batch *Batch)
}

type DatasourcePricing struct {
	CPM float64 `json:"cpm"`
}

type DatasourceCoverage struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type Datasource struct {
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	Category        DatasourceCategoryID   `json:"category"`
	Type            string                 `json:"type"`
	Modality        []string               `json:"modality"`
	Topics          []string               `json:"topics"`
	Languages       []string               `json:"languages,omitempty"`
	Source          string                 `json:"source,omitempty"`
	ExampleQueries  []string               `json:"example_queries"`
	Pricing         DatasourcePricing      `json:"pricing"`
	ResponseSchema  map[string]interface{} `json:"response_schema,omitempty"`
	UpdateFrequency string                 `json:"update_frequency,omitempty"`
	Size            int                    `json:"size,omitempty"`
	Coverage        *DatasourceCoverage    `json:"coverage,omitempty"`
}

type DatasourceCategory struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DatasetCount int    `json:"dataset_count"`
}

type DatasourcesListOptions struct {
	Category DatasourceCategoryID
}

type DatasourcesListResponse struct {
	Success     bool         `json:"success"`
	Error       string       `json:"error,omitempty"`
	Datasources []Datasource `json:"datasources,omitempty"`
}

type DatasourcesCategoriesResponse struct {
	Success    bool                 `json:"success"`
	Error      string               `json:"error,omitempty"`
	Categories []DatasourceCategory `json:"categories,omitempty"`
}

type GetAssetsOptions struct {
	Token string
}
