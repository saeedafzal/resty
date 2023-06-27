package request

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/tui/dialog"
	"github.com/saeedafzal/resty/util"
)

// AddHeaderDialog creates the form dialog used to add new header values.
// TODO: Move inline functions out?
func (r Panel) AddHeaderDialog() tview.Primitive {
	m := r.model
	name, value := "", ""

	form := tview.NewForm().
		AddInputField("Name", value, 0, nil, func(text string) {
			name = text
		}).
		AddInputField("Value", value, 0, nil, func(text string) {
			value = text
		}).
		AddButton("Add Header", func() {
			m.RequestData.Headers[name] = []string{value}
			r.updateRequestHeadersTable()
		}).
		AddButton("Cancel", func() {
			m.Pages.HidePage(util.AddHeaderDialogPage)
		}).
		SetFieldBackgroundColor(util.HexToColour("#6F00FF")).
		SetButtonBackgroundColor(util.HexToColour("#6F00FF"))

	form.
		SetBorder(true).
		SetTitle("Add New Header").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			return r.dialogInputCapture(event, util.AddHeaderDialogPage)
		})

	return dialog.Dialog(form)
}

// NOTE: Creates the form dialog with single input field to update
// a header value. It is pre-populated with the value being edited.
func (r Panel) showEditDialogPage() {
	pages := r.model.Pages
	v := ""

	form := tview.NewForm().
		AddInputField("Value", r.editHeaderValue, 0, nil, func(text string) {
			v = text
		}).
		AddButton("Edit Header", func() {
			r.editRequestHeadersHandler(v)
		}).
		AddButton("Cancel", func() {
			pages.RemovePage(util.EditHeaderDialogPage)
		}).
		SetFieldBackgroundColor(util.HexToColour("#6F00FF")).
		SetButtonBackgroundColor(util.HexToColour("#6F00FF"))

	form.
		SetBorder(true).
		SetTitle("Edit Header Value").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			return r.dialogInputCapture(event, util.EditHeaderDialogPage)
		})

	pages.AddPage(util.EditHeaderDialogPage, dialog.Dialog(form), true, true)
}

// NOTE: Default input capture for dialog components.
// TODO: Should it be moved to dialog.go?
func (r Panel) dialogInputCapture(event *tcell.EventKey, page string) *tcell.EventKey {
	m := r.model

	if event.Key() == tcell.KeyEsc || (event.Rune() == 'q' && !m.IsInputField()) {
		m.Pages.HidePage(page)
		return nil
	}

	return event
}

// NOTE: Handler for editing header value.
func (r Panel) editRequestHeadersHandler(value string) {
	table := r.requestHeadersTable
	headers := r.model.RequestData.Headers

	row, col := table.GetSelection()

	if col == 0 { // Header key
		v := headers[r.editHeaderValue]
		delete(headers, r.editHeaderValue)
		headers[value] = v
	} else { // Value
		cell := table.GetCell(row, 0)
		headers[cell.Text] = []string{value}
	}

	r.updateRequestHeadersTable()
	r.model.Pages.RemovePage(util.EditHeaderDialogPage)
}
