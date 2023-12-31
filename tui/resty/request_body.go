package resty

import (
	"os"
	"os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (r *Resty) requestBody() *tview.TextArea {
	textarea := r.requestBodyTextArea

	textarea.
		SetBorder(true).
		SetTitle("Request Body").
		SetInputCapture(r.requestBodyInputCapture)

	r.components[1] = textarea
	return textarea
}

func (r *Resty) requestBodyInputCapture(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyCtrlE {
		r.externalEditor()
		return nil
	}

	return event
}

func (r *Resty) externalEditor() {
	file, err := os.CreateTemp("", "temp")
	if err != nil {
		panic(err)
	}
	defer r.closeFile(file)

	textarea := r.requestBodyTextArea
	currentText := textarea.GetText()
	if _, err := file.WriteString(currentText); err != nil {
		panic(err)
	}

	editor, exists := os.LookupEnv("EDITOR")
	if !exists {
		editor = "vim"
	}

	r.app.Suspend(func() {
		cmd := exec.Command(editor, file.Name())
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			panic(err)
		}
	})

	b, err := os.ReadFile(file.Name())
	if err != nil {
		panic(err)
	}

	textarea.SetText(string(b), false)
}

func (r *Resty) closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		panic(err)
	}
	if err := os.Remove(file.Name()); err != nil {
		panic(err)
	}
}
