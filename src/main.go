package main

import (
	"fmt"
	"log"
	"os"

	"github.com/paij0se/lmmp3"
	"github.com/paij0se/ymp3cli/src/cli"
	"github.com/paij0se/ymp3cli/src/client"
	"github.com/paij0se/ymp3cli/src/server"
	"github.com/phayes/freeport"
)

var (
	Version = "0.0.11"
)

func startServer() {
	port, err := freeport.GetFreePort()

	if err != nil {
		log.Panicln(err)

	}
	go client.StartClient(fmt.Sprintf(":%d", port))
	server.StartServer(fmt.Sprintf(":%d", port))
}

func main() {
	// create the folder if it doesn't exist
	if _, err := os.Stat("music"); os.IsNotExist(err) {
		os.Mkdir("music", 0777)

	}
	switch {
	// if the input is empty
	case len(os.Args) != 2:
		startServer()
	case os.Args[1] == "--h" || os.Args[1] == "--help":
		cli.HelpCommand()
	case os.Args[1] == "--v" || os.Args[1] == "--version":
		fmt.Println(Version)
	case os.Args[1][len(os.Args[1])-4:] == ".mp3":
		cli.PlaySongCli(os.Args[1])
	case os.Args[1][:4] == "http":
		lmmp3.DownloadAndConvert(os.Args[1])
	default:
		startServer()
	}
}
