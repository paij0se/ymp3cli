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
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	/*
		Routes
	*/
	e.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, "hello 💀")
		return nil
	})
	e.POST("/download", server.DownloadSong)
	e.POST("/y", server.AskForPlayTheSong)
	/*

	   show the songs in the music folder

	*/
	e.GET("/songs", func(c echo.Context) error {
		files, err := ioutil.ReadDir("music")
		if err != nil {
			log.Fatal(err)
		}
		for i, file := range files {
			n := strconv.Itoa(i)
			json.NewEncoder(c.Response()).Encode(map[string]string{"[" + n + "]": file.Name()})

		}
		return nil
	})

	e.Logger.Fatal(e.Start(":8000"))
}
