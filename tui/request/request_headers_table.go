package request

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// NOTE: Initialises the table widget for displays request headers.
func (r Panel) initRequestHeadersTable() {
	table := r.requestHeadersTable.
		SetBorders(true).
		SetFixed(0, 0).
		SetSelectable(false, false)

	table.
		SetFocusFunc(func() { table.SetSelectable(true, true) }).
		SetBlurFunc(func() { table.SetSelectable(false, false) }).
		SetInputCapture(r.requestHeadersInputCapture)

	r.updateRequestHeadersTable()
	r.model.Components[2] = table
}

// NOTE: Creates a new table cell for headers.
func (r Panel) headerCell(name string) *tview.TableCell {
	return tview.NewTableCell(name).
		SetExpansion(1).
		SetSelectable(false).
		SetTextColor(tcell.ColorYellow)
}

// NOTE: Creates a new table cell for normal content.
func (r Panel) cell(name string) *tview.TableCell {
	return tview.NewTableCell(name).
		SetExpansion(1).
		SetSelectable(true).
		SetMaxWidth(1)
}

// NOTE: Update the contents of the request headers table
func (r Panel) updateRequestHeadersTable() {
	r.requestHeadersTable.
		Clear().
		SetCell(0, 0, r.headerCell("Name")).
		SetCell(0, 1, r.headerCell("Value"))

	i := 1
	for k, v := range r.model.RequestData.Headers {
		r.requestHeadersTable.
			SetCell(i, 0, r.cell(k)).
			SetCell(i, 1, r.cell(v[0]))
		i++
	}
}

func (r Panel) requestHeadersInputCapture(event *tcell.EventKey) *tcell.EventKey {
	table := r.requestHeadersTable
	headers := r.model.RequestData.Headers

	switch event.Rune() {
	case 'd':
		row, _ := table.GetSelection()
		cell := table.GetCell(row, 0)
		delete(headers, cell.Text)
		table.RemoveRow(row)
		return nil
	case 'e':
		row, col := table.GetSelection()
		cell := table.GetCell(row, col)
		r.editHeaderValue = cell.Text
		r.showEditDialogPage()
		return nil
	}

	return event
}
