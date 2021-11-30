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

func DeleteSound(url string, sound string) {
	type Payload struct {
		Delete uint32 `json:"delete"`
	}

	number, _ := strconv.Atoi(sound)

	data := Payload{
		Delete: uint32(number - 1),
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)

	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("DELETE", url+"delete", body)

	if err != nil {
		log.Println(err)

	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)

	}

	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// print the response of the server
	fmt.Println(string(res))

}
