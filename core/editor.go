package core

import (
	"os"
	"os/exec"

	"github.com/saeedafzal/tview"
)

// TODO: Need some kind of status bar etc to display errors on UI
// instead of panic
func OpenEditor(app *tview.Application, text string, readonly bool) string {
	f, err := os.CreateTemp("", "body")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer os.Remove(f.Name())

	if _, err := f.WriteString(text); err != nil {
		panic(err)
	}

	flags := ""
	if readonly {
		flags = "-M"
	}

	app.Suspend(func() {
		cmd := exec.Command("nvim", flags, f.Name())
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	})

	b, err := os.ReadFile(f.Name())
	if err != nil {
		panic(err)
	}

	return string(b)
}
