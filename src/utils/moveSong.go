package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

func MoveSong() {
	switch runtime.GOOS {
	case "linux", "darwin":
		del := exec.Command("sh", "-c", "mv *.mp3 music")
		//fmt.Println("all mp3 files moved to the music folder")
		delError := del.Run()
		if delError != nil {
			fmt.Println(delError)
		}
	case "windows":
		del := exec.Command(`cmd`, `/C`, "move *.mp3 music")
		//fmt.Println("all mp3 files moved to the music folder")
		delError := del.Run()
		if delError != nil {
			fmt.Println(delError)
		}
	default:
		fmt.Println("unknown OS")
	}
}
