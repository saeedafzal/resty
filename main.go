package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/tui"
)

var (
	version   string
	commit    string
	buildTime string
)

func main() {
	// Version flag
	if flags() {
		printVersion()
		return
	}

	// Setup TUI
	app := tview.NewApplication().
		EnableMouse(true)

	// Start application
	if err := app.SetRoot(tui.NewTUI(app), true).Run(); err != nil {
		log.Panicln(err)
	}
}

func flags() bool {
	var v bool
	flag.BoolVar(&v, "v", false, "Display application version.")
	flag.BoolVar(&v, "version", false, "Display application version.")
	flag.Parse()
	return v
}

func printVersion() {
	fmt.Println("\n=== Resty ===")
	fmt.Printf("Version:    %s\n", version)
	fmt.Printf("Commit:     %s\n", commit)
	fmt.Printf("Build Time: %s\n", buildTime)
}
