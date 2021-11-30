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

func Spotify(c echo.Context) error {
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

		} else if !tools.S.MatchString(url) {
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)

			json.NewEncoder(c.Response()).Encode(map[string]string{"error": "not a spotify url!"})

		} else {
			var stdout, stderr bytes.Buffer

			cmd := exec.Command("sh", "-c", "spotdl "+url)

			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			if cmd.Run() != nil {
				log.Println(err)

			}

			executedOut := stdout.String() + stderr.String()
			output := strings.ReplaceAll(executedOut, "sh: 1: kill: No such process", "")
			out := noansi.NoAnsi(output)

			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)

			json.NewEncoder(c.Response()).Encode(map[string]string{"url": url, "output": out, "status": "success"})

			tools.MoveSong()

		}

		return nil

	case "windows":
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

		} else if !tools.S.MatchString(url) {
			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)

			json.NewEncoder(c.Response()).Encode(map[string]string{"error": "not a spotify url!"})

		} else {
			var stdout, stderr bytes.Buffer

			cmd := exec.Command("cmd", "/c", ("spotdl " + url))

			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			if cmd.Run() != nil {
				log.Println(err)

			}

			output := stdout.String() + stderr.String()

			c.Response().Header().Set("Content-Type", "application/json")
			c.Response().WriteHeader(http.StatusCreated)

			json.NewEncoder(c.Response()).Encode(map[string]string{"url": url, "output": output, "status": "success"})

			tools.MoveSong()
		}

		return nil

	default:
		log.Println("Unknown OS")

		return nil
	}

}
