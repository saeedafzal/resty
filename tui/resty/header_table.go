package resty

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/saeedafzal/resty/tui/modal"
)

func (r Resty) initRequestHeadersTable() {
	table := r.requestHeaderTable
	table.
		SetFocusFunc(func() { table.SetSelectable(true, false) }).
		SetBlurFunc(func() { table.SetSelectable(false, false) }).
		SetInputCapture(r.requestHeaderTableInputCapture)
}

func (r Resty) requestHeaderTableInputCapture(event *tcell.EventKey) *tcell.EventKey {
	table := r.requestHeaderTable

	switch event.Rune() {
	case 'a':
		r.pages.AddPage(
			"header",
			modal.HeaderModal(r.pages, "Add Header", "", "", r.addHeaderCallback),
			true,
			true,
		)
		return nil
	case 'd':
		row, _ := table.GetSelection()
		cell := table.GetCell(row, 0)
		table.RemoveRow(row)
		delete(r.requestData.Headers, cell.Text)
		return nil
	case 'e':
		row, _ := table.GetSelection()
		name := table.GetCell(row, 0).Text
		value := table.GetCell(row, 1).Text
		r.pages.AddPage(
			"header",
			modal.HeaderModal(r.pages, "Edit Header", name, value, r.editHeaderCallback),
			true,
			true,
		)
		return nil
	}

	return event
}

// Used by request form add header button
func (r Resty) addHeaderAddSelected(key, value string) {
	if strings.TrimSpace(key) == "" || strings.TrimSpace(value) == "" {
		return
	}

	r.requestData.Headers[key] = []string{value}
	// TODO: Add to table directly instead of re-rendering entire summary
	r.updateRequestSummary()
	r.pages.RemovePage("header")
}

func (r Resty) addHeaderCallback(name, value string) {
	r.addHeaderAddSelected(name, value)
	r.app.SetFocus(r.requestHeaderTable)
}

func (r Resty) editHeaderCallback(name, value string) {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(value) == "" {
		return
	}

	r.requestData.Headers[name] = []string{value}

	table := r.requestHeaderTable
	index, _ := table.GetSelection()
	table.GetCell(index, 0).SetText(name)
	table.GetCell(index, 1).SetText(value)

	r.pages.RemovePage("header")
	r.app.SetFocus(table)
}
