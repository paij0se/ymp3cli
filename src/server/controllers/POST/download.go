package POST

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	noansi "github.com/ELPanaJose/api-deno-compiler/src/routes/others"
	"github.com/labstack/echo"
	"github.com/paij0se/ymp3cli/src/server/tools"
)

func Download(c echo.Context) error {
	switch runtime.GOOS {
	case "darwin", "linux":
		var inputUrl tools.Url

		reqBody, err := ioutil.ReadAll(c.Request().Body)

		if err != nil {
			fmt.Fprintf(c.Response(), "Error")

		}

		json.Unmarshal(reqBody, &inputUrl)

		url := inputUrl.Url

		if url == "" {
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)

			json.NewEncoder(c.Response()).Encode(map[string]string{"error": "empty url!"})

		} else if !tools.V.MatchString(url) {
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)

			json.NewEncoder(c.Response()).Encode(map[string]string{"error": "not a youtube url!"})

		} else {
			var stdout, stderr bytes.Buffer

			cmd := exec.Command("sh", "-c", ("youtube-dl -x --audio-format mp3 " + url))

			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			if cmd.Run() != nil {
				log.Println(err)

			}

			executedOut := stdout.String() + stderr.String()
			output := strings.ReplaceAll(executedOut, "sh: 1: kill: No such process", "")
			out := noansi.NoAnsi(output)

			log.Println(out)

			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)

			json.NewEncoder(c.Response()).Encode(map[string]string{"video_downloaded": url, "output": output, "status": "success"})

			tools.MoveSong()

		}

		return nil

	default:
		log.Println("Unknown OS")

		return nil
	}

}
