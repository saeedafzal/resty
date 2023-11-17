package request

import (
	"os"
	"os/exec"

	"github.com/gdamore/tcell/v2"
)

func (p Panel) initRequestBodyTextArea() {
	p.requestBodyTextArea.
		SetBorder(true).
		SetTitle("Request Body").
		SetInputCapture(p.responseBodyInputCapture)

	p.model.Components[1] = p.requestBodyTextArea
}

func (p Panel) responseBodyInputCapture(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyCtrlE {
		p.model.App.Suspend(p.startEditor)
		return nil
	}

	return event
}

func (p Panel) startEditor() {
	textarea := p.requestBodyTextArea
	text := textarea.GetText()

	// TODO: Dynamically set extension? Only if content type has been set?
	tempFile := "tmp.json"

	if err := p.writeTextToFile(text, tempFile); err != nil {
		// TODO: Display error feedback
		return
	}
	defer os.Remove(tempFile)

	// Get terminal editor if it exists
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

	editedText, err := p.readTextFromFile(tempFile)
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
