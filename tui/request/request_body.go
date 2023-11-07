package request

import (
	"os"
	"os/exec"

	"github.com/gdamore/tcell/v2"
)

func (r Panel) requestBodyInit() {
	r.requestBodyTextArea.
		SetBorder(true).
		SetTitle("Request Body").
		SetInputCapture(r.responseBodyInputCapture)

	r.model.Components[1] = r.requestBodyTextArea
}

func (r Panel) responseBodyInputCapture(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyCtrlE {
		r.model.App.Suspend(r.startEditor)
		return nil
	}

	return event
}

func (r Panel) startEditor() {
	textarea := r.requestBodyTextArea
	text := textarea.GetText()

	// TODO: Dynamically set extension? Only if content type has been set.
	// TODO: Set to actual temp directory
	tempFile := "tmp.json"

	if err := r.writeTextToFile(text, tempFile); err != nil {
		// TODO: Display error feedback
		return
	}
	defer os.Remove(tempFile)

	// NOTE: Get terminal editor
	editor, exists := os.LookupEnv("EDITOR")
	if !exists {
		editor = "vim"
	}

	cmd := exec.Command(editor, tempFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		// TODO: Display error feedback
		return
	}

	editedText, err := r.readTextFromFile(tempFile)
	if err != nil {
		// TODO: Display error feedback
		return
	}

	textarea.SetText(editedText, false)
}

func (_ Panel) writeTextToFile(text, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	return err
}

func (_ Panel) readTextFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
