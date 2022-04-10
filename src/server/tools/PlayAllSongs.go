package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/dhowden/tag"
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
		defer f.Close()
		fmt.Println("Playing:", file.Name())
		e.GET("/currentSong", func(c echo.Context) error {
			dir, err := os.Getwd()
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			song, _ := os.Open("music/" + file.Name())
			defer song.Close()
			tag, err := tag.ReadFrom(song)
			if err != nil {
				log.Println(err)
				return c.String(http.StatusInternalServerError, err.Error())
			}
			match := regexp.MustCompile(`(|v\/|vi=|vi\/|youtu.be\/)[a-zA-Z0-9_-]{11}`)
			return c.JSON(http.StatusOK, map[string]string{"path": dir + "/music/" + f.Name(), "by": tag.Artist(), "title": tag.Title(), "img": "https://img.youtube.com/vi/" + match.FindString(tag.Comment()) + "/hqdefault.jpg", "file": f.Name(), "url": tag.Comment()})
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
