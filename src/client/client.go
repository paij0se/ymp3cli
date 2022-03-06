package client

import (
	"fmt"
	"net"
	"time"

	clear "github.com/ELPanaJose/pairat/src/utils"
	"github.com/paij0se/ymp3cli/src/client/controllers"
	"github.com/paij0se/ymp3cli/src/client/handlers"
	"github.com/paij0se/ymp3cli/src/client/helpers"
	"github.com/paij0se/ymp3cli/src/client/validators"
)

var questions = []handlers.BaseQuestion{
	{
		Name: "Download a Song from Youtube🔴",
		Option: &handlers.QuestionOption{
			Handler: controllers.DownloadSound,
			Label:   "Enter the Youtube URL of the video to download🔴",
			Options: func(url string) []string {
				return []string{}
			},
			Validator: validators.String,
		},
	},
	{
		Name: "Download a Song from SoundCloud☁️",
		Option: &handlers.QuestionOption{
			Handler: controllers.DownloadSoundcloud,
			Label:   "Enter the SoundCloud URL of the song to download☁️",
			Options: func(url string) []string {
				return []string{}
			},
			Validator: validators.String,
		},
	},
	{
		Name: "Download a Song/Playlist from Spotify🥗",
		Option: &handlers.QuestionOption{
			Handler: controllers.DownloadSpotify,
			Label:   "Enter the Spotify url to download (playlist/song)🥗",
			Options: func(url string) []string {
				return []string{}
			},
			Validator: validators.String,
		},
	},
	{
		Name: "Suffle the Playlist🔄",
		Option: &handlers.QuestionOption{
			Handler:   controllers.PlaySoundAll,
			Label:     "Enter any number to shuffle the Playlist🔄",
			Options:   controllers.GetSongs,
			Validator: validators.Number,
		},
	},
	{
		Name: "Listen a single Song🎵",
		Option: &handlers.QuestionOption{
			Handler:   controllers.PlaySound,
			Label:     "Enter a number to play the Song🎵",
			Options:   controllers.GetSongs,
			Validator: validators.Number,
		},
	},
	{
		Name: "Delete a song🔻",
		Option: &handlers.QuestionOption{
			Handler:   controllers.DeleteSound,
			Label:     "Enter a number to delete Song🔻",
			Options:   controllers.GetSongs,
			Validator: validators.Number,
		},
	},
}

func StartClient(port string) {
	baseURL := "http://localhost" + port + "/"
	// clear the screen
	clear.Clear()
	fmt.Println("Server started at: " + baseURL)

	for {
		_, err := net.DialTimeout("tcp", ("localhost" + port), (time.Millisecond * 200))

		if err == nil {

			break
		}

		time.Sleep(time.Millisecond * 200)
	}

	for {
		helpers.Welcome()

		handlers.QuestionHandler(baseURL, questions)
	}
}
