package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/paij0se/ymp3cli/src/server/tools"
)

//./main -p 8080
func requestSpotify() error {
	b, _ := json.Marshal(tools.Url{Url: "https://open.spotify.com/track/7JRN5xOUIrnI4crUMOt6X4?si=4caf7c66e37d4b77"})
	r, err := http.Post("http://localhost:8080/spotify", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer r.Body.Close()
	b, _ = ioutil.ReadAll(r.Body)
	fmt.Println(string(b))
	return nil
}

func requestYoutube() error {

	return nil
}
