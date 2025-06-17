#!/bin/bash
set -e
go fmt main.go
CGO_ENABLED=0 GOOS=linux go build .
strip go-helloword-http
