package tools

import (
	"log"
	"os/exec"
	"runtime"
)

func MoveSong() {
	switch runtime.GOOS {

	case "linux", "darwin":
		err := exec.Command("sh", "-c", "mv *.mp3 music").Run()

		if err != nil {
			log.Println(err)

		}

	case "windows":
		err := exec.Command(`cmd`, `/C`, "move *.mp3 music").Run()

		if err != nil {
			log.Println(err)

		}

	default:
		log.Println("Unknown OS")

	}
}
