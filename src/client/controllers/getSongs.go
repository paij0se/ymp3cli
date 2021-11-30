package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetSongs(url string) []string {
	resp, err := http.Get(url + "songs")

	if err != nil {
		log.Fatalln(err)

	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)

	}

	strBody := string(body)

	if len(strBody) == 0 {
		fmt.Println("🐶No songs found, Download them👻")

		return []string{}
	}

	songs := strings.Split(strBody, "\n")

	return songs[0 : len(songs)-1]
}
