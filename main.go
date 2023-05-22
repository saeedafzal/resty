package main

import (
	"flag"
	"fmt"

	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/internal/tui"
)

var (
	version   = "0.0.1"
	buildTime string
)

func main() {
	displayVersion := flag.Bool("version", false, "Display application version.")
	flag.Parse()

	if *displayVersion {
		fmt.Println("RESTY")
		fmt.Println(fmt.Sprintf("Version:    %s", version))
		fmt.Println(fmt.Sprintf("Build Time: %s", buildTime))
		return
	}

	app := tview.NewApplication().EnableMouse(true)
	ui := tui.NewUI(app)

	if err := app.SetRoot(ui.Layout(), true).Run(); err != nil {
		panic(err)
	}
}
