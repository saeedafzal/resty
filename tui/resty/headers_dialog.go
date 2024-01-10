package resty

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/helper"
	"github.com/saeedafzal/resty/tui/dialog"
)

func (r Resty) headersDialog() tview.Primitive {
	table := r.headersTable

	table.
		SetFixed(0, 0).
		SetSelectable(true, false)

	table.
		SetBorder(true).
		SetTitle("Request Headers").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 'q' || event.Key() == tcell.KeyESC {
				r.pages.RemovePage(helper.HEADER_DIALOG)
				return nil
			}

			switch event.Rune() {
			case 'a':
				r.pages.AddPage(
					helper.FORM_DIALOG,
					dialog.InputTwoDialog(r.pages, "Add Header", "", "", r.addHeaderCallback),
					true,
					true,
				)
				return nil
			case 'd':
				i, _ := table.GetSelection()
				cell := table.GetCell(i, 0)
				delete(r.requestData.Headers, cell.Text)
				r.headersTable.RemoveRow(i)
				r.updateRequestSummary()
				return nil
			case 'e':
				i, _ := table.GetSelection()
				name := table.GetCell(i, 0).Text
				value := table.GetCell(i, 1).Text
				r.pages.AddPage(
					helper.FORM_DIALOG,
					dialog.InputTwoDialog(r.pages, "Edit Header", name, value, r.editHeaderCallback),
					true,
					true,
				)
				return nil
			}

			return event
		})

	r.updateHeadersTable()
	return dialog.NewDialog(table, 0, 0)
}

func (r Resty) updateHeadersTable() {
	table := r.headersTable.Clear()

	table.
		SetCell(0, 0, helper.HeaderCell("Name")).
		SetCell(0, 1, helper.HeaderCell("Value"))

	index := 1
	for k, v := range r.requestData.Headers {
		table.
			SetCell(index, 0, helper.Cell(k)).
			SetCell(index, 1, helper.Cell(v[0]))
		index = index + 1
	}
}

func (r Resty) addHeaderCallback(name, value string) {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(value) == "" {
		return
	}

	r.requestData.Headers[name] = []string{value}

	table := r.headersTable
	i := table.GetRowCount()
	table.
		SetCell(i, 0, helper.Cell(name)).
		SetCell(i, 1, helper.Cell(value))

	r.updateRequestSummary()
	r.pages.RemovePage(helper.FORM_DIALOG)
}

func (r Resty) editHeaderCallback(name, value string) {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(value) == "" {
		return
	}

	r.requestData.Headers[name] = []string{value}

	table := r.headersTable
	index, _ := table.GetSelection()
	table.GetCell(index, 0).SetText(name)
	table.GetCell(index, 1).SetText(value)

	r.updateRequestSummary()
	r.pages.RemovePage(helper.FORM_DIALOG)
}
