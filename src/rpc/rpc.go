package rpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/hugolgst/rich-go/client"
)

type Respose struct {
	Message string
}

func DefaultRpc(port int) {
	images := []string{"https://cdn.discordapp.com/emojis/756615654404259870.webp?size=96&quality=lossless", "https://cdn.discordapp.com/emojis/763985835329978398.webp?size=96&quality=lossless", "https://cdn.discordapp.com/emojis/897294079275192350.webp?size=96&quality=lossless", "https://cdn.discordapp.com/emojis/765001183869403176.webp?size=96&quality=lossless", "https://cdn.discordapp.com/emojis/937819339233574952.webp?size=96&quality=lossless"}
	randImage := rand.Intn(len(images))
	pick := images[randImage]

	err := client.Login("851297648111517697")
	if err != nil {
		fmt.Println("No discord detected")
	}
	for {
		time.Sleep(time.Second * 1)

		url := "http://localhost:" + fmt.Sprintf("%d", port) + "/currentSong"
		resp, err := http.Get(url)

		if err != nil {
			log.Fatalln(err)

		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)

		}

		r := string(body)
		song := "Listening " + r
		var res Respose
		json.Unmarshal([]byte(r), &res)
		if res.Message == "Not Found" { // {"message":"Not Found"}
			song = "Listening Nothing💀"
		}
		err = client.SetActivity(client.Activity{
			State:      "🎵🖥️",
			Details:    song,
			LargeImage: pick,
			LargeText:  "🙏",
			SmallImage: "wallpaperbetter_com_1366x768",
			SmallText:  "yessir",
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

}

func Speedrpc(song string) {
	images := []string{"https://cdn.discordapp.com/emojis/756615654404259870.webp?size=96&quality=lossless", "https://cdn.discordapp.com/emojis/763985835329978398.webp?size=96&quality=lossless", "https://cdn.discordapp.com/emojis/897294079275192350.webp?size=96&quality=lossless", "https://cdn.discordapp.com/emojis/765001183869403176.webp?size=96&quality=lossless", "https://cdn.discordapp.com/emojis/937819339233574952.webp?size=96&quality=lossless"}
	randImage := rand.Intn(len(images))
	pick := images[randImage]
	err := client.Login("851297648111517697")
	if err != nil {
		fmt.Println("No discord detected")
	}
	for {
		time.Sleep(time.Second * 1)
		err = client.SetActivity(client.Activity{
			State:      "🥷🏿",
			Details:    "remixing " + song,
			LargeImage: pick,
			LargeText:  "🙏",
			SmallImage: "wallpaperbetter_com_1366x768",
			SmallText:  "yessir",
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

}
