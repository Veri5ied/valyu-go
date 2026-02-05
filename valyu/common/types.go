package common

type SearchType string

const (
	SearchTypeAll         SearchType = "all"
	SearchTypeWeb         SearchType = "web"
	SearchTypeProprietary SearchType = "proprietary"
	SearchTypeNews        SearchType = "news"
)

type CountryCode string

const (
	CountryCodeAll CountryCode = "ALL"
	CountryCodeAR  CountryCode = "AR"
	CountryCodeAU  CountryCode = "AU"
	CountryCodeAT  CountryCode = "AT"
	CountryCodeBE  CountryCode = "BE"
	CountryCodeBR  CountryCode = "BR"
	CountryCodeCA  CountryCode = "CA"
	CountryCodeCL  CountryCode = "CL"
	CountryCodeCN  CountryCode = "CN"
	CountryCodeDK  CountryCode = "DK"
	CountryCodeFI  CountryCode = "FI"
	CountryCodeFR  CountryCode = "FR"
	CountryCodeDE  CountryCode = "DE"
	CountryCodeHK  CountryCode = "HK"
	CountryCodeIN  CountryCode = "IN"
	CountryCodeID  CountryCode = "ID"
	CountryCodeIT  CountryCode = "IT"
	CountryCodeJP  CountryCode = "JP"
	CountryCodeKR  CountryCode = "KR"
	CountryCodeMY  CountryCode = "MY"
	CountryCodeMX  CountryCode = "MX"
	CountryCodeNL  CountryCode = "NL"
	CountryCodeNZ  CountryCode = "NZ"
	CountryCodeNO  CountryCode = "NO"
	CountryCodePL  CountryCode = "PL"
	CountryCodePT  CountryCode = "PT"
	CountryCodePH  CountryCode = "PH"
	CountryCodeRU  CountryCode = "RU"
	CountryCodeSA  CountryCode = "SA"
	CountryCodeZA  CountryCode = "ZA"
	CountryCodeES  CountryCode = "ES"
	CountryCodeSE  CountryCode = "SE"
	CountryCodeCH  CountryCode = "CH"
	CountryCodeTW  CountryCode = "TW"
	CountryCodeTR  CountryCode = "TR"
	CountryCodeGB  CountryCode = "GB"
	CountryCodeUS  CountryCode = "US"
)

type ResponseLength string

const (
	ResponseLengthShort  ResponseLength = "short"
	ResponseLengthMedium ResponseLength = "medium"
	ResponseLengthLarge  ResponseLength = "large"
	ResponseLengthMax    ResponseLength = "max"
)

type ExtractEffort string

const (
	ExtractEffortNormal ExtractEffort = "normal"
	ExtractEffortHigh   ExtractEffort = "high"
	ExtractEffortAuto   ExtractEffort = "auto"
)

type DeepResearchMode string

const (
	DeepResearchModeFast     DeepResearchMode = "fast"
	DeepResearchModeStandard DeepResearchMode = "standard"
	DeepResearchModeHeavy    DeepResearchMode = "heavy"
)

type DeepResearchStatus string

const (
	DeepResearchStatusQueued    DeepResearchStatus = "queued"
	DeepResearchStatusRunning   DeepResearchStatus = "running"
	DeepResearchStatusCompleted DeepResearchStatus = "completed"
	DeepResearchStatusFailed    DeepResearchStatus = "failed"
	DeepResearchStatusCancelled DeepResearchStatus = "cancelled"
)

type BatchStatus string

const (
	BatchStatusOpen                BatchStatus = "open"
	BatchStatusProcessing          BatchStatus = "processing"
	BatchStatusCompleted           BatchStatus = "completed"
	BatchStatusCompletedWithErrors BatchStatus = "completed_with_errors"
	BatchStatusCancelled           BatchStatus = "cancelled"
)

type DatasourceCategoryID string

const (
	DatasourceCategoryResearch       DatasourceCategoryID = "research"
	DatasourceCategoryHealthcare     DatasourceCategoryID = "healthcare"
	DatasourceCategoryPatents        DatasourceCategoryID = "patents"
	DatasourceCategoryMarkets        DatasourceCategoryID = "markets"
	DatasourceCategoryCompany        DatasourceCategoryID = "company"
	DatasourceCategoryEconomic       DatasourceCategoryID = "economic"
	DatasourceCategoryPredictions    DatasourceCategoryID = "predictions"
	DatasourceCategoryTransportation DatasourceCategoryID = "transportation"
	DatasourceCategoryLegal          DatasourceCategoryID = "legal"
	DatasourceCategoryPolitics       DatasourceCategoryID = "politics"
)
