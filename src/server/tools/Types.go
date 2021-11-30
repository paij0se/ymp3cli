package tools

import "regexp"

var (
	V = regexp.MustCompile(`(http:|https:)?\/\/(www\.)?(youtube.com|youtu.be)\/(watch)?(\?v=)?(\S+)?`)
	S = regexp.MustCompile(`(http:|https:)?\/\/(www\.)?(open.spotify.com|spotify.com)\/(watch)?(\?v=)?(\S+)?`)
)

type Url struct {
	Url string
}

type Nsong struct {
	Nsong int
}

type Delete struct {
	Delete int
}
