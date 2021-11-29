package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func DeleteSong(song int) error {
	files, err := ioutil.ReadDir("music")
	if err != nil {
		fmt.Println(err)
	}
	os.Remove("music/" + files[song].Name())
	return nil
}
