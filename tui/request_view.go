package tui

import (
	"net/http"

	"github.com/rivo/tview"
)

func (u UI) requestView() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(u.formView(), 0, 1, true).
		AddItem(u.requestHeadersTable(), 0, 1, false)

	return flex
}

func (u UI) formView() *tview.Form {
	form := tview.NewForm().
		AddDropDown("Method", []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}, 0, func(option string, _ int) {
			u.request.Method = option
		}).
		AddInputField("URL", "", 0, nil, func(text string) {
			u.request.Url = text
		}).
		AddButton("Add Headers", nil).
		AddButton("Send", func() {
			u.doRequest()
		})

	form.SetBorder(true).SetTitle("Request Form")
	return form
}

func (u UI) requestHeadersTable() *tview.Table {
	return u.createTable("Headers")
}
