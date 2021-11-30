package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DownloadSpotify(url string, sound string) {
	postBody := map[string]string{"url": sound}
	jsonBody, err := json.Marshal(postBody)
	if err != nil {
		log.Fatalln(err)
	}
	responseBody := bytes.NewBuffer(jsonBody)
	resp, err := http.Post((url + "spotify"), "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// print the response of the server
	fmt.Println(string(body))
}
