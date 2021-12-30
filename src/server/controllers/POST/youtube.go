package POST

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/labstack/echo"
	lmmp3 "github.com/paij0se/lmmp3"
	"github.com/paij0se/ymp3cli/src/server/tools"
)

func Download(c echo.Context) error {
	c.Response().WriteHeader(http.StatusCreated)
	var inputUrl tools.Url
	json.NewDecoder(c.Request().Body).Decode(&inputUrl)
	url := inputUrl.Url
	tools.ErrControl(c, "youtube", url, tools.V)
	if tools.ErrControl(c, "youtube", url, tools.V) {

		json.NewEncoder(c.Response()).Encode(map[string]string{"video_downloaded": url})
		switch runtime.GOOS {
		case "windows":
			lmmp3.DownloadAndConvert(url)
			del := exec.Command("cmd", "/C", "del", "*.mpeg")
			if del.Run() != nil {
				panic("failed to delete files")
			}
			tools.MoveSong()
		default:
			lmmp3.DownloadAndConvert(url)
			tools.MoveSong()
		}

	}
	return nil

}
