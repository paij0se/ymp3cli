package main

import (
	"os/exec"
	"testing"

	"github.com/paij0se/lmmp3"
)

func TestSound(t *testing.T) {
	lmmp3.DownloadAndConvert("https://www.youtube.com/watch?v=5xYDXp7fkY4")
	del := exec.Command("cmd", "/C", "del", "*.mpeg")
	if del.Run() != nil {
		panic("failed to delete files")
	}
}
