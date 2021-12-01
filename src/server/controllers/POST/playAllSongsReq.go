package POST

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/labstack/echo"
	"github.com/paij0se/ymp3cli/src/server/tools"
)

func PlayAllSongsReq(c echo.Context) error {
	var n tools.Nsong

	reqBody, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		fmt.Fprintf(c.Response(), "Error")

	}

	json.Unmarshal(reqBody, &n)

	if err != nil {
		log.Println(err)

	}

	json.Unmarshal(reqBody, &n)
	if n.Nsong == 1 {
		// play the song and send the response at the same time in  a goroutine

		tools.PlayAllSongs()

		json.NewEncoder(c.Response()).Encode(map[string]string{"shuffle": "Playing all songs"})
	}

	return nil
}
