package core

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	routes "github.com/paij0se/ymp3/src/server"
)

func Core() {
	// clear the screen
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
	e.POST("/download", routes.DownloadSong)

	e.Logger.Fatal(e.Start(":8000"))
}
