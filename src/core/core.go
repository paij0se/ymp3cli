package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	server "github.com/paij0se/ymp3cli/src/server"
)

func Core() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, "hello 💀")

		return nil
	})

	e.POST("/download", server.DownloadSong)
	e.POST("/spotify", server.SpotifyDownloader)

	e.POST("/y", server.AskForPlayTheSong)

	// Show the songs / music stored in the folder.
	e.GET("/songs", func(c echo.Context) error {
		files, err := ioutil.ReadDir("music")

		if err != nil {
			log.Fatal(err)

		}

		for i, file := range files {
			json.NewEncoder(c.Response()).Encode(map[string]string{"[" + strconv.Itoa(i) + "]": file.Name()})

		}

		return nil
	})
	e.DELETE("/delete", server.DeleteRequest)

	e.Logger.Fatal(e.Start(":8000"))
}
