package server

import "regexp"

var (
	v = regexp.MustCompile(`(http:|https:)?\/\/(www\.)?(youtube.com|youtu.be)\/(watch)?(\?v=)?(\S+)?`)
	s = regexp.MustCompile(`(http:|https:)?\/\/(www\.)?(open.spotify.com|spotify.com)\/(watch)?(\?v=)?(\S+)?`)
)

type url struct {
	Url string
}
type nsong struct {
	Nsong int
}
type delete struct {
	Delete int
}
