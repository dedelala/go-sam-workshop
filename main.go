package main // import "github.com/dedelala/go-sam-workshop/wc"

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rsp := events.APIGatewayProxyResponse{
		Body:       "blep",
		StatusCode: http.StatusOK,
	}
	return rsp, nil
}

func main() {
	lambda.Start(handler)
}
