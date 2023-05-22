package main

import (
	"fmt"

	"github.com/saeedafzal/resty/tui"
	"github.com/spf13/pflag"
)

var (
	version   string
	commit    string
	buildTime string
)

func main() {
	v := pflag.Bool("version", false, "Display application version.")
	pflag.Parse()

	if *v {
		fmt.Println("=== Resty ===")
		fmt.Printf("Version:     %s\n", version)
		fmt.Printf("Commit:      %s\n", commit)
		fmt.Printf("Build Time:  %s\n", buildTime)
		return
	}

	m := tui.NewModel()
	t := tui.NewTUI(m)
	if err := m.App.SetRoot(t.Root(), true).Run(); err != nil {
		panic(err)
	}
}
