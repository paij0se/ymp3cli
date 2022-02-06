package cli

import "fmt"

func HelpCommand() {
	fmt.Println(`   
  Usage: ymp3cli -[OPTION]
  -h ,-help: Display the help command
  -v ,-version: Display the version of ymp3cli
  -p ,-play: Play a single song
  -u ,-update: Update ymp3cli to the latest version
  -d ,-download ,-download: Download a song from youtube
  -port: Port to run the server on

  Usage: ymp3cli -p [SONG]
  ymp3cli -p <song.mp3>: play a single song
  example: ymp3cli -p song.mp3
 
  Usage: ymp3cli -d [Link]
  ymp3cli -d <link>: download a song from youtube
  example: ymp3cli -d https://www.youtube.com/watch?v=dQw4w9WgXcQ
 
	 MIT License
	 Made it by pai
	 https://paijose.cf
 
 `)
}

/*
	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&help, "h", false, "show help")

	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.BoolVar(&showVersion, "v", false, "show version")

	flag.StringVar(&url, "d", "", "download a song from youtube ")
	flag.StringVar(&url, "download", "", "download a song from youtube ")

	flag.StringVar(&song, "p", "", "play a single song")
	flag.StringVar(&song, "play", "", "play a single song")

	flag.IntVar(&port, "port", 0, "port to run the server on")
*/
