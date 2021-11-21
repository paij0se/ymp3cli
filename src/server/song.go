package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	noansi "github.com/ELPanaJose/api-deno-compiler/src/routes/others"
	"github.com/labstack/echo"
)

/*

I need to storage the songs in a folder, and then I show the songs in the CLI client

*/
func DownloadSong(c echo.Context) error {
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
		// create the folder
		CreateFolder()
		// move the mp3 files
		MoveSong()
		// play the song
		cmd2 := exec.Command("sh", "-c", "mpg321 src/client/music/*.mp3")
		fmt.Println("Playing the song...")
		uwu := cmd2.Run()
		if uwu != nil {
			fmt.Println(err)
		}

	}
	// send the response with the headers
	return nil
}
