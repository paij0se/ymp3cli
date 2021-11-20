package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	header := widgets.NewParagraph()
	header.Text = "Press q to quit, Press a or d to switch tabs"
	header.SetRect(0, 0, 50, 1)
	header.Border = false
	header.TextStyle.Bg = ui.ColorMagenta
	// help part
	p2 := widgets.NewParagraph()
	p2.Text = "Press q to quit\nPress a or d to switch tabs\n"
	p2.Title = "Keys"
	p2.SetRect(5, 5, 40, 15)
	p2.BorderStyle.Fg = ui.ColorYellow

	// songs part
	bc := widgets.NewParagraph()
	// request songs
	postBody, _ := json.Marshal(map[string]string{
		"url": "",
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:8000/download", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	bc.Text = sb
	bc.Title = "Songs"
	bc.SetRect(5, 5, 40, 15)
	bc.BorderStyle.Fg = ui.ColorCyan

	tabpane := widgets.NewTabPane("help", "songs")
	tabpane.SetRect(0, 1, 50, 4)
	tabpane.Border = true

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			ui.Render(p2)
		case 1:
			ui.Render(bc)
		}
	}

	ui.Render(header, tabpane, p2)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "a":
			tabpane.FocusLeft()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		case "d":
			tabpane.FocusRight()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		}
	}
}
