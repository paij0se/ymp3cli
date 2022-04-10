package rpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hugolgst/rich-go/client"
)

type Song struct {
	Path  string
	By    string
	Title string
	Img   string
	File  string
	Url   string
}

func Rpc(port int) {
	now := time.Now()
	for {
		time.Sleep(time.Second * 5)
		url := "http://localhost:" + fmt.Sprintf("%d", port) + "/currentSong"
		resp, error := http.Get(url)
		if error != nil {
			log.Println(error)
		}
		body, error := ioutil.ReadAll(resp.Body)
		if error != nil {
			log.Println(error)
		}
		var song Song
		error = json.Unmarshal(body, &song)
		if error != nil {
			log.Println(error)
		}
		// {"by":"","file":"music/Just Water.mp3","img":"https://img.youtube.com/vi//hqdefault.jpg","path":"/home/jose/ymp3cli/music/music/Just Water.mp3","title":""}
		if song.By == "" && song.Title == "" && song.Img == "https://img.youtube.com/vi//hqdefault.jpg" && song.Url == "" {
			song.By = "unknown"
			song.Title = song.File
			song.Img = "https://cdn.discordapp.com/emojis/822805787771928597.webp"
			song.Url = "https://ymp3cli.tk"
		}
		id := client.Login("851297648111517697")
		if error != nil {
			fmt.Println("No discord detected")
		}
		client.SetActivity(client.Activity{
			State:      "By " + song.By,
			Details:    "Listening " + song.Title,
			LargeImage: song.Img,
			Timestamps: &client.Timestamps{
				Start: &now,
			},
			Buttons: []*client.Button{
				{
					Label: "Play on YouTube",
					Url:   song.Url,
				},
				{
					Label: "Download ymp3cli",
					Url:   "https://github.com/paij0se/ymp3cli/releases/latest",
				},
			},
		})
		if id != nil {
			// lol this is so stupid
			fmt.Print("")
		}
	}
}
func Speedrpc(song string) {
	err := client.Login("851297648111517697")
	if err != nil {
		fmt.Println("No discord detected")
	}
	now := time.Now()
	err = client.SetActivity(client.Activity{
		Details:    "Remixing " + song,
		LargeImage: "https://cdn.discordapp.com/emojis/822805787771928597.webp?size=128&quality=lossless",
		Timestamps: &client.Timestamps{
			Start: &now,
		},
		Buttons: []*client.Button{
			{
				Label: "GitHub",
				Url:   "https://github.com/paij0se/ymp3cli",
			},
			{
				Label: "Website",
				Url:   "https://ymp3cli.tk",
			},
		},
	})

	if err != nil {
		fmt.Println("Error in rpc")
	}
	fmt.Print("")
}
