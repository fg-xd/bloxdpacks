#!/usr/bin/bash
GOOS=windows GOARCH=amd64 go build -o ./dist/win/bloxdpacks.exe ./bloxdpacks.go
GOOS=darwin GOARCH=arm64 go build -o ./dist/mac/bloxdpacks ./bloxdpacks.go
GOOS=linux GOARCH=amd64 go build -o ./dist/linux/bloxdpacks ./bloxdpacks.go
# Build Script
