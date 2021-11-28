package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	noansi "github.com/ELPanaJose/api-deno-compiler/src/routes/others"
	"github.com/labstack/echo"
)

func DeleteRequest(c echo.Context) error {
	var d delete
	reqBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Fprintf(c.Response(), "Error")
	}
	json.Unmarshal(reqBody, &d)

	files, err := ioutil.ReadDir("music")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Println("file size:", file.Size())
	}
	// delete the song and send the response
	json.NewEncoder(c.Response()).Encode(map[string]string{"song_deleted": files[d.Delete].Name()})
	DeleteSong(d.Delete)
	return nil
}

func AskForPlayTheSong(c echo.Context) error {
	var n nsong
	reqBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Fprintf(c.Response(), "Error")
	}
	json.Unmarshal(reqBody, &n)

	fmt.Println(n.Nsong)
	files, err := ioutil.ReadDir("music")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Println("file size:", file.Size())
	}
	// play the song and send the response
	json.NewEncoder(c.Response()).Encode(map[string]string{"song_played": files[n.Nsong].Name()})
	PlaySongOneByOne(n.Nsong)

	return nil
}

func DownloadSong(c echo.Context) error {
	switch runtime.GOOS {
	case "linux", "darwin":
		var inputUrl url

		reqBody, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Fprintf(c.Response(), "Error")
		}
		json.Unmarshal(reqBody, &inputUrl)

		url := inputUrl.Url
		// check if the url is empty and match only youtube links
		switch {
		case len(url) == 0:
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)
			json.NewEncoder(c.Response()).Encode(map[string]string{"error": "empty url!"})
		case !v.MatchString(url):
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)
			json.NewEncoder(c.Response()).Encode(map[string]string{"error": "not a youtube url!"})
		default:
			fmt.Println(url)
			var stdout, stderr bytes.Buffer
			// download the video
			// https://www.youtube.com/watch?v=rcdvi74dUjQ

			cmd := exec.Command("sh", "-c", "youtube-dl -x --audio-format mp3 "+url)
			// show the output
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			peo := cmd.Run()
			if peo != nil {
				fmt.Println(err)
			}
			// capture the stderr and stdout
			executedOut := stdout.String() + stderr.String()
			out2 := strings.ReplaceAll(executedOut, "sh: 1: kill: No such process", "")
			output := noansi.NoAnsi(out2)
			fmt.Println(output)
			// send thge response
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)
			json.NewEncoder(c.Response()).Encode(map[string]string{"video_downloaded": url, "output": output, "status": "success"})
			// move the mp3 files
			MoveSong()

		}
		return nil
	case "windows":
		var inputUrl url

		reqBody, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Fprintf(c.Response(), "Error")
		}
		json.Unmarshal(reqBody, &inputUrl)

		url := inputUrl.Url
		// check if the url is empty and match only youtube links
		switch {
		case len(url) == 0:
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)
			json.NewEncoder(c.Response()).Encode(map[string]string{"error": "empty url!"})
		case !v.MatchString(url):
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)
			json.NewEncoder(c.Response()).Encode(map[string]string{"error": "not a youtube url!"})
		default:
			fmt.Println(url)
			var stdout, stderr bytes.Buffer
			// download the video
			cmd := exec.Command(`cmd`, `/C`, "youtube-dl -x --audio-format mp3 "+url)
			// show the output
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			peo := cmd.Run()
			if peo != nil {
				fmt.Println(err)
			}
			// capture the stderr and stdout
			output := stdout.String() + stderr.String()
			fmt.Println(output)
			// send thge response
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)
			json.NewEncoder(c.Response()).Encode(map[string]string{"video_downloaded": url, "output": output, "status": "success"})
			// move the mp3 files
			MoveSong()

		}
		return nil
	default:
		fmt.Println("unknown OS")
		return nil
	}
}
