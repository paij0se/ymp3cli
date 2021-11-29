package cliRequests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ChooseSong(song int) {
	//Encode the data
	// the post body is just a int value
	postBody := map[string]int{"nsong": song}
	// convert podBody to json
	jsonBody, err := json.Marshal(postBody)
	if err != nil {
		log.Fatalln(err)
	}
	responseBody := bytes.NewBuffer(jsonBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("http://127.0.0.1:8000/y", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	// print sb in green color
	fmt.Printf("\x1b[32m%s\x1b[0m\n", sb)
}
