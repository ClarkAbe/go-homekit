#!/bin/bash
echo "Build Linux mipsle..."
CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -ldflags "-s -w" -o ./bin/homekit.linux_mipsle.bin main.go 
echo "Build Linux amd64..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/homekit.linux_amd64.bin main.go 
echo "Build Windows amd64..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/homekit.windows_mipsle.exe main.go 
echo "Build Darwin..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/homekit.darwin_amd64.bin main.go 
