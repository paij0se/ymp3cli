package tools

import (
	"io/ioutil"
	"log"
	"os"
)

func DeleteSong(song int) error {
	files, err := ioutil.ReadDir("music")

	if err != nil {
		log.Panicln(err)

	}

	os.Remove("music/" + files[song].Name())

	return nil
}
