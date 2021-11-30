package GET

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/labstack/echo"
)

func Songs(c echo.Context) error {
	files, err := ioutil.ReadDir("music")

	if err != nil {
		log.Println(err)

	}

	for _, file := range files {
		json.NewEncoder(c.Response()).Encode(file.Name())

	}

	return nil
}
