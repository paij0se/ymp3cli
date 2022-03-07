package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func PlaySound(url string, sound string) {
	number, _ := strconv.Atoi(sound)

	postBody := map[string]uint32{"nsong": uint32(number - 1)}

	jsonBody, err := json.Marshal(postBody)

	if err != nil {
		log.Fatalln(err)

	}

	responseBody := bytes.NewBuffer(jsonBody)

	resp, err := http.Post(url+"y", "application/json", responseBody)

	if err != nil {
		log.Fatalln(err)
		log.Println("Error: ", err)

	}

	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(res))
}
