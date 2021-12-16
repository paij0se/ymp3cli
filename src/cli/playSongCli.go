package cli

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func PlaySongCli(song string) {
	fmt.Println("Playing: ", song)
	file, err := os.Open(song)

	if err != nil {
		log.Println(err)

	}

	defer file.Close()

	d, err := mp3.NewDecoder(file)

	if err != nil {

		fmt.Println(err)
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)

	if err != nil {

		fmt.Println(err)
	}

	defer c.Close()

	p := c.NewPlayer()

	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {

		fmt.Println(err)
	}
}
