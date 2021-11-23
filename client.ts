import * as peo from "https://denopkg.com/iamnathanj/cursor@v2.2.0/mod.ts";
import * as ink from "https://deno.land/x/ink@1.3/mod.ts";
import { emptyDir } from "https://deno.land/std/fs/mod.ts";
import Ask from "https://deno.land/x/ask@1.0.6/mod.ts";
const url = "http://127.0.0.1:8000/";
// clear screen
await peo.clearScreen();
console.log(
  (await Deno.readTextFile("welcome.txt")) + "\n",
  "welcome to ymp3cli\n",
  "Type <ctrl + c> to exit.\n",
  "version 0.0.1\n"
);
while (true) {
  // get the available songs and display them with the option to play one by one
  const avaiableSongs = await fetch(`${url}songs`, {
    method: "GET",
  });
  const songs = await avaiableSongs.text();
  console.log(ink.colorize(`<magenta>avaiable songs:\n${songs}</magenta>`));

  const ask = new Ask(); // global options are also supported! (see below)

  const nsongs = await ask.prompt([
    {
      name: "n",
      type: "number",
      message: "type the number of the song you want to listen (if there is no songs just type s + enter to skip):",
    },
  ]);
  const play = nsongs.n;

  const nSong = await fetch(`${url}y`, {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      nsong: play,
    }),
  });
  const content1 = await nSong.text();
  console.log(content1);

  const input = prompt(ink.colorize("<green>? youtube url for download the song (you can type s + enter for skip):</green>"));
  console.log(ink.colorize(`<yellow>downloading: ${input}</yellow>`));
  const del = prompt(
    ink.colorize("<red>delete previous songs? (y:yes, n:no):</red>")
  );
  if (del === "y") {
    // delete previous songs
    emptyDir("./music");
  }
  const rawResponse = await fetch(`${url}download`, {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      url: input,
    }),
  });
  const content = await rawResponse.json();
  console.log(content);
}
