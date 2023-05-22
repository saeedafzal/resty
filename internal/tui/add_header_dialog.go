package tui

import "github.com/rivo/tview"

var addHeaderForm = tview.NewForm()

func (u UI) addHeaderDialog() *tview.Flex {
	width, height := 40, 10

	var name, value string

	addHeaderForm.
		AddInputField("Name", "", 0, nil, func(text string) {
			name = text
		}).
		AddInputField("Value", "", 0, nil, func(text string) {
			value = text
		}).
		AddButton("Add Header", func() {
			u.request.Headers.Add(name, value)
			u.updateHeadersTable()
		}).
		AddButton("Cancel", func() {
			u.pages.HidePage(addHeaderDialog)
		})
	addHeaderForm.SetBorder(true).SetTitle("Add Headers")

	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(addHeaderForm, height, 1, false).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}
