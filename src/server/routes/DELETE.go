package routes

import (
	"github.com/labstack/echo"
	"github.com/paij0se/ymp3cli/src/server/controllers/DELETE"
)

func Delete(e *echo.Echo) {
	e.DELETE("/delete", DELETE.Delete)

}
