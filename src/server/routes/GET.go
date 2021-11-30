package routes

import (
	"github.com/labstack/echo"
	"github.com/paij0se/ymp3cli/src/server/controllers/GET"
)

func Get(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {

		return c.String(200, "Hello, World!")
	})

	e.GET("/songs", GET.Songs)

}
