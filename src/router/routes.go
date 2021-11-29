package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	controllers "github.com/paij0se/ymp3cli/src/controllers"
)

func SetUpRoutes() {
	e := echo.New()
	//e.Use(middleware.Logger())
	// Middlewares
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		c.String(http.StatusOK, "hello 💀")

		return nil
	})
	e.POST("/download", controllers.DownloadSongYt)
	e.POST("/spotify", controllers.SpotifyDownloader)
	e.POST("/y", controllers.AskForPlayTheSong)
	// Show the avaliable songs in the music directory.
	e.GET("/songs", func(c echo.Context) error {
		files, err := ioutil.ReadDir("music")

		if err != nil {
			fmt.Println(err)

		}

		for i, file := range files {
			json.NewEncoder(c.Response()).Encode(map[string]string{"[" + strconv.Itoa(i) + "]": file.Name()})

		}

		return nil
	})
	e.DELETE("/delete", controllers.DeleteRequest)

	e.Logger.Fatal(e.Start(":8000"))
}
