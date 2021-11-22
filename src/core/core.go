package core

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	server "github.com/paij0se/ymp3cli/src/server"
)

func Core() {
	// kill mpg321
	server.KillSong()
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

	e.Logger.Fatal(e.Start(":8000"))
}
