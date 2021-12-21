package rpc

import (
	"fmt"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func RpcListening(songName string) {
	err := client.Login("851297648111517697")
	if err != nil {
		fmt.Println("No discord detected")
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		State:      "🎵🖥️",
		Details:    "Listening " + songName,
		LargeImage: "logo",
		LargeText:  "🙏",
		SmallImage: "wallpaperbetter_com_1366x768",
		SmallText:  "yessir",
		Timestamps: &client.Timestamps{
			Start: &now,
		},
		Buttons: []*client.Button{
			&client.Button{
				Label: "GitHub",
				Url:   "https://github.com/paij0se/ymp3cli",
			},
		},
	})

	if err != nil {
		fmt.Println("Error in rpc")
	}
	time.Sleep(time.Second * 10)
	fmt.Println("")
}

func YtRpc(url string) {
	err := client.Login("851297648111517697")
	if err != nil {
		fmt.Println("No discord detected")
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		State:      "🎵🖥️",
		Details:    "Downloading... " + url,
		LargeImage: "logo",
		LargeText:  "🙏",
		SmallImage: "wallpaperbetter_com_1366x768",
		SmallText:  "yessir",
		Timestamps: &client.Timestamps{
			Start: &now,
		},
		Buttons: []*client.Button{
			&client.Button{
				Label: "GitHub",
				Url:   "https://github.com/paij0se/ymp3cli",
			},
		},
	})

	if err != nil {
		fmt.Println("Error in rpc")
	}

	time.Sleep(time.Second * 10)
	fmt.Println("")

}

func DefaultRpc() {
	err := client.Login("851297648111517697")
	if err != nil {
		fmt.Println("No discord detected")
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		State:      "🎵🖥️",
		Details:    "Listening music in the terminal",
		LargeImage: "logo",
		LargeText:  "🙏",
		SmallImage: "wallpaperbetter_com_1366x768",
		SmallText:  "yessir",
		Timestamps: &client.Timestamps{
			Start: &now,
		},
		Buttons: []*client.Button{
			&client.Button{
				Label: "GitHub",
				Url:   "https://github.com/paij0se/ymp3cli",
			},
		},
	})

	if err != nil {
		fmt.Println("Error in rpc")
	}

	time.Sleep(time.Second * 10)
	fmt.Println("")
}
