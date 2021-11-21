import * as peo from "https://denopkg.com/iamnathanj/cursor@v2.2.0/mod.ts";
import * as ink from "https://deno.land/x/ink@1.3/mod.ts";
import { emptyDir } from "https://deno.land/std/fs/mod.ts";
import { showSongs } from "./utils/showSongs.ts";
// clear screen
await peo.clearScreen();

while (true) {
  showSongs();

  const input = prompt(ink.colorize("<green>youtube url:</green>"));
  const del = prompt(
    ink.colorize("<red>delete previous songs? (y:yes, n:no):</red>")
  );
  if (del === "y") {
    // delete previous songs
    emptyDir("./music");
  }
  console.log(ink.colorize("<blue>avaiable songs:</blue>"));
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
