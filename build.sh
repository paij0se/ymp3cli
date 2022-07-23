#!/bin/bash

GOOS=linux go build -ldflags "-s -w" -o ymp3cli src/main.go
zip ymp3cli-linux.zip ymp3cli

GOOS=windows go build -ldflags "-s -w" -o ymp3cli.exe src/main.go
zip ymp3cli-windows.zip ymp3cli.exe

rm ymp3cli.exe ymp3cli


