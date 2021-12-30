package POST

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strings"

	noansi "github.com/ELPanaJose/api-deno-compiler/src/routes/others"
	"github.com/labstack/echo"
	"github.com/paij0se/ymp3cli/src/server/tools"
)

//this functions works for check any kind of error of the client

func Spotify(c echo.Context) (err error) {
	var stdout, stderr bytes.Buffer
	var inputUrl tools.Url
	c.Response().WriteHeader(http.StatusCreated)

	err = json.NewDecoder(c.Request().Body).Decode(&inputUrl)

	url := inputUrl.Url

	if tools.ErrControl(c, "spotify", url, tools.S) {

		cmd := exec.Command("python3", "-m", "spotdl", url)

		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err = cmd.Run(); err != nil {
			log.Println(err)
			stderr.Write([]byte(err.Error()))

		}

		executedOut := stdout.String() + stderr.String()
		output := strings.ReplaceAll(executedOut, "sh: 1: kill: No such process", "")
		out := noansi.NoAnsi(output)

		json.NewEncoder(c.Response()).Encode(map[string]string{"url": url, "output": out, "status": "success"})

		tools.MoveSong()
	}

	return

}
