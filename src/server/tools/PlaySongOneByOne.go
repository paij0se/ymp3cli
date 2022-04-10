package tools

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/dhowden/tag"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/labstack/echo"
)

func PlaySongOneByOne(song uint32, e *echo.Echo) error {
	files, err := ioutil.ReadDir("music")

	if err != nil {
		log.Println(err)

	}

	if len(files) == 0 {
		log.Println("No stored music.")

		return nil
	}

	f, err := os.Open("music/" + files[song].Name())
	e.GET("/currentSong", func(c echo.Context) error {
		dir, err := os.Getwd()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		song, _ := os.Open("music/" + files[song].Name())
		defer song.Close()
		tag, err := tag.ReadFrom(song)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}
		match := regexp.MustCompile(`(|v\/|vi=|vi\/|youtu.be\/)[a-zA-Z0-9_-]{11}`)
		return c.JSON(http.StatusOK, map[string]string{"path": dir + "/music/" + f.Name(), "by": tag.Artist(), "title": tag.Title(), "img": "https://img.youtube.com/vi/" + match.FindString(tag.Comment()) + "/hqdefault.jpg", "file": f.Name(), "url": tag.Comment()})
	})

	if err != nil {
		log.Println(err)

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

	return nil
}
