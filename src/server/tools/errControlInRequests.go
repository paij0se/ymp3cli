package tools

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/labstack/echo"
)

func ErrControl(c echo.Context, origin, url string, reg *regexp.Regexp) (output bool) {
	output = true
	if output = !(url == ""); !output {

		json.NewEncoder(c.Response()).Encode(map[string]string{"error": "empty url!"})
		return

	} else if output = reg.MatchString(url); !output {

		json.NewEncoder(c.Response()).Encode(map[string]string{"error": fmt.Sprintf("not a %s url!", origin)})
		return
	}
	return

}
