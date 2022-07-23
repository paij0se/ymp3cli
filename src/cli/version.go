package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Version struct {
	Name string `json:"name"`
}

func CheckVersion(v string) {
	res, err := http.Get("https://api.github.com/repos/paij0se/ymp3cli/releases/latest")
	if err != nil {
		log.Println("Unable to check the version 😔")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var version Version
	json.Unmarshal(body, &version)
	version.Name = strings.Replace(version.Name, "{", "", -1)
	version.Name = strings.Replace(version.Name, "}", "", -1)
	version.Name = strings.Replace(version.Name, "v", "", -1)
	// {v0.7.0} -> 0.7.0
	if version.Name != v {
		fmt.Println("New version available! 🎉", version.Name, "Your current version is:", v)
		fmt.Println("You can update ymp3cli with the following command:")
		fmt.Println("Linux & Mac:")
		fmt.Println("$ ymp3cli -u")
		fmt.Println("Windows:")
		fmt.Println("https://github.com/paij0se/ymp3cli/releases/latest")
		time.Sleep(time.Second * 10)
	}

}
