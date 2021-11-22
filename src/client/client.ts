import * as peo from "https://denopkg.com/iamnathanj/cursor@v2.2.0/mod.ts";
import * as ink from "https://deno.land/x/ink@1.3/mod.ts";
import { emptyDir } from "https://deno.land/std/fs/mod.ts";
import { showSongs, fileNames } from "./utils/showSongs.ts";
// clear screen
await peo.clearScreen();
console.log(
  (await Deno.readTextFile("welcome.txt")) + "\n",
  "welcome to ymp3cli\n",
  "Type <ctrl + c> to exit.\n",
  "version 0.0.1\n"
);
while (true) {
  console.log(ink.colorize(`<magenta>avaiable songs:\n${showSongs}</magenta>`));

  const input = prompt(ink.colorize("<green>youtube url:</green>"));
  console.log(ink.colorize(`<yellow>downloading: ${input}</yellow>`));
  const play = prompt(
    ink.colorize("<blue>play a random song? (y:yes, n:no):</blue>")
  );
  if (play === "y") {
    const randomSong = fileNames[Math.floor(Math.random() * fileNames.length)];
    console.log(ink.colorize(`<yellow>playing: ${randomSong} </yellow>`));
    // play a random song from the music folder
    const process = Deno.run({
      cmd: ["mpg321", "music/" + randomSong],
      stdout: "piped",
      stderr: "piped",
    });

    const output = await process.output();
    const outStr = new TextDecoder().decode(output);
    console.log(outStr);
  }
  const del = prompt(
    ink.colorize("<red>delete previous songs? (y:yes, n:no):</red>")
  );
  if (del === "y") {
    // delete previous songs
    emptyDir("./music");
  }
  const rawResponse = await fetch("http://127.0.0.1:8000/download", {
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
