package resty

import (
	"os"
	"os/exec"

	"github.com/gdamore/tcell/v2"
)

func (r Resty) initRequestBody() {
	textarea := r.requestBody.
		SetPlaceholder("Enter body here...")

	textarea.
		SetBorder(true).
		SetTitle("Request Body").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyCtrlE {
				r.startExternalEditor()
				return nil
			}

			return event
		})

	r.components[1] = textarea
}

func (r Resty) startExternalEditor() {
	// Create temp file to open
	f, err := os.CreateTemp(os.TempDir(), "*.tmp")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer os.Remove(f.Name())

	// Get request body and write into file
	currentText := r.requestBody.GetText()
	if _, err := f.WriteString(currentText); err != nil {
		panic(err)
	}

	// Get editor
	editor, ok := os.LookupEnv("EDITOR")
	if !ok {
		editor = "vim"
	}

	// Suspend application and launch editor
	r.app.Suspend(func() {
		cmd := exec.Command(editor, f.Name())
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	})

	// FIXME: https://github.com/rivo/tview/issues/932
	nscr, _ := tcell.NewScreen()
	r.app.SetScreen(nscr)

	// Read file
	b, err := os.ReadFile(f.Name())
	if err != nil {
		panic(err)
	}

	r.requestBody.SetText(string(b), false)
}
