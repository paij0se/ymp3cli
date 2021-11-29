package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func MoveSong() {
	switch runtime.GOOS {
	case "linux", "darwin":
		del := exec.Command("sh", "-c", "mv *.mp3 music")
		//fmt.Println("all mp3 files moved to the music folder")
		delError := del.Run()
		if delError != nil {
			fmt.Println(delError)
		}
	case "windows":
		del := exec.Command(`cmd`, `/C`, "move *.mp3 music")
		//fmt.Println("all mp3 files moved to the music folder")
		delError := del.Run()
		if delError != nil {
			fmt.Println(delError)
		}
	default:
		fmt.Println("unknown OS")
	}
}

func DeleteSong(song int) error {
	files, err := ioutil.ReadDir("music")
	if err != nil {
		fmt.Println(err)
	}
	os.Remove("music/" + files[song].Name())
	return nil
}

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
