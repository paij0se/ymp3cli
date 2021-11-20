import * as peo from "https://denopkg.com/iamnathanj/cursor@v2.2.0/mod.ts";
// clear screen
await peo.clearScreen();

while (true) {
  const input = prompt("youtube url:");

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
