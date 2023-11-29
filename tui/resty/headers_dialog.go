package resty

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/tui/dialog"
)

func (r *Resty) headersDialog() tview.Primitive {
	table := r.requestHeadersTable.Clear()

	table.
		SetFixed(0, 0).
		SetSelectable(true, false).
		SetCell(0, 0, r.headerCell("Name")).
		SetCell(0, 1, r.headerCell("Value"))

	for k, v := range r.requestData.Headers {
		r.addHeaderToTable(k, v[0])
	}

	table.
		SetBorder(true).
		SetTitle("Request Headers").
		SetInputCapture(r.headersDialogInputHandler)

	return dialog.NewDialog(table, 0, 0)
}

func (_ *Resty) headerCell(text string) *tview.TableCell {
	return tview.NewTableCell(text).
		SetExpansion(1).
		SetAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorYellow).
		SetSelectable(false)
}

func (_ *Resty) cell(text string) *tview.TableCell {
	return tview.NewTableCell(text).
		SetExpansion(1).
		SetAlign(tview.AlignCenter).
		SetSelectable(true)
}

func (r *Resty) headersDialogInputHandler(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'q':
		r.pages.RemovePage("HEADERS_DIALOG")
		return nil
	case 'a':
		r.pages.AddPage("ADD_EDIT_HEADERS_DIALOG", r.addEditHeadersDialog("Add", "", ""), true, true)
		return nil
	case 'e':
		t := r.requestHeadersTable
		rowIndex, _ := t.GetSelection()
		name := t.GetCell(rowIndex, 0).Text
		value := t.GetCell(rowIndex, 1).Text
		r.pages.AddPage("ADD_EDIT_HEADERS_DIALOG", r.addEditHeadersDialog("Edit", name, value), true, true)
		return nil
	case 'd':
		r.deleteHeaderInTable()
		return nil
	}

	return event
}

func (r *Resty) addHeaderToTable(name, value string) {
	t := r.requestHeadersTable
	row := t.GetRowCount()

	t.
		SetCell(row, 0, r.cell(name)).
		SetCell(row, 1, r.cell(value))
}

func (r *Resty) editHeaderInTable(name, value string) {
	t := r.requestHeadersTable

	rowIndex, _ := t.GetSelection()
	t.GetCell(rowIndex, 0).SetText(name)
	t.GetCell(rowIndex, 1).SetText(value)
}

func (r *Resty) deleteHeaderInTable() {
	t := r.requestHeadersTable
	rowIndex, _ := t.GetSelection()

	delete(r.requestData.Headers, t.GetCell(rowIndex, 0).Text)
	t.RemoveRow(rowIndex)

	r.updateRequestSummary()
}
