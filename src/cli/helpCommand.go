package cli

import "fmt"

func HelpCommand() {
	fmt.Println(`   
  Usage: ymp3cli --[OPTION]
  --h,--help: Display the help command
  --v,--version: Display the version of ymp3cli
 
  Usage: ymp3cli [SONG]
  ymp3cli <song.mp3>: play a single song
  example: ymp3cli music/song.mp3
 
  Usage: ymp3cli [Link]
  ymp3cli <link>: download a song from youtube
  example: ymp3cli https://www.youtube.com/watch?v=dQw4w9WgXcQ
 
	 MIT License
	 Made it by pai
	 https://paijose.cf
 
 `)
}
