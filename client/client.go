package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	clear "github.com/ELPanaJose/pairat/src/utils"
	"github.com/manifoldco/promptui"
	"github.com/paij0se/ymp3cli/client/cli"
)

func cliDeleteSong(deleteSong int) {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		return err
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "❓🎶Number of the song you want to delete (you can type <99 + Enter> to skip)🐻:",
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	number, err := strconv.Atoi(result)
	if err != nil {
		fmt.Printf("Error converting %s to int %v\n", result, err)
		return
	}
	deleteSong = number

	cli.DeleteSong(deleteSong)
}

func Clientmain() {
	for {
		//open the welcome.txt file
		file, err := os.Open("client/welcome.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		// convert file into string
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		// print the content of the text file
		fmt.Println(string(data))
		fmt.Printf("version 0.0.5\nType <ctrl + c> to exit.\n")
		fmt.Println("Avaliable songs:")
		// get the songs
		cli.GetSongs()
		// ask for download the song
		validateUrl := func(input string) error {
			if len(input) == 0 {
				return errors.New("empty url")
			}
			return nil
		}

		var url string
		templates := &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Valid:   "{{ . | green }} ",
			Invalid: "{{ . | red }} ",
			Success: "{{ . | bold }} ",
		}
		promptUrl := promptui.Prompt{
			Label:     "🤓Insert a Youtube url to download the song (you can type <99 + Enter> for skip)🍱:",
			Validate:  validateUrl,
			Templates: templates,
			Default:   url,
		}

		resultUrl, err := promptUrl.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		// Download the song
		cli.DownloadRequest(resultUrl)
		// spotify download ------------------------------------------------------------------------------------
		validateSpotify := func(input string) error {
			if len(input) == 0 {
				return errors.New("empty url")
			}
			return nil
		}

		promptSpotify := promptui.Prompt{
			Label:     "🐢Insert a Spotify url to download the song/playlist (you can type <99 + Enter> for skip)🥒:",
			Validate:  validateSpotify,
			Templates: templates,
			Default:   url,
		}

		resultSpotify, err2 := promptSpotify.Run()

		if err2 != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		// download the spotify song/playlist
		cli.DownloadSpotify(resultSpotify)
		//show the avaliable songs
		fmt.Println("Avaliable songs:")
		cli.GetSongs()
		// ask the number of the song you want to listen
		validate := func(input string) error {
			_, err := strconv.ParseFloat(input, 64)
			return err
		}

		prompt := promptui.Prompt{
			Label:     "❓🎶Number of the song you want to listen (you can type <99 + Enter> to skip)🙈:",
			Templates: templates,
			Validate:  validate,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		// convert restult to int
		number, err := strconv.Atoi(result)
		if err != nil {
			fmt.Printf("Error converting %s to int %v\n", result, err)
			return
		}
		cli.ChooseSong(number)
		// ask if delete the song
		cliDeleteSong(number)
		// clear the screen
		time.Sleep(time.Second * 1)
		clear.Clear()
	}
}
