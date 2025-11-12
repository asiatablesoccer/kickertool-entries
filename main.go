// Package main .
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/crispgm/kickertool-entries/pkg/kickertool"
)

// Response .
type Response struct {
	Message     string                        `json:"message"`
	Disciplines map[string][]kickertool.Entry `json:"disciplines"`
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	api := kickertool.New(os.Getenv("KICKERTOOL_ACCESS_TOKEN"))
	tournamentID, ok := request.QueryStringParameters["tournamentID"]
	if !ok || !strings.HasPrefix(tournamentID, "tio:") {
		message := "Invalid tournament ID"
		log.Println(message)
		return buildResponse(message, nil), nil
	}
	tournament, err := api.GetTournament(tournamentID)
	if err != nil {
		message := "Fetching tournament failed"
		log.Println(message)
		return buildResponse("Fetching tournament failed", nil), nil
	}

	allEntries := make(map[string][]kickertool.Entry)
	for _, disc := range tournament.Disciplines {
		entries, err := api.GetDisciplineEntries(tournamentID, disc.ID)
		if err != nil {
			message := fmt.Sprint("Fetching discipline", disc.Name, "failed")
			log.Println(message)
			return buildResponse(message, nil), nil
		}
		allEntries[disc.Name] = entries
		time.Sleep(20 * time.Millisecond)
	}
	return buildResponse("Success", allEntries), nil
}

func buildResponse(message string, data map[string][]kickertool.Entry) *events.APIGatewayProxyResponse {
	response := Response{
		Message:     message,
		Disciplines: data,
	}
	body, err := json.Marshal(response)
	if err != nil {
		response.Message = err.Error()
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
