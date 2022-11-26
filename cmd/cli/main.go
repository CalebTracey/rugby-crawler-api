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
	compName := "test comp"
	date := "2022/12/12"
	respA, err := client.GetLeaderboardDataWithResponse(context.Background(),
		openapi3.GetLeaderboardDataJSONRequestBody{
			CompId:   &compId,
			CompName: &compName,
			Date:     &date,
		})
	if err != nil {
		log.Fatalf("Couldn't get competition %s", err)
	}

	fmt.Printf("\tLeaderboard Data: %v\n", *respA.JSON201.LeaderboardData)
	fmt.Printf("\tMessage: %v\n", *respA.JSON201.Message)

	respB, err := client.GetAllLeaderboardDataWithResponse(context.Background())
	if err != nil {
		log.Fatalf("Couldn't get competition %s", err)
	}

	fmt.Printf("\tLeaderboard Data List: %v\n", *respB.JSON201.LeaderboardDataList)
	fmt.Printf("\tMessage: %v\n", *respB.JSON201.Message)
}
