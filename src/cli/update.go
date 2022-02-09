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
		err := exec.Command(`cmd`, `/C`, "go get -u github.com/paij0se/ymp3cli").Run()
		if err != nil {
			log.Println(err)
		}
		s.Stop()
	} else {
		err := exec.Command("sh", "-c", "curl https://raw.githubusercontent.com/paij0se/ymp3cli/main/install.sh | bash").Run()
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Update complete!")
		s.Stop()
	}
	s.Stop()

}
