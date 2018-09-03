#!/bin/bash
GO111MODULE=on GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o wc || exit 1
zip wc.zip wc || exit 1
sam local start-api
