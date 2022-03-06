package cli

import (
	"github.com/paij0se/ymp3cli/src/cli/scdl/pkg/soundcloud"
)

func SoundCloudTest(url string) {
	soundcloud.ExtractSong(url)
}
