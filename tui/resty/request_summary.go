package resty

import (
	"strings"

	"github.com/rivo/tview"
)

func (r Resty) requestSummary() *tview.Flex {
	r.initRequestHeadersTable()

	textview := r.requestSummaryTextView.
		SetDynamicColors(true)

	table := r.requestHeaderTable
	table.
		SetFocusFunc(func() { table.SetSelectable(true, false) }).
		SetBlurFunc(func() { table.SetSelectable(false, false) })

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(textview, 4, 0, false).
		AddItem(table, 0, 1, true)

	flex.
		SetBorder(true).
		SetTitle("Request Summary")

	r.components[2] = r.requestHeaderTable
	return flex
}

func (r Resty) updateRequestSummary() {
	data := r.requestData

	b := strings.Builder{}
	b.WriteString("Method: [-:-:b]" + data.Method + "[-:-:-]\n")
	b.WriteString("URL:    [-:-:b]" + data.Url + "[-:-:-]\n\n")
	b.WriteString("[yellow:-:b]Headers[-:-:-]")

	textview := r.requestSummaryTextView
	textview.Clear()
	textview.SetText(b.String())

	table := r.requestHeaderTable.Clear()
	index := 0
	for k, v := range data.Headers {
		table.SetCellSimple(index, 0, k)
		table.SetCellSimple(index, 1, strings.Join(v, ","))
		index += 1
	}
}
