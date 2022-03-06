package cmd

import (
	"github.com/paij0se/ymp3cli/src/cli/scdl/pkg/soundcloud"
)

func scdl(args []string) {
	url := args[0]

	if Find == true {

		soundcloud.Search(url)
		// exit
		return
	}
	// download song
	soundcloud.ExtractSong(url)

}
