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
)

func main() {
    client, err := valyu.NewClient("") // Uses VALYU_API_KEY env var
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    resp, _ := client.Search(ctx, "machine learning transformers", nil)
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
client, err := valyu.NewClient("your-api-key")
```

## APIs

### Search

```go
resp, _ := client.Search(ctx, "query", &valyu.SearchOptions{
    SearchType:    valyu.SearchTypeProprietary,
    MaxNumResults: 20,
    IncludedSources: []string{"valyu/valyu-arxiv"},
})
```

### Contents

```go
resp, _ := client.Contents(ctx, []string{"https://example.com"}, &valyu.ContentsOptions{
    Summary:        true,
    ResponseLength: valyu.ResponseLengthMax,
})
```

### Answer

```go
resp, _ := client.Answer(ctx, "What are transformers?", &valyu.AnswerOptions{
    SearchType: valyu.SearchTypeAll,
    FastMode:   true,
})
```

### DeepResearch

```go
task, _ := client.DeepResearch.Create(ctx, &valyu.DeepResearchCreateOptions{
    Query: "AI safety research summary",
    Mode:  valyu.DeepResearchModeFast,
})

result, _ := client.DeepResearch.Wait(ctx, task.DeepResearchID, nil)
fmt.Println(result.Output)
```

### Batch

```go
batch, _ := client.Batch.Create(ctx, &valyu.CreateBatchOptions{
    Name: "Research Batch",
    Mode: valyu.DeepResearchModeStandard,
})

client.Batch.AddTasks(ctx, batch.BatchID, &valyu.AddBatchTasksOptions{
    Tasks: []valyu.BatchTaskInput{
        {Query: "Topic 1"},
        {Query: "Topic 2"},
    },
})

result, _ := client.Batch.WaitForCompletion(ctx, batch.BatchID, nil)
```

### Datasources

```go
sources, _ := client.Datasources.List(ctx, nil)
categories, _ := client.Datasources.Categories(ctx)
```

## Configuration

```go
client, err := valyu.NewClient("api-key",
    valyu.WithBaseURL("https://custom.api.com"),
    valyu.WithTimeout(60 * time.Second),
    valyu.WithHTTPClient(customClient),
)
```

## Error Handling

```go
resp, err := client.Search(ctx, "query", nil)
if err != nil {
    log.Fatal("Network error:", err)
}
if !resp.Success {
    log.Fatal("API error:", resp.Error)
}
```

## License

MIT License - see [LICENSE](LICENSE) for details.
