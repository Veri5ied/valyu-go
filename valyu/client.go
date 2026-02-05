package valyu

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Veri5ied/valyu-go/valyu/answer"
	"github.com/Veri5ied/valyu-go/valyu/batch"
	"github.com/Veri5ied/valyu-go/valyu/contents"
	"github.com/Veri5ied/valyu-go/valyu/datasources"
	"github.com/Veri5ied/valyu-go/valyu/deepresearch"
	"github.com/Veri5ied/valyu-go/valyu/internal/api"
	"github.com/Veri5ied/valyu-go/valyu/search"
)

const (
	DefaultBaseURL = "https://api.valyu.ai/v1"
	DefaultTimeout = 30 * time.Second
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client

	Search       *search.Service
	Answer       *answer.Service
	Contents     *contents.Service
	DeepResearch *deepresearch.Service
	Batch        *batch.Service
	Datasources  *datasources.Service
}

func New(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		apiKey = os.Getenv("VALYU_API_KEY")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("VALYU_API_KEY is not set")
	}

	c := &Client{
		baseURL: DefaultBaseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	apiClient := api.New(c.baseURL, c.apiKey, c.httpClient)

	c.Search = search.New(apiClient)
	c.Answer = answer.New(apiClient)
	c.Contents = contents.New(apiClient)
	c.DeepResearch = deepresearch.New(apiClient)
	c.Batch = batch.New(apiClient)
	c.Datasources = datasources.New(apiClient)

	return c, nil
}
