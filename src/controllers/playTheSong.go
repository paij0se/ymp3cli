package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	utils "github.com/paij0se/ymp3cli/src/utils"
)

func AskForPlayTheSong(c echo.Context) error {
	var n nsong
	reqBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Fprintf(c.Response(), "Error")
	}
	json.Unmarshal(reqBody, &n)

	//fmt.Println(n.Nsong)
	files, err := ioutil.ReadDir("music")
	if err != nil {
		fmt.Println(err)
	}
	// if the files is out of range, return error
	if n.Nsong > len(files) {
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().WriteHeader(http.StatusCreated)
		json.NewEncoder(c.Response()).Encode(map[string]string{"error": "out of range"})
		return nil
	} else {
		// play the song and send the response
		json.NewEncoder(c.Response()).Encode(map[string]string{"song_played": files[n.Nsong].Name()})
		utils.PlaySongOneByOne(n.Nsong)
		return nil
	}

}
