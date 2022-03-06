package routes

import (
	"github.com/labstack/echo"
	"github.com/paij0se/ymp3cli/src/server/controllers/POST"
)

func Post(e *echo.Echo) {
	e.POST("/download", POST.Download)
	e.POST("/download/soundcloud", POST.DownloadSouncloud)
	e.POST("/spotify", POST.Spotify)
	e.POST("/y", POST.Y)
	e.POST("/y/all", POST.PlayAllSongsReq)
}
