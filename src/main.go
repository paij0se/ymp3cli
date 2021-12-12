package main

import (
	"fmt"
	"log"
	"os"

	"github.com/paij0se/ymp3cli/src/client"
	"github.com/paij0se/ymp3cli/src/server"
	"github.com/phayes/freeport"
)

func main() {
	switch {
	case len(os.Args) != 2:
		if _, err := os.Stat("music"); os.IsNotExist(err) {
			os.Mkdir("music", 0777)

		}

		port, err := freeport.GetFreePort()

		if err != nil {
			log.Panicln(err)

		}
		go client.StartClient(fmt.Sprintf(":%d", port))
		server.StartServer(fmt.Sprintf(":%d", port))
	case os.Args[1] == "--h":
		fmt.Println(`   

          --h: Display the help command
          --v: Display the versions

			 `)
	case os.Args[1] == "--v":
		fmt.Println("v0.0.9")
	default:
		if _, err := os.Stat("music"); os.IsNotExist(err) {
			os.Mkdir("music", 0777)

		}

		port, err := freeport.GetFreePort()

		if err != nil {
			log.Panicln(err)

		}
		go client.StartClient(fmt.Sprintf(":%d", port))
		server.StartServer(fmt.Sprintf(":%d", port))
	}
}
