export const fileNames: string[] = [];

for await (const dirEntry of Deno.readDir("music")) {
  if (dirEntry.isFile) {
    fileNames.push(dirEntry.name);
  }
}

export const showSongs: string = fileNames.join("\n");
