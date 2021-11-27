package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetSongs() {
	resp, err := http.Get("http://127.0.0.1:8000/songs")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	// check if the body is empty
	if len(sb) == 0 {
		fmt.Println("No songs found, Download them😛")
	} else {
		// Print the body in blue
		fmt.Printf("\x1b[34m%s\x1b[0m\n", sb)
	}
}
