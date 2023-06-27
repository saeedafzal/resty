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
	var v bool
	flag.BoolVar(&v, "version", false, "Display application version.")
	flag.Parse()

	if v {
		fmt.Println("\n=== Resty ===")
		fmt.Printf("Version:     %s\n", version)
		fmt.Printf("Commit:      %s\n", commit)
		fmt.Printf("Build Time:  %s\n", buildTime)
		return
	}

	m := model.NewModel()
	tui.NewTUI(m)

	m.App.SetRoot(m.Pages, true)
	if err := m.App.Run(); err != nil {
		panic(err)
	}
}
