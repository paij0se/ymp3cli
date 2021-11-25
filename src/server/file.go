package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func MoveSong() {
	del := exec.Command("sh", "-c", "mv *.mp3 music")
	fmt.Println("all mp3 files moved to the music folder")
	delError := del.Run()
	if delError != nil {
		fmt.Println(delError)
	}
}

func DeleteSong(song int) error {
	files, err := ioutil.ReadDir("music")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		os.Remove("music/" + files[song].Name())
		fmt.Println("file size:", file.Size())
	}

	return nil
}

func PlaySongOneByOne(song int) error {
	files, err := ioutil.ReadDir("music")
	if err != nil {
		log.Fatal(err)
	} // check if there is any mp3 file in the music folder
	if len(files) == 0 {
		fmt.Println("No mp3 files in the music folder.")
	} else {

		for _, file := range files {
			// play the song
			f, err := os.Open("music/" + files[song].Name())
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			d, err := mp3.NewDecoder(f)
			if err != nil {
				return err
			}
			fmt.Println("file size:", file.Size())

			c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
			if err != nil {
				return err
			}
			defer c.Close()

			p := c.NewPlayer()
			defer p.Close()

			fmt.Printf("playing: %s\n", f.Name())

			if _, err := io.Copy(p, d); err != nil {
				return err
			}
		}
	}
	return nil
}
