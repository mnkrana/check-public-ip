package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var response events.APIGatewayProxyResponse

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ip := request.RequestContext.Identity.SourceIP
	log.Printf("SourceIP %v =", ip)
	response = events.APIGatewayProxyResponse{Body: string(ip), StatusCode: 200}
	return response, nil
}
