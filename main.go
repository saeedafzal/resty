package main

import (
	"fmt"
	"log"

	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/tui"
	"github.com/spf13/pflag"
)

var (
	version,
	commit,
	buildTime string
)

func main() {
	// Flags
	versionFlag, helpFlag := InitFlags()

	if versionFlag {
		printVersion()
		return
	}

	if helpFlag {
		pflag.Usage()
		return
	}

	// Setup TUI
	app := tview.
		NewApplication().
		EnableMouse(true)

	// Start application
	if err := app.SetRoot(tui.NewTUI(app), true).Run(); err != nil {
		log.Panicln(err)
	}
}

func InitFlags() (bool, bool) {
	var versionFlag, helpFlag bool

	pflag.BoolVarP(&versionFlag, "version", "v", false, "Display application version.")
	pflag.BoolVarP(&helpFlag, "help", "h", false, "Help for Resty.")

	pflag.Parse()
	return versionFlag, helpFlag
}

func printVersion() {
	fmt.Println("\n=== Resty ===")
	fmt.Printf("Version:    %s\n", version)
	fmt.Printf("Commit:     %s\n", commit)
	fmt.Printf("Build Time: %s\n", buildTime)
}
