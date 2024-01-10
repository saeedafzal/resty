package dialog

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/helper"
)

// NewDialog creates a modal view on the center of the screen
// with the specified width and height.
func NewDialog(p tview.Primitive, width, height int) *tview.Flex {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

// InputTwoDialog displays a dialog with two inputs.
func InputTwoDialog(pages *tview.Pages, label, name, value string, callback func(name, value string)) tview.Primitive {
	form := tview.NewForm().
		SetButtonBackgroundColor(tcell.ColorIndigo).
		SetFieldBackgroundColor(tcell.ColorIndigo).
		AddInputField("Name", name, 0, nil, func(text string) { name = text }).
		AddInputField("Value", value, 0, nil, func(text string) { value = text }).
		AddButton(label, func() { callback(name, value) }).
		AddButton("Cancel", func() {
			pages.RemovePage(helper.FORM_DIALOG)
		})

	form.
		SetBorder(true).
		SetTitle(label).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 'q' || event.Key() == tcell.KeyESC {
				pages.RemovePage(helper.FORM_DIALOG)
				return nil
			}
			return event
		})

	return NewDialog(form, 40, 9)
}
