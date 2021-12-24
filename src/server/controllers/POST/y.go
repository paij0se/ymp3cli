package POST

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/paij0se/ymp3cli/src/server/tools"
)

func Y(c echo.Context) error {
	var n tools.Nsong

	reqBody, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		fmt.Fprintf(c.Response(), "Error")

	}

	json.Unmarshal(reqBody, &n)

	files, err := ioutil.ReadDir("music")

	if err != nil {
		log.Println(err)

	}

	json.Unmarshal(reqBody, &n)

	if n.Nsong > len(files) {
		c.Response().WriteHeader(http.StatusBadRequest)
		c.Response().WriteHeader(http.StatusCreated)

		json.NewEncoder(c.Response()).Encode(map[string]string{"error": "out of range"})

		return nil
	}

	json.NewEncoder(c.Response()).Encode(map[string]string{"song_played": files[n.Nsong].Name()})

	tools.PlaySongOneByOne(uint32(n.Nsong), c.Echo())

	return nil
}
