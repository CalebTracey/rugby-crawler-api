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
	respC, err := client.GetLeaderboardDataWithResponse(context.Background(),
		openapi3.GetLeaderboardDataJSONRequestBody{
			CompId:   &compId,
			CompName: &compName,
			Date:     &date,
		})
	if err != nil {
		log.Fatalf("Couldn't get competition %s", err)
	}

	fmt.Printf("\tCompetition Id: %s\n", *respC.JSON201.Id)
	fmt.Printf("\tCompetition Name: %s\n", *respC.JSON201.Name)
	fmt.Printf("\tCompetition Teams: %v\n", *respC.JSON201.Teams)
	fmt.Printf("\tMessage: %v\n", *respC.JSON201.Message)
}
