package main

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/tui"
	"github.com/spf13/pflag"
)

var version string

func main() {
	// Handle CLI flags and exit if necessary
	if flags() {
		return
	}

	// Start terminal ui
	app := tview.NewApplication()
	tui := tui.New(app)

	app.
		EnableMouse(true).
		SetRoot(tui.Root(), true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func flags() bool {
	var v, h bool
	pflag.BoolVarP(&v, "version", "v", false, "Display application version.")
	pflag.BoolVarP(&h, "help", "h", false, "Usage of Resty.")
	pflag.Parse()

	if v {
		fmt.Println("Resty", version)
		return true
	}

	if h {
		pflag.Usage()
		return true
	}

	return false
}
