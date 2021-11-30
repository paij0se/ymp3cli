package DELETE

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/paij0se/ymp3cli/src/server/tools"
)

func Delete(c echo.Context) error {
	var deleted tools.Delete

	reqBody, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		fmt.Fprintf(c.Response(), "Error")

	}

	json.Unmarshal(reqBody, &deleted)

	files, err := ioutil.ReadDir("music")

	if err != nil {
		log.Panicln(err)

	}

	if deleted.Delete > len(files) {
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().WriteHeader(http.StatusCreated)

		json.NewEncoder(c.Response()).Encode(map[string]string{"error": "Out of range"})

		return nil
	}

	json.NewEncoder(c.Response()).Encode(map[string]string{"song_deleted": files[deleted.Delete].Name()})
	tools.DeleteSong(deleted.Delete)

	return nil
}
