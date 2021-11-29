package main

import (
	"os"
	"time"

	client "github.com/paij0se/ymp3cli/client"
	router "github.com/paij0se/ymp3cli/src/router"
)

func main() {
	// create the folder music if doesent exists
	if _, err := os.Stat("music"); os.IsNotExist(err) {
		os.Mkdir("music", 0777)
	}
	// execute the client in a go routine
	go client.Clientmain()
	// execute the router, then wait 2 seconds 🍦
	router.SetUpRoutes()
	time.Sleep(2 * time.Second)

}
