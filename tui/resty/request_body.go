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
		r.app.Suspend(r.startEditor)
		return nil
	}

	return event
}

func (r *Resty) startEditor() {
	// Get current text in text area
	textarea := r.requestBodyTextArea
	text := textarea.GetText()

	// TODO: Set extension on temp file?
	temp := "tmp"
	if err := r.writeTextToFile(text, temp); err != nil {
		// TODO: Error feedback
		return
	}
	defer os.Remove(temp)

	// Get terminal editor or use vim as default
	editor, exists := os.LookupEnv("EDITOR")
	if !exists {
		editor = "vim"
	}

	// Open editor
	cmd := exec.Command(editor, temp)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		// TODO: Error feedback
		return
	}

	editedText, err := r.readTextFromFile(temp)
	if err != nil {
		// TODO: Error feedback
		return
	}

	textarea.SetText(editedText, false)
}

func (_ *Resty) writeTextToFile(text, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	return err
}

func (_ *Resty) readTextFromFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
