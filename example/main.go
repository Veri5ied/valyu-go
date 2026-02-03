package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Veri5ied/valyu-go/valyu"
)

func main() {
	client, err := valyu.NewClient("")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	fmt.Println("=== Search ===")
	searchResp, _ := client.Search(ctx, "What is machine learning?", nil)
	if searchResp.Success {
		fmt.Printf("Found %d results\n", len(searchResp.Results))
	}

	fmt.Println("\n=== Contents ===")
	contentsResp, _ := client.Contents(ctx, []string{"https://docs.valyu.ai"}, nil)
	if contentsResp.Success {
		fmt.Printf("Processed %d URLs\n", contentsResp.URLsProcessed)
	}

	fmt.Println("\n=== Answer ===")
	answerResp, _ := client.Answer(ctx, "What is RAG?", nil)
	if answerResp.Success {
		fmt.Printf("Answer received with %d sources\n", len(answerResp.SearchResults))
	}

	fmt.Println("\n=== DeepResearch ===")
	task, _ := client.DeepResearch.Create(ctx, &valyu.DeepResearchCreateOptions{
		Query: "Latest AI developments",
		Mode:  valyu.DeepResearchModeFast,
	})
	if task.Success {
		fmt.Printf("Task created: %s\n", task.DeepResearchID)
	}

	fmt.Println("\n=== Datasources ===")
	sources, _ := client.Datasources.List(ctx, nil)
	if sources.Success {
		fmt.Printf("Available: %d datasources\n", len(sources.Datasources))
	}
}
