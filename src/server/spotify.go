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

func SpotifyDownloader(c echo.Context) error {
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
		case url == "":
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)
			json.NewEncoder(c.Response()).Encode(map[string]string{"error": "empty url!"})
		case !s.MatchString(url):
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)
			json.NewEncoder(c.Response()).Encode(map[string]string{"error": "not a spotify url!"})
		default:
			//fmt.Println(url)
			var stdout, stderr bytes.Buffer
			// download the video

			cmd := exec.Command("sh", "-c", "spotdl "+url)
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
			// send thge response
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)
			json.NewEncoder(c.Response()).Encode(map[string]string{"url": url, "output": output, "status": "success"})
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
		case url == "":
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)
			json.NewEncoder(c.Response()).Encode(map[string]string{"error": "empty url!"})
		case !s.MatchString(url):
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)
			json.NewEncoder(c.Response()).Encode(map[string]string{"error": "not a spotify url!"})
		default:
			//fmt.Println(url)
			var stdout, stderr bytes.Buffer
			// download the video
			cmd := exec.Command(`cmd`, `/C`, "spotdl "+url)
			// show the output
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			peo := cmd.Run()
			if peo != nil {
				fmt.Println(err)
			}
			// capture the stderr and stdout
			output := stdout.String() + stderr.String()
			//fmt.Println(output)
			// send thge response
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)
			json.NewEncoder(c.Response()).Encode(map[string]string{"url": url, "output": output, "status": "success"})
			// move the mp3 files
			MoveSong()

		}
		return nil
	default:
		fmt.Println("unknown OS")
		return nil
	}

}
