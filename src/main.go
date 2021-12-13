package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/paij0se/ymp3cli/src/client"
	"github.com/paij0se/ymp3cli/src/server"
	"github.com/phayes/freeport"
)

var (
	Version = "0.0.10"
)

func main() {
	// create the folder if it doesn't exist
	if _, err := os.Stat("music"); os.IsNotExist(err) {
		os.Mkdir("music", 0777)

	}
	switch {
	case len(os.Args) != 2:
		if _, err := os.Stat("music"); os.IsNotExist(err) {
			os.Mkdir("music", 0777)

		}

		port, err := freeport.GetFreePort()

		if err != nil {
			log.Panicln(err)

		}

		go client.StartClient(fmt.Sprintf(":%d", port))
		server.StartServer(fmt.Sprintf(":%d", port))
	case os.Args[1] == "--h" || os.Args[1] == "--help":
		fmt.Printf(`   
 Usage: ymp3cli --[OPTION]
 --h,--help: Display the help command
 --v,--version: Display the version of ymp3cli
 Usage: ymp3cli [SONG]
 ymp3cli <song.mp3>: play a single song
 example: ymp3cli music/song.mp3
    MIT License
    Made it by pai
    https://paijose.cf

`)
	case os.Args[1] == "--v" || os.Args[1] == "--version":
		fmt.Println(Version)
	case os.Args[1][len(os.Args[1])-4:] == ".mp3":
		fmt.Println("Playing: ", os.Args[1])
		file, err := os.Open(os.Args[1])

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

	default:
		port, err := freeport.GetFreePort()

		if err != nil {
			log.Panicln(err)

		}
		go client.StartClient(fmt.Sprintf(":%d", port))
		server.StartServer(fmt.Sprintf(":%d", port))
	}
}
