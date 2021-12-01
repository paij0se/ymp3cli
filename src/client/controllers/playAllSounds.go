package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PlaySoundAll(url string, sound string) {

	postBody := map[string]uint32{"nsong": 1}

	jsonBody, err := json.Marshal(postBody)

	if err != nil {
		log.Fatalln(err)

	}

	responseBody := bytes.NewBuffer(jsonBody)

	resp, err := http.Post(url+"y/all", "application/json", responseBody)

	if err != nil {
		log.Fatalln(err)

	}

	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(res))
}
