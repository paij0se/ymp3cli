<h1>A simple terminal tool for downloading and listening music from YouTube, Spotify and SoundCloud </h1>
<p>consumes <10MB of ram!!!</p>
<div align=center>

<img src="https://media.discordapp.net/attachments/907631182240436305/950177164718915604/unknown.png?width=625&height=400"/>

<h1>Discord RPC</h1>

<img src="https://media.discordapp.net/attachments/957431274237218849/957681287458590810/unknown.png"/>
	
<h1>Change the playback speed<h1>

https://user-images.githubusercontent.com/69026987/153721642-be1a61c7-aa3c-4ec9-898b-37d73d994b17.mp4
	
</div>

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=paij0se/ymp3cli&type=Date)](https://star-history.com/#paij0se/ymp3cli&Date)


<h1>Why???</h1>
idk, is free.

<h1>Prerequisites</h1>

- golang 1.17+
- python 3.6.1 or above (for spotdl)
- ffmpeg 4.2 or above (added to PATH)
- spotDL https://github.com/spotDL/spotify-downloader#installation

### Installing FFmpeg

- [Windows Tutorial](https://windowsloop.com/install-ffmpeg-windows-10/)
- OSX - `brew install ffmpeg`
- Linux - `sudo apt install ffmpeg`

<h1>Install instructions</h1>

- You can download the binary file from:  https://github.com/paij0se/ymp3cli/releases

- Or with curl
```bash
curl https://raw.githubusercontent.com/paij0se/ymp3cli/main/install.sh | bash
```

- verify the installation with
```bash
$ ymp3cli --h

  Usage: ymp3cli -[OPTION]
  -h ,-help: Display the help command
  -v ,-version: Display the version of ymp3cli
  -p ,-play: Play a single song
  -u ,-update: Update ymp3cli to the latest version
  -d ,-download: Download a song from youtube
  -s ,-speed: That allows changing the playback speed
  -sd ,-soundcloud: Download a song from soundcloud

  Usage: ymp3cli -p [SONG]
  ymp3cli -p <song.mp3>: play a single song
  example: ymp3cli -p song.mp3

  Usage: ymp3cli -s [SONG]
  ymp3cli -s <song.mp3>: change the playback speed
  example: ymp3cli -s song.mp3

  Usage: ymp3cli -d [Link]
  ymp3cli -d <link>: download a song from youtube
  example: ymp3cli -d https://www.youtube.com/watch?v=dQw4w9WgXcQ

	 MIT License
	 Made it by pai
	 https://paijose.cf




$ ymp3cli # start ymp3cli

$ ymp3cli -d https://www.youtube.com/watch?v=dQw4w9WgXcQ # download a song from youtube

$ ymp3cli -p song.mp3 # play a single song

```

<h1>Build instructions</h1>

for linux install the oto dependencies

```bash
sudo apt install libasound2-dev
```
for macOS Oto requies `AudioToolbox.framework`, but this is automatically linked.

run `go get .` to install the dependencies.

Build ymp3cli with `go build -o ymp3cli src/main.go`

and for execute ymp3cli just run `./ymp3cli`.

<h1>Config</h1>

In /home/user/.ymp3cli/config.yaml

```yaml
# port: the port that the server will listen on
port: 8888
```

<h1>TODO:</h1>

- [x] client in golang
- [x] download the videos without youtube-dl
- [x] works correctly in windows
- [x] Discord rpc
- [x] able to pause and rewind the songs
- [ ] A playlists system
- [x] able to play a song one by one

<h1>Custom clients?</h1>
- visit the wiki https://github.com/paij0se/ymp3cli/wiki/Routes

<h1>Alternative clients</h1>
- The old deno client https://github.com/bruh-boys/ymp3cli-old-client

<h1>Thanks to</h1>
- Flames https://github.com/FlamesX-128
