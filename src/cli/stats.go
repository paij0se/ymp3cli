package cli

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/user"
)

/*
Disclaimer:
I don't  track you.
Its just a simple stats for the ammount of users.
Its open source!

you can remove it if you want.
*/

func Stats() {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	username := user.Username
	postBody, _ := json.Marshal(map[string]string{
		"Name": username,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://ymp3cli-api.herokuapp.com/user", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
}
