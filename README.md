<p align="center">
  <img src="https://valyu.ai/logo.svg" alt="Valyu" width="120">
</p>

<h1 align="center">Valyu Go SDK</h1>

<p align="center">
  <a href="https://pkg.go.dev/github.com/Veri5ied/valyu-go"><img src="https://pkg.go.dev/badge/github.com/Veri5ied/valyu-go.svg" alt="Go Reference"></a>
  <a href="https://github.com/Veri5ied/valyu-go/blob/main/LICENSE"><img src="https://img.shields.io/github/license/Veri5ied/valyu-go" alt="License"></a>
  <a href="https://github.com/Veri5ied/valyu-go/releases"><img src="https://img.shields.io/github/v/release/Veri5ied/valyu-go" alt="Release"></a>
</p>

<p align="center">
  The official Go SDK for <a href="https://valyu.ai">Valyu's</a> DeepSearch API.<br>
  Access academic papers, real-time web content, structured data, and AI-powered research.
</p>

---

## Installation

```bash
go get github.com/Veri5ied/valyu-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/Veri5ied/valyu-go/valyu"
    "github.com/Veri5ied/valyu-go/valyu/search"
)

func main() {
    client, err := valyu.New("") // Uses VALYU_API_KEY env var
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    resp, _ := client.Search.Search(ctx, "machine learning transformers", nil)
    for _, r := range resp.Results {
        fmt.Println(r.Title)
    }
}
```

## Authentication

```bash
export VALYU_API_KEY="your-api-key"
```

Or pass directly:

```go
client, err := valyu.New("your-api-key")
```

## APIs

### Search

```go
import "github.com/Veri5ied/valyu-go/valyu/search"
import "github.com/Veri5ied/valyu-go/valyu/common"

resp, _ := client.Search.Search(ctx, "query", &search.Options{
    SearchType:    common.SearchTypeProprietary,
    MaxNumResults: 20,
    IncludedSources: []string{"valyu/valyu-arxiv"},
})
```

### Contents

```go
import "github.com/Veri5ied/valyu-go/valyu/contents"
import "github.com/Veri5ied/valyu-go/valyu/common"

resp, _ := client.Contents.Get(ctx, []string{"https://example.com"}, &contents.Options{
    Summary:        true,
    ResponseLength: common.ResponseLengthMax,
})
```

### Answer

```go
import "github.com/Veri5ied/valyu-go/valyu/answer"
import "github.com/Veri5ied/valyu-go/valyu/common"

resp, _ := client.Answer.Answer(ctx, "What are transformers?", &answer.Options{
    SearchType: common.SearchTypeAll,
    FastMode:   true,
})
```

### DeepResearch

```go
import "github.com/Veri5ied/valyu-go/valyu/deepresearch"
import "github.com/Veri5ied/valyu-go/valyu/common"

task, _ := client.DeepResearch.Create(ctx, &deepresearch.CreateOptions{
    Query: "AI safety research summary",
    Mode:  common.DeepResearchModeFast,
})

status, _ := client.DeepResearch.Get(ctx, task.DeepResearchID)
```

### Batch

```go
import "github.com/Veri5ied/valyu-go/valyu/batch"
import "github.com/Veri5ied/valyu-go/valyu/common"

newBatch, _ := client.Batch.Create(ctx, &batch.CreateOptions{
    Name: "Research Batch",
    Mode: common.DeepResearchModeStandard,
})

status, _ := client.Batch.Get(ctx, newBatch.BatchID)
```

### Datasources

```go
sources, _ := client.Datasources.List(ctx)
```

## Configuration

```go
client, err := valyu.New("api-key",
    valyu.WithBaseURL("https://custom.api.com"),
    valyu.WithTimeout(60 * time.Second),
    valyu.WithHTTPClient(customClient),
)
```

## Error Handling

```go
resp, err := client.Search.Search(ctx, "query", nil)
if err != nil {
    log.Fatal("Network error:", err)
}
if !resp.Success {
    log.Fatal("API error:", resp.Error)
}
```

## License

MIT License - see [LICENSE](LICENSE) for details.
