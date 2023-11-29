package resty

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/tui/dialog"
	"github.com/saeedafzal/resty/tui/helper"
)

func (r *Resty) addEditHeadersDialog(t, name, value string) tview.Primitive {
	form := tview.NewForm().
		AddInputField("Name", name, 0, nil, func(text string) { name = text }).
		AddInputField("Value", value, 0, nil, func(text string) { value = text }).
		AddButton(t+" Header", func() {
			r.addHeaderBtnHandler(t, name, value)
		}).
		AddButton("Cancel", func() {
			r.pages.RemovePage("ADD_EDIT_HEADERS_DIALOG")
		})

	form.
		SetBorder(true).
		SetTitle(t + " Header").
		SetInputCapture(r.addEditHeadersDialogInputHandler)

	return dialog.NewDialog(form, 40, 9)
}

func (r *Resty) addEditHeadersDialogInputHandler(event *tcell.EventKey) *tcell.EventKey {
	if 'q' == event.Rune() && !helper.IsInput(r.app) {
		r.pages.RemovePage("ADD_EDIT_HEADERS_DIALOG")
		return nil
	}

	return event
}

func (r *Resty) addHeaderBtnHandler(t, name, value string) {
	if name != "" || value != "" {
		r.requestData.Headers[name] = []string{value}

		if t == "Add" {
			r.addHeaderToTable(name, value)
		} else if t == "Edit" {
			r.editHeaderInTable(name, value)
		}

		r.updateRequestSummary()
		r.pages.RemovePage("ADD_EDIT_HEADERS_DIALOG")
	}
}
