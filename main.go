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

// lines counts the number of lines in s.
func lines(s string) int {
	n := 0
	for _, b := range []byte(s) {
		if b == '\n' {
			n++
		}
	}
	if n == 0 {
		return 1
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
	msg := map[string]int{}
	switch req.Path {
	case "/lines":
		msg["lines"] = lines(req.Body)
	case "/words":
		msg["words"] = words(req.Body)
	default:
		msg["lines"] = lines(req.Body)
		msg["words"] = words(req.Body)
	}

	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(msg); err != nil {
		return respond(http.StatusInternalServerError,
			`{"error":"failed to encode response"}`)
	}

	return respond(http.StatusOK, b.String())
}

func main() {
	lambda.Start(handler)
}
