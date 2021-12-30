package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//./main -p 8080
func makePostRequest(content *strings.Reader, service string) (err error) {
	r, err := http.Post("http://localhost:8080/"+service, "application/json", content)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))
	return
}
func requestSpotify() (err error) {
	content := strings.NewReader(`{"Url":"https://open.spotify.com/track/7JRN5xOUIrnI4crUMOt6X4?si=4caf7c66e37d4b77"}`)
	if err = makePostRequest(content, "spotify"); err != nil {
		fmt.Println(err)
	}
	content = strings.NewReader(`{"Url":""}`)
	if err = makePostRequest(content, "spotify"); err != nil {
		fmt.Println(err)
	}
	content = strings.NewReader(`{"Url":"https://www.youtube.com/watch?v=ACZURx8hYJQ"}`)
	if err = makePostRequest(content, "spotify"); err != nil {
		fmt.Println(err)
	}
	return
}

func requestYoutube() (err error) {
	content := strings.NewReader(`{"Url": "https://www.youtube.com/watch?v=ACZURx8hYJQ"}`)
	if err = makePostRequest(content, "youtube"); err != nil {
		fmt.Println(err)
	}
	content = strings.NewReader(`{"Url":""}`)
	if err = makePostRequest(content, "youtube"); err != nil {
		fmt.Println(err)
	}
	content = strings.NewReader(`{"Url":"https://open.spotify.com/track/7JRN5xOUIrnI4crUMOt6X4?si=4caf7c66e37d4b77"}`)
	if err = makePostRequest(content, "youtube"); err != nil {
		fmt.Println(err)
	}

	return
}

func main() {
	log.Println(requestSpotify())
	log.Println(requestYoutube())

}
