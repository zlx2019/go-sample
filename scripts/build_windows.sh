#!/bin/sh

# Windows
export CGO_ENABLED=0
export GOOS=windows
export GOARCH=amd64
go build -o ../build/ ../cmd/app/app.go