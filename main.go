package main

import (
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/tui"
)

func main() {
	app := tview.NewApplication().EnableMouse(true)
	ui := tui.NewUI(app)

	if err := app.SetRoot(ui.Layout(), true).Run(); err != nil {
		panic(err)
	}
}
