package dialog

import (
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/model"
)

type HeaderFormDialog struct {
	model       *model.Model
	title       string
	name, value string
	callback    func(name, value string)
}

func AddHeaderDialog(model *model.Model, title string, callback func(name, value string)) HeaderFormDialog {
	return HeaderFormDialog{
		model,
		title,
		"",
		"",
		callback,
	}
}

func EditHeaderDialog(model *model.Model, title, name, value string, callback func(name, value string)) HeaderFormDialog {
	return HeaderFormDialog{
		model,
		title,
		name,
		value,
		callback,
	}
}

func (a HeaderFormDialog) Root() *tview.Grid {
	return dialog(a.addHeaderForm(), 40, 10)
}

func (a HeaderFormDialog) addHeaderForm() *tview.Form {
	form := tview.NewForm().
		AddInputField("Name", a.name, 0, nil, func(text string) { a.name = text }).
		AddInputField("Value", a.value, 0, nil, func(text string) { a.value = text }).
		AddButton(a.title, func() { a.callback(a.name, a.value) }).
		AddButton("Cancel", a.destroyDialog)

	form.
		SetBorder(true).
		SetTitle(a.title)

	return form
}

func (a HeaderFormDialog) destroyDialog() {
	a.model.Pages.RemovePage("HEADER_FORM_DIALOG")
}
