package dialog

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/model"
)

type HeaderDialog struct {
	model *model.Model
	table *tview.Table
}

func NewHeaderDialog(model *model.Model) HeaderDialog {
	return HeaderDialog{
		model,
		tview.NewTable(),
	}
}

func (h HeaderDialog) Root() *tview.Grid {
	h.initTable()
	return dialog(h.table, 0, 0)
}

// Draws the initial view of the table.
func (h HeaderDialog) initTable() {
	h.table.
		SetFixed(0, 0).
		SetCell(0, 0, h.titleCell("Name")).
		SetCell(0, 1, h.titleCell("Value")).
		SetSelectable(true, false)

	headers := h.model.RequestData.Headers
	for k, v := range headers {
		h.addNewHeader(k, v[0])
	}

	h.table.
		SetBorder(true).
		SetTitle("Request Headers").
		SetInputCapture(h.tableInputCapture)
}

func (_ HeaderDialog) titleCell(text string) *tview.TableCell {
	return tview.NewTableCell(text).
		SetExpansion(1).
		SetAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorYellow).
		SetSelectable(false)
}

func (_ HeaderDialog) cell(text string) *tview.TableCell {
	return tview.NewTableCell(text).
		SetExpansion(1).
		SetAlign(tview.AlignCenter).
		SetSelectable(true)
}

func (h HeaderDialog) tableInputCapture(event *tcell.EventKey) *tcell.EventKey {
	pages := h.model.Pages

	switch event.Rune() {
	case 'a':
		d := AddHeaderDialog(h.model, "Add Header", h.addHeaderCallback)
		pages.AddAndSwitchToPage("HEADER_FORM_DIALOG", d.Root(), true)
		return nil
	case 'd':
		h.deleteHeader()
		return nil
	case 'e':
		rowIndex, _ := h.table.GetSelection()
		name := h.table.GetCell(rowIndex, 0).Text
		value := h.table.GetCell(rowIndex, 1).Text
		d := EditHeaderDialog(h.model, "Edit Header", name, value, h.editHeaderCallback)
		pages.AddAndSwitchToPage("HEADER_FORM_DIALOG", d.Root(), true)
		return nil
	}

	// Destroy page on exit
	if event.Key() == tcell.KeyESC || event.Rune() == 'q' {
		pages.RemovePage("HEADERS_DIALOG")
	}

	return event
}

func (h HeaderDialog) addNewHeader(name, value string) {
	t := h.table
	row := t.GetRowCount()

	t.
		SetCell(row, 0, h.cell(name)).
		SetCell(row, 1, h.cell(value))
}

func (h HeaderDialog) addHeaderCallback(name, value string) {
	if name == "" || value == "" {
		return
	}

	m := h.model
	m.RequestData.Headers[name] = []string{value}
	h.model.UpdateRequestSummary()

	t := h.table
	row := t.GetRowCount()

	t.
		SetCell(row, 0, h.cell(name)).
		SetCell(row, 1, h.cell(value))

	m.Pages.RemovePage("HEADER_FORM_DIALOG")
}

// Deletes the currently selected header.
func (h HeaderDialog) deleteHeader() {
	rowIndex, _ := h.table.GetSelection()
	nameCell := h.table.GetCell(rowIndex, 0)

	delete(h.model.RequestData.Headers, nameCell.Text)
	h.model.UpdateRequestSummary()

	h.table.RemoveRow(rowIndex)
}

func (h HeaderDialog) editHeaderCallback(name, value string) {
	if name == "" || value == "" {
		return
	}

	m := h.model
	m.RequestData.Headers[name] = []string{value}
	h.model.UpdateRequestSummary()

	t := h.table
	rowIndex, _ := t.GetSelection()

	t.GetCell(rowIndex, 0).SetText(name)
	t.GetCell(rowIndex, 1).SetText(value)

	m.Pages.RemovePage("HEADER_FORM_DIALOG")
}
