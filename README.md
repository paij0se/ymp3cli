🛑this cli still has a lot of bugs🛑

<h1>A simple tool to download music from youtube/spotify and listen in the terminal.</h1>

<h1>Prerequisites</h1>

- golang
- youtube-dl https://github.com/ytdl-org/youtube-dl/releases/tag/2021.06.06
- python 3.6.1 or above (for youtubedl and spotdl)
- ffmpeg 4.2 or above (added to PATH)
- spotDL https://github.com/spotDL/spotify-downloader#installation

> **_YouTube Music must be available in your country for spotDL to work. This is because we use YouTube Music to filter search results. You can check if YouTube Music is available in your country, by visiting [YouTube Music](https://music.youtube.com)._**

## Installation

### Installing FFmpeg

- [Windows Tutorial](https://windowsloop.com/install-ffmpeg-windows-10/)
- OSX - `brew install ffmpeg`
- Linux - `sudo apt install ffmpeg`

<h1>Install instructions</h1>

- you can download the binary file from:  https://github.com/paij0se/ymp3cli/releases/tag/v0.0.5

<h1>Build instructions</h1>

- linux and macOS

for linux install the oto dependencies `sudo apt install libasound2-dev`,
for macOS Oto requies `AudioToolbox.framework`, but this is automatically linked.

run `./install.sh` to install the dependencies.

Build ymp3cli with `go build -o ymp3cli src/main.go`

and for execute ymp3cli just run `./ymp3cli`.

<img src="https://you-can.ml/monda/yessir.png">

<img src="https://you-can.ml/monda/ymp3cli.png">

<h1>TODO:</h1>

- [x] client in golang
- [ ] download the videos without youtube-dl
- [ ] works correctly in windows (Not tested yet)❓
- [ ] able to pause and rewind the songs
- [x] able to play a song one by one

<h1>Custom clients?</h1>
- visit the wiki https://github.com/paij0se/ymp3cli/wiki/Routes

<h1>Alternative clients</h1>
- The old deno client https://github.com/bruh-boys/ymp3cli-old-client
