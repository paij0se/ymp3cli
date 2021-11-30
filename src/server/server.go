package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/paij0se/ymp3cli/src/server/routes"
)

func StartServer(port string) {
	e := echo.New()

	// Middleware.
	e.Use(middleware.Recover())

	// Start routes.
	routes.Delete(e)
	routes.Get(e)
	routes.Post(e)

	// Start server.
	e.Logger.Fatal(e.Start(port))
}
