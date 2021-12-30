package DELETE

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/labstack/echo"
	"github.com/paij0se/ymp3cli/src/server/tools"
)

func Delete(c echo.Context) error {
	var deleted tools.Delete

	json.NewDecoder(c.Request().Body).Decode(&deleted)

	files, err := ioutil.ReadDir("music")

	if err != nil {
		log.Panicln(err)

	}

	if deleted.Delete > len(files) {

		json.NewEncoder(c.Response()).Encode(map[string]string{"error": "Out of range"})

		return nil
	}

	json.NewEncoder(c.Response()).Encode(map[string]string{"song_deleted": files[deleted.Delete].Name()})
	tools.DeleteSong(deleted.Delete)

	return nil
}
