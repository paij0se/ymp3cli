#!/bin/bash
echo downloading dependencies
cd src ; go get . ; cd ..
go mod download github.com/ELPanaJose/pairat
go mod download github.com/manifoldco/promptui
go get github.com/manifoldco/promptui@v0.9.0
echo dependencies installed, now run ./ymp3cli to start
