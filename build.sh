#!/bin/bash
CGO_ENABLED=0 GOOS=linux go build .
strip go-helloword-http
