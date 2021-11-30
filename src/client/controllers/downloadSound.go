package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func DownloadSound(url string, sound string) {
	postBody := map[string]string{"url": sound}

	jsonBody, err := json.Marshal(postBody)

	if err != nil {
		log.Fatalln(err)

	}

	responseBody := bytes.NewBuffer(jsonBody)

	resp, err := http.Post((url + "download"), "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)

	}

	defer resp.Body.Close()
}
