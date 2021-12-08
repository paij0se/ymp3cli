package POST

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
	lmmp3 "github.com/paij0se/lmmp3"
	"github.com/paij0se/ymp3cli/src/server/tools"
)

func Download(c echo.Context) error {
	var inputUrl tools.Url
	reqBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Fprintf(c.Response(), "Error")

	}
	json.Unmarshal(reqBody, &inputUrl)
	url := inputUrl.Url
	switch {
	case url == "":
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().WriteHeader(http.StatusCreated)
		json.NewEncoder(c.Response()).Encode(map[string]string{"error": "empty url!"})
	case !tools.V.MatchString(url):
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().WriteHeader(http.StatusCreated)
		json.NewEncoder(c.Response()).Encode(map[string]string{"error": "not a youtube url!"})
	default:
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().WriteHeader(http.StatusCreated)
		json.NewEncoder(c.Response()).Encode(map[string]string{"video_downloaded": url})
		lmmp3.DownloadAndConvert(url) // download and convert the video
		del := exec.Command("cmd", "/C", "del", "*.mpeg")
		if del.Run() != nil {
			panic("failed to delete files")
		}
		tools.MoveSong() // move the song to the music folder
	}
	return nil

}
