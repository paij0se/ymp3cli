#!/bin/bash
git clone https://github.com/paij0se/ymp3cli
cd ymp3cli
echo downloading dependencies
pip3 install spotdl
cd src ; go get . ; cd ..
go mod download github.com/ELPanaJose/pairat
go mod download github.com/manifoldco/promptui
go get github.com/manifoldco/promptui@v0.9.0
echo dependencies installed, now run <ymp3cli>
go build -o ymp3cli src/main.go
sudo mv ymp3cli /bin
