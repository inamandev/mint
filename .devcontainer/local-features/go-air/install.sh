#!/usr/bin/env bash
echo "installing air for Go hot reload"
go install github.com/air-verse/air@latest
echo "air installed successfully"
air -v
