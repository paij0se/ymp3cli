package rpc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/dhowden/tag"
	"github.com/hugolgst/rich-go/client"
)

func Rpc(port int) string {
	now := time.Now()
	for {
		time.Sleep(time.Second * 1)
		url := "http://localhost:" + fmt.Sprintf("%d", port) + "/currentSong"
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		r := string(body)
		if !strings.HasSuffix(r, ".mp3") {
		} else {
			f, err := os.Open(r)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			m, err := tag.ReadFrom(f)
			if err != nil {
				log.Println(err)
			}
			match := regexp.MustCompile(`(|v\/|vi=|vi\/|youtu.be\/)[a-zA-Z0-9_-]{11}`)
			img := "https://img.youtube.com/vi/" + match.FindString(m.Comment()) + "/hqdefault.jpg"
			id := client.Login("851297648111517697")
			if err != nil {
				fmt.Println("No discord detected")
			}
			err = client.SetActivity(client.Activity{
				State:      "By " + m.Artist(),
				Details:    "Listening " + m.Title(),
				LargeImage: img,
				Timestamps: &client.Timestamps{
					Start: &now,
				},
				Buttons: []*client.Button{
					&client.Button{
						Label: "Play on YouTube",
						Url:   m.Comment(),
					},
					{
						Label: "Download ymp3cli",
						Url:   "https://github.com/paij0se/ymp3cli/releases/latest",
					},
				},
			})

			if id != nil {
				fmt.Println("Error in rpc")
			}
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
			&client.Button{
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
