package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Veri5ied/valyu-go/valyu"
)

func pretty(v interface{}) string {
	if v == nil {
		return "<nil>"
	}
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func main() {
	if os.Getenv("VALYU_API_KEY") == "" {
		log.Fatal("set VALYU_API_KEY env var")
	}

	client, err := valyu.NewClient("")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	// Search
	fmt.Println("=== Search ===")
	sr, err := client.Search(ctx, "What is retrieval-augmented generation?", &valyu.SearchOptions{MaxNumResults: 3})
	if err != nil {
		fmt.Println("Search error:", err)
	} else if !sr.Success {
		fmt.Println("Search API error:", sr.Error)
	} else {
		for i, r := range sr.Results {
			fmt.Printf("%d) %s — %s\n", i+1, r.Title, r.URL)
		}
	}

	// Answer (non-stream)
	fmt.Println("\n=== Answer ===")
	ans, err := client.Answer(ctx, "Explain RAG in AI, briefly.", nil)
	if err != nil {
		fmt.Println("Answer error:", err)
	} else if !ans.Success {
		fmt.Println("Answer API error:", ans.Error)
	} else {
		fmt.Println("Answer contents:", pretty(ans.Contents))
		fmt.Printf("Found %d sources\n", len(ans.SearchResults))
	}

	// AnswerStream (SSE)
	fmt.Println("=== AnswerStream ===")
	streamCh, err := client.AnswerStream(ctx, "What is RAG in AI?", &valyu.AnswerOptions{
		SearchType: valyu.SearchTypeAll,
	})
	if err != nil {
		fmt.Printf("AnswerStream error: %v\n", err)
	} else {
		var fullContent strings.Builder
		var streamResults []valyu.SearchResult

		for chunk := range streamCh {
			switch chunk.Type {
			case "error":
				fmt.Printf("Stream error: %s\n", chunk.Error)
			case "search_results":
				streamResults = append(streamResults, chunk.SearchResults...)
				fmt.Printf("Got %d search results\n", len(chunk.SearchResults))
			case "content":
				fullContent.WriteString(chunk.Content)
				if chunk.Content != "" {
					fmt.Print(chunk.Content)
				}
			case "metadata":
				fmt.Printf("\nStream metadata - TxID: %s\n", chunk.TxID)
			case "done":
				fmt.Println("\n[Stream completed]")
			}
		}

		fmt.Printf("\nFinal streamed content (%d chars)\n", fullContent.Len())
		fmt.Printf("Found %d sources via stream\n", len(streamResults))
	}

	// Datasources
	fmt.Println("\n=== Datasources ===")
	ds, err := client.Datasources.List(ctx, nil)
	if err != nil {
		fmt.Println("Datasources error:", err)
	} else if !ds.Success {
		fmt.Println("Datasources API error:", ds.Error)
	} else {
		fmt.Printf("Available datasources: %d\n", len(ds.Datasources))
	}
}
