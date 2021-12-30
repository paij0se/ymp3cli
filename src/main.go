package main

import (
	"flag"
	"fmt"
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
	version     = "0.0.14"
	help        bool
	showVersion bool
	url         string
	port        int
	song        string
)

func startServer() (err error) {

	if port == 0 {
		port, err = freeport.GetFreePort()
		if err != nil {
			fmt.Println("Error getting free port")
			return
		}
	}

	go rpc.DefaultRpc(port)
	go client.StartClient(fmt.Sprintf(":%d", port))
	server.StartServer(fmt.Sprintf(":%d", port))
	return
}
func init() {
	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&help, "h", false, "show help")

	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.BoolVar(&showVersion, "v", false, "show version")

	flag.StringVar(&url, "d", "", "download a song from youtube ")
	flag.StringVar(&url, "download", "", "download a song from youtube ")

	flag.StringVar(&song, "p", "", "play a single song")
	flag.StringVar(&song, "play", "", "play a single song")

	flag.IntVar(&port, "port", 0, "port to run the server on")
	flag.Parse()
}

func main() {
	// create the folder if it doesn't exist
	if _, err := os.Stat("music"); os.IsNotExist(err) {
		os.Mkdir("music", 0777)

	}
	if url != "" {
		lmmp3.DownloadAndConvert(os.Args[1])
		if runtime.GOOS == "windows" {
			del := exec.Command("del", "*.mpeg")
			if del.Run() != nil {
				fmt.Println("Error deleting the mpeg files")
			}
		}

	} else if song != "" {
		cli.PlaySongCli(song)
	} else if help {
		cli.HelpCommand()
	} else if showVersion {
		fmt.Println(version)

	} else {
		startServer()
	}
}
