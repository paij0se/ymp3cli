package cli

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
)

func Update() {
	s := spinner.New(spinner.CharSets[38], 100*time.Millisecond) // Build our new spinner
	s.Start()
	s.Suffix = " Updating..."
	if runtime.GOOS == "windows" {
		fmt.Println("Download the last verion of ymp3cli from https://github.com/paij0se/ymp3cli/releases/latest")
		s.Stop()
	} else {
		err := exec.Command("sh", "-c", "curl https://raw.githubusercontent.com/paij0se/ymp3cli/main/install.sh | bash").Run()
		if err != nil {
			log.Println("Error updating ymp3cli, Download manually from https://github.com/paij0se/ymp3cli/releases/latest")
		}
		fmt.Println("Update complete!")
		s.Stop()
	}
	s.Stop()

}
