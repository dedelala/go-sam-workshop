package main // import "github.com/dedelala/go-sam-workshop"

import (
	"bytes"
	"encoding/json"
	"net/http"
	"unicode"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// words counts the number of continuous runs of non-space runes in s.
func words(s string) int {
	b, n := true, 0
	for _, r := range []rune(s) {
		switch {
		case unicode.IsSpace(r):
			b = true
		case b:
			b = false
			n++
		}
	}
	return n
}

// respond returns a handler response in a slightly less verbose way.
func respond(code int, msg string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       msg,
		StatusCode: code,
	}, nil
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var b bytes.Buffer

	msg := map[string]int{"words": words(req.Body)}

	if err := json.NewEncoder(&b).Encode(msg); err != nil {
		return respond(http.StatusInternalServerError,
			`{"error":"failed to encode response"}`)
	}

	return respond(http.StatusOK, b.String())
}

func main() {
	lambda.Start(handler)
}
