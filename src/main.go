package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/paij0se/lmmp3"
	"github.com/paij0se/ymp3cli/src/cli"
	"github.com/paij0se/ymp3cli/src/client"
	"github.com/paij0se/ymp3cli/src/rpc"
	"github.com/paij0se/ymp3cli/src/server"
	"github.com/phayes/freeport"
)

var (
	version = "0.0.14"
)

func startServer() {
	port, err := freeport.GetFreePort()

	if err != nil {
		log.Panicln(err)

	}
	go rpc.DefaultRpc(port)
	go client.StartClient(fmt.Sprintf(":%d", port))
	server.StartServer(fmt.Sprintf(":%d", port))
}

func main() {
	// create the folder if it doesn't exist
	if _, err := os.Stat("music"); os.IsNotExist(err) {
		os.Mkdir("music", 0777)

	}
	switch {
	case len(os.Args) != 2: // if the input is empty
		startServer()
	case os.Args[1] == "--h" || os.Args[1] == "--help":
		cli.HelpCommand()
	case os.Args[1] == "--v" || os.Args[1] == "--version":
		fmt.Println(version)
	case os.Args[1][len(os.Args[1])-4:] == ".mp3":
		cli.PlaySongCli(os.Args[1])
	case os.Args[1][:4] == "http": // if the input is a youtube Url
		lmmp3.DownloadAndConvert(os.Args[1])
		if runtime.GOOS == "windows" {
			del := exec.Command("cmd", "/C", "del", "*.mpeg")
			if del.Run() != nil {
				fmt.Println("Error deleting the mpeg files")
			}
		}
	default:
		startServer()
	}
}
