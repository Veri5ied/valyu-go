# Academic Research API

An example API server demonstrating the Valyu Go SDK's capabilities for academic research. This example showcases how to use the Search and Answer APIs to build an intelligent research assistant.

## Features

- **Paper Search**: Search academic papers from arXiv and other scholarly sources
- **Topic Summarization**: Get AI-powered summaries of research topics with sources
- **Paper Analysis**: Deep analysis of specific research papers

## Prerequisites

- Go 1.21 or higher
- Valyu API key from [valyu.ai](https://valyu.ai)

## Installation

1. Clone the repository and navigate to this example:

```bash
cd examples/academic-research-api
```

2. Set your API key:

```bash
export VALYU_API_KEY="your-api-key-here"
```

3. Run the server:

```bash
go run main.go
```

The server will start on port `8081` by default (configurable via `PORT` environment variable).

## API Endpoints

### 1. Search Papers

Search for academic papers on a specific topic.

**Endpoint**: `POST /search`

**Request**:

```json
{
  "query": "transformer neural networks",
  "max_results": 10
}
```

**Response**:

```json
{
  "query": "transformer neural networks",
  "results": [
    {
      "title": "Attention Is All You Need",
      "url": "https://arxiv.org/abs/1706.03762",
      "abstract": "The dominant sequence transduction models...",
      "source": "valyu/valyu-arxiv"
    }
  ],
  "count": 10
}
```

**Example**:

```bash
curl -X POST http://localhost:8081/search \
  -H "Content-Type: application/json" \
  -d '{"query": "quantum computing", "max_results": 5}'
```

### 2. Summarize Research Topic

Get a comprehensive AI-powered summary of recent research on a topic.

**Endpoint**: `POST /summarize`

**Request**:

```json
{
  "topic": "CRISPR gene editing applications"
}
```

**Response**:

```json
{
  "topic": "CRISPR gene editing applications",
  "summary": "...",
  "sources": [...]
}
```

**Example**:

```bash
curl -X POST http://localhost:8081/summarize \
  -H "Content-Type: application/json" \
  -d '{"topic": "machine learning interpretability"}'
```

### 3. Analyze Paper

Get a detailed analysis of a specific research paper.

**Endpoint**: `POST /papers`

**Request**:

```json
{
  "url": "https://arxiv.org/abs/1706.03762"
}
```

**Response**:

```json
{
  "url": "https://arxiv.org/abs/1706.03762",
  "analysis": "...",
  "sources": [...]
}
```

**Example**:

```bash
curl -X POST http://localhost:8081/papers \
  -H "Content-Type: application/json" \
  -d '{"url": "https://arxiv.org/abs/2103.00020"}'
```

### 4. Health Check

Check if the server is running.

**Endpoint**: `GET /health`

**Response**:

```json
{
  "status": "healthy"
}
```

## Authentication

You can provide your API key in two ways:

1. **Environment variable** (recommended):

```bash
export VALYU_API_KEY="your-api-key"
```

2. **Request header**:

```bash
curl -X POST http://localhost:8081/search \
  -H "X-API-Key: your-api-key" \
  -H "Content-Type: application/json" \
  -d '{"query": "climate change"}'
```

## Use Cases

- **Literature Review**: Quickly search and summarize papers for a research topic
- **Paper Discovery**: Find relevant academic papers from arXiv
- **Research Analysis**: Get AI-powered insights on specific papers
- **Academic Monitoring**: Track latest research in specific fields

## SDK Features Demonstrated

- `Search.Search()` - Search academic papers with filters
- `Answer.Answer()` - AI-powered question answering with sources
- Custom search options (SearchType, MaxNumResults, IncludedSources)
- Error handling and response validation
- Context management and timeouts

## Configuration

- **Port**: Set via `PORT` environment variable (default: 8081)
- **API Key**: Set via `VALYU_API_KEY` environment variable or `X-API-Key` header
- **Timeout**: Search requests timeout after 30s, Answer requests after 60s
