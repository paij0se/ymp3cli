package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/labstack/echo"
)

func PlayAllSongs(e *echo.Echo) error {
	files, err := ioutil.ReadDir("music")
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		fmt.Println("🐶No songs found, Download them👻")
	}
	for _, file := range files {
		os.Open("music/" + file.Name())
		f, err := os.Open("music/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Playing:", file.Name())
		e.GET("/currentSong", func(c echo.Context) error {
			return c.String(http.StatusOK, file.Name()[:len(file.Name())-4])
		})
		streamer, format, err := mp3.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		defer streamer.Close()
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		done := make(chan bool)
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))
		<-done
	}
	return nil
}
