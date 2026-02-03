package valyu

import (
	"context"
	"fmt"
	"regexp"
	"strings"
)

type searchRequest struct {
	Query              string   `json:"query"`
	SearchType         string   `json:"search_type,omitempty"`
	MaxNumResults      int      `json:"max_num_results,omitempty"`
	MaxPrice           float64  `json:"max_price,omitempty"`
	IsToolCall         *bool    `json:"is_tool_call,omitempty"`
	RelevanceThreshold float64  `json:"relevance_threshold,omitempty"`
	IncludedSources    []string `json:"included_sources,omitempty"`
	ExcludeSources     []string `json:"exclude_sources,omitempty"`
	Category           string   `json:"category,omitempty"`
	StartDate          string   `json:"start_date,omitempty"`
	EndDate            string   `json:"end_date,omitempty"`
	CountryCode        string   `json:"country_code,omitempty"`
	ResponseLength     string   `json:"response_length,omitempty"`
	FastMode           bool     `json:"fast_mode,omitempty"`
	URLOnly            bool     `json:"url_only,omitempty"`
}

func validateDateFormat(date string) bool {
	if date == "" {
		return true
	}
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return dateRegex.MatchString(date)
}

func validateSource(source string) bool {
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		return true
	}

	if strings.Contains(source, ".") {
		domainRegex := regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+(/.+)?$`)
		if domainRegex.MatchString(source) {
			return true
		}
	}

	parts := strings.Split(source, "/")
	if len(parts) == 2 {
		providerRegex := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
		if providerRegex.MatchString(parts[0]) && providerRegex.MatchString(parts[1]) {
			return true
		}
	}

	return false
}

func (c *Client) Search(ctx context.Context, query string, opts *SearchOptions) (*SearchResponse, error) {
	if opts == nil {
		opts = &SearchOptions{}
	}

	searchType := opts.SearchType
	if searchType == "" {
		searchType = SearchTypeAll
	}

	maxNumResults := opts.MaxNumResults
	if maxNumResults == 0 {
		maxNumResults = 10
	}

	relevanceThreshold := opts.RelevanceThreshold
	if relevanceThreshold == 0 {
		relevanceThreshold = 0.5
	}

	isToolCall := opts.IsToolCall
	if isToolCall == nil {
		defaultToolCall := true
		isToolCall = &defaultToolCall
	}

	validSearchTypes := map[SearchType]bool{
		SearchTypeAll:         true,
		SearchTypeWeb:         true,
		SearchTypeProprietary: true,
		SearchTypeNews:        true,
	}
	if !validSearchTypes[searchType] {
		return &SearchResponse{
			Success: false,
			Error:   "Invalid searchType provided. Must be one of: all, web, proprietary, news",
			Query:   query,
		}, nil
	}

	if !validateDateFormat(opts.StartDate) {
		return &SearchResponse{
			Success: false,
			Error:   "Invalid startDate format. Must be YYYY-MM-DD",
			Query:   query,
		}, nil
	}

	if !validateDateFormat(opts.EndDate) {
		return &SearchResponse{
			Success: false,
			Error:   "Invalid endDate format. Must be YYYY-MM-DD",
			Query:   query,
		}, nil
	}

	if maxNumResults < 1 || maxNumResults > 100 {
		return &SearchResponse{
			Success: false,
			Error:   "maxNumResults must be between 1 and 100",
			Query:   query,
		}, nil
	}

	for _, source := range opts.IncludedSources {
		if !validateSource(source) {
			return &SearchResponse{
				Success: false,
				Error:   fmt.Sprintf("Invalid includedSources format. Invalid source: %s", source),
				Query:   query,
			}, nil
		}
	}

	for _, source := range opts.ExcludeSources {
		if !validateSource(source) {
			return &SearchResponse{
				Success: false,
				Error:   fmt.Sprintf("Invalid excludeSources format. Invalid source: %s", source),
				Query:   query,
			}, nil
		}
	}

	req := searchRequest{
		Query:              query,
		SearchType:         string(searchType),
		MaxNumResults:      maxNumResults,
		IsToolCall:         isToolCall,
		RelevanceThreshold: relevanceThreshold,
	}

	if opts.MaxPrice > 0 {
		req.MaxPrice = opts.MaxPrice
	}
	if len(opts.IncludedSources) > 0 {
		req.IncludedSources = opts.IncludedSources
	}
	if len(opts.ExcludeSources) > 0 {
		req.ExcludeSources = opts.ExcludeSources
	}
	if opts.Category != "" {
		req.Category = opts.Category
	}
	if opts.StartDate != "" {
		req.StartDate = opts.StartDate
	}
	if opts.EndDate != "" {
		req.EndDate = opts.EndDate
	}
	if opts.CountryCode != "" {
		req.CountryCode = string(opts.CountryCode)
	}
	if opts.ResponseLength != "" {
		req.ResponseLength = string(opts.ResponseLength)
	}
	if opts.FastMode {
		req.FastMode = opts.FastMode
	}
	if opts.URLOnly {
		req.URLOnly = opts.URLOnly
	}

	var response SearchResponse
	if err := c.doRequest(ctx, "POST", "/deepsearch", req, &response); err != nil {
		return &SearchResponse{
			Success: false,
			Error:   err.Error(),
			Query:   query,
		}, nil
	}

	return &response, nil
}
