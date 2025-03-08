package resty

import (
	"os"
	"os/exec"

	"github.com/gdamore/tcell/v2"
)

func (r Resty) initRequestBodyTextArea() {
	textarea := r.requestBodyTextArea.
		SetPlaceholder("CTRL+E to open editor...")

	textarea.
		SetBorder(true).
		SetTitle("Request Body").
		SetInputCapture(r.requestBodyTextAreaInputCapture)

	r.components[1] = textarea
}

func (r Resty) requestBodyTextAreaInputCapture(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyCtrlE {
		f, err := os.CreateTemp("", "body")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		defer os.Remove(f.Name())

		if _, err := f.WriteString(r.requestBodyTextArea.GetText()); err != nil {
			panic(err)
		}

		r.app.Suspend(func() {
			cmd := exec.Command("nvim", f.Name())
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

		r.requestBodyTextArea.SetText(string(b), false)
		return nil
	}

	return event
}
