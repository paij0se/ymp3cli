package server

import (
	"fmt"
	"os/exec"
)

func PlayMuisc() {
	del := exec.Command("sh", "-c", "mpg321 music/*.mp3")
	fmt.Println("music folder created")
	delError := del.Run()
	if delError != nil {
		fmt.Println(delError)
	}
}

func KillSong() {
	del := exec.Command("sh", "-c", "killall mpg321")
	fmt.Println("mpg321 killed.")
	delError := del.Run()
	if delError != nil {
		fmt.Println(delError)
	}
}

func MoveSong() {
	del := exec.Command("sh", "-c", "mv *.mp3 music")
	fmt.Println("all mp3 files moved to the music folder")
	delError := del.Run()
	if delError != nil {
		fmt.Println(delError)
	}
}
