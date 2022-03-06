package POST

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/paij0se/ymp3cli/src/cli/scdl/pkg/soundcloud"
	"github.com/paij0se/ymp3cli/src/server/tools"
)

func DownloadSouncloud(c echo.Context) error {
	c.Response().WriteHeader(http.StatusCreated)
	var inputUrl tools.Url
	json.NewDecoder(c.Request().Body).Decode(&inputUrl)
	url := inputUrl.Url
	soundcloud.ExtractSong(url)
	tools.MoveSong()
	return nil

}
