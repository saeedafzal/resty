package main

import (
	"flag"
	"fmt"

	"github.com/saeedafzal/resty/tui"
	"github.com/saeedafzal/tview"
)

var version string

func main() {
	if flags() {
		return
	}

	app := tview.NewApplication().
		EnableMouse(true).
		EnablePaste(true)
	tui := tui.New(app)

	if err := app.SetRoot(tui.Root(), true).Run(); err != nil {
		panic(err)
	}
}

func flags() bool {
	v := flag.Bool("version", false, "Show application version.")
	h := flag.Bool("help", false, "Usage of resty.")
	flag.Parse()

	if *v {
		fmt.Println("Resty:", version)
		return true
	}

	if *h {
		flag.Usage()
		return true
	}

	return false
}
