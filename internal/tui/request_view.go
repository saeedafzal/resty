package tui

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/rivo/tview"
)

var (
	methods                = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}
	requestSummaryTextView = tview.NewTextView()
	headersTable           = tview.NewTable()
)

func (u UI) requestView() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(u.formView(), 0, 1, true).
		AddItem(u.requestSummaryView(), 0, 1, false)

	return flex
}

func (u UI) formView() *tview.Form {
	form := tview.NewForm().
		AddDropDown("Method", methods, 0, func(option string, _ int) {
			u.request.Method = option
			u.formatRequestSummaryView()
		}).
		AddInputField("URL", "", 0, nil, func(text string) {
			u.request.Url = text
			u.formatRequestSummaryView()
		}).
		AddButton("Add Headers", func() {
			u.pages.ShowPage(addHeaderDialog)
			u.app.SetFocus(addHeaderForm)
		}).
		AddButton("Send", func() {
			u.pages.ShowPage(sendingDialogPage)
			u.doRequest()
			u.pages.HidePage(sendingDialogPage)
		})

	form.SetBorder(true).SetTitle("Request Form")
	return form
}

func (u UI) requestSummaryView() *tview.Flex {
	requestSummaryTextView.
		SetDynamicColors(true)
	u.formatRequestSummaryView()

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(requestSummaryTextView, 0, 1, false).
		AddItem(u.requestHeadersView(), 0, 4, false)

	flex.SetBorder(true).SetTitle("Request Summary")

	return flex
}

func (u UI) requestHeadersView() *tview.Flex {
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetText("[yellow:-:b]Headers[-:-:-]")
	textView.SetBorderPadding(0, 0, 1, 1)

	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(textView, 1, 0, false).
		AddItem(u.headersTable(), 0, 1, false)
}

func (u UI) formatRequestSummaryView() {
	doc := strings.Builder{}
	doc.WriteString(fmt.Sprintf("[-:-:b]Method[-:-:-]: %s\n", u.request.Method))
	doc.WriteString(fmt.Sprintf("[-:-:b]URL[-:-:-]:    %s", u.request.Url))

	requestSummaryTextView.SetText(doc.String())
}

func (_ UI) sendingDialog() *tview.Modal {
	modal := tview.NewModal().SetText("Sending HTTP request...")
	modal.SetBorder(true).SetTitle("Sending")
	return modal
}

func (_ UI) headersTable() *tview.Table {
	headersTable.SetSelectable(true, true).
		SetBorders(true).
		SetFixed(1, 1).
		SetCell(0, 0, tableCell("Name")).
		SetCell(0, 1, tableCell("Value"))

	return headersTable
}

func (u UI) updateHeadersTable() {
	headersTable.Clear()
	index := 1

	headersTable.SetCell(0, 0, tableCell("Name")).
		SetCell(0, 1, tableCell("Value"))

	for n := range u.request.Headers {
		headersTable.SetCell(index, 0, selectableTableCell(n)).
			SetCell(index, 1, selectableTableCell(u.request.Headers.Get(n)))
	}
}

// TODO: Move to different file?
func tableCell(text string) *tview.TableCell {
	return tview.NewTableCell(text).
		SetSelectable(false).
		SetExpansion(1)
}

func selectableTableCell(text string) *tview.TableCell {
	return tview.NewTableCell(text).
		SetSelectable(true).
		SetExpansion(1)
}
