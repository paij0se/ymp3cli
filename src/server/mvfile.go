package server

import (
	"fmt"
	"os/exec"
)

func CreateFolder() {
	del := exec.Command("sh", "-c", "mkdir src/music")
	fmt.Println("music folder created")
	delError := del.Run()
	if delError != nil {
		fmt.Println(delError)
	}
}

func MoveSong() {
	del := exec.Command("sh", "-c", "mv *.mp3 src/music")
	fmt.Println("all mp3 files moved to the music folder")
	delError := del.Run()
	if delError != nil {
		fmt.Println(delError)
	}
}
