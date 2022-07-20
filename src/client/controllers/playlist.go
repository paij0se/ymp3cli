package controllers

func Playlist(url string) []string {
	go PlaySoundAll(url, "")
	return GetSongs(url)
}
