import * as ink from "https://deno.land/x/ink@1.3/mod.ts";

export async function showSongs() {
  const cmd = Deno.run({
    cmd: ["ls", "-a", "music"],
    stderr: "piped",
    stdout: "piped",
    stdin: "null",
  });

  const rawOutput = await cmd.output();
  cmd.close();
  const output: string = new TextDecoder().decode(rawOutput); // here decode de output hex => text
  console.log(ink.colorize(`<magenta>${output}</magenta>`));
}
