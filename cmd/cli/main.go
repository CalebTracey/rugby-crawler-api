package main

import (
	"context"
	"fmt"
	"github.com/calebtracey/rugby-crawler-api/pkg/openapi3"
	"log"
)

func main() {
	client, err := openapi3.NewClientWithResponses("http://0.0.0.0:6080")
	if err != nil {
		log.Fatalf("Couldn't instantiate client: %s", err)
	}

	compId := "123"
	date := "2022/12/12"
	respC, err := client.CrawlCompWithResponse(context.Background(),
		openapi3.CrawlCompJSONRequestBody{
			CompetitionID: &compId,
			Date:          &date,
		})
	if err != nil {
		log.Fatalf("Couldn't get competition %s", err)
	}

	fmt.Printf("\tCompetition Name: %s\n", *respC.JSON201.Name)
	fmt.Printf("\tMessage: %v\n", *respC.JSON201.Message)
}
