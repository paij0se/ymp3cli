#!/bin/bash
echo Running server...
mkdir -p music
sleep 1
go run main.go 2>&1 >/dev/null &
deno run -A client.ts
