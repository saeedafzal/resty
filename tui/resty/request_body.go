package resty

import (
	"github.com/saeedafzal/resty/core"

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
		text := core.OpenEditor(r.app, r.requestBodyTextArea.GetText(), false)
		r.requestBodyTextArea.SetText(text, false)
		return nil
	}

	return event
}
