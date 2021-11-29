package main

import (
	"os"

	client "github.com/paij0se/ymp3cli/client"
	core "github.com/paij0se/ymp3cli/src/core"
)

func main() {
	// create the folder music in ./local/music if doesent exists
	if _, err := os.Stat("music"); os.IsNotExist(err) {
		os.Mkdir("music", 0777)
	}
	// execute the client in a go routine
	go client.Clientmain()
	// execute the server
	core.Core()

}
