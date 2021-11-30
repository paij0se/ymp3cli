package client

import (
	"net"
	"time"

	"github.com/paij0se/ymp3cli/src/client/controllers"
	"github.com/paij0se/ymp3cli/src/client/handlers"
	"github.com/paij0se/ymp3cli/src/client/helpers"
	"github.com/paij0se/ymp3cli/src/client/validators"
)

var questions = []handlers.BaseQuestion{
	{
		Name: "Download a song",
		Option: &handlers.QuestionOption{
			Handler: controllers.DownloadSound,
			Label:   "Enter the URL of the video to download.",
			Options: func(url string) []string {
				return []string{}
			},
			Validator: validators.String,
		},
	},
	{
		Name: "Listen a song",
		Option: &handlers.QuestionOption{
			Handler:   controllers.PlaySound,
			Label:     "Enter a number to play the sound",
			Options:   controllers.GetSongs,
			Validator: validators.Number,
		},
	},
	{
		Name: "Delete a song",
		Option: &handlers.QuestionOption{
			Handler:   controllers.DeleteSound,
			Label:     "Enter a number to delete sound",
			Options:   controllers.GetSongs,
			Validator: validators.Number,
		},
	},
}

func StartClient(port string) {
	baseURL := "http://localhost" + port + "/"

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
