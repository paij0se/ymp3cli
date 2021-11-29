package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func PlaySongOneByOne(song int) error {
	files, err := ioutil.ReadDir("music")
	if err != nil {
		fmt.Println(err)
	} // check if there is any mp3 file in the music folder
	if len(files) == 0 {
		fmt.Println("No mp3 files in the music folder.")
	} else {

		// play the song
		f, err := os.Open("music/" + files[song].Name())
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		d, err := mp3.NewDecoder(f)
		if err != nil {
			return err
		}

		c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
		if err != nil {
			return err
		}
		defer c.Close()

		p := c.NewPlayer()
		defer p.Close()

		if _, err := io.Copy(p, d); err != nil {
			return err
		}
	}

	return nil
}
