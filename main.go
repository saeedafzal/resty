package main

import (
	"flag"
	"fmt"

	"github.com/saeedafzal/resty/model"
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

	// Initialise ui
	m := model.NewModel()
	t := tui.NewTUI(m)

	m.App.SetRoot(t.Root(), true)

	// Run application
	if err := m.App.Run(); err != nil {
		panic(err)
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
	fmt.Printf("Version:     %s\n", version)
	fmt.Printf("Commit:      %s\n", commit)
	fmt.Printf("Build Time:  %s\n", buildTime)
}
