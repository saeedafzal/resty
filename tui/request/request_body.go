package request

import "github.com/rivo/tview"

func (r Panel) requestBodyTextView() *tview.TextArea {
	textArea := tview.NewTextArea()

	textArea.
		SetBorder(true).
		SetTitle("Request Body")

	r.model.Components[1] = textArea
	return textArea
}
