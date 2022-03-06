package POST

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/labstack/echo"
	"github.com/paij0se/ymp3cli/src/server/tools"
)

func Spotify(c echo.Context) (err error) {
	var inputUrl tools.Url
	c.Response().WriteHeader(http.StatusCreated)

	err = json.NewDecoder(c.Request().Body).Decode(&inputUrl)

	url := inputUrl.Url

	if tools.ErrControl(c, "spotify", url, tools.S) {

		switch runtime.GOOS {
		case "linux", "darwin":
			cmd := exec.Command("/bin/bash", "-c", "spotdl "+url)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}

		case "windows":
			cmd := exec.Command("cmd", "/c", "spotdl "+url)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}

		default:
			fmt.Println("Your OS is not supported")
		}

		json.NewEncoder(c.Response()).Encode(url + "  downloaded")

		tools.MoveSong()
	}

	return

}
