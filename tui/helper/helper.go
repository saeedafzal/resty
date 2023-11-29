package helper

import (
	"reflect"
	"strings"

	"github.com/rivo/tview"
)

func IsInput(app *tview.Application) bool {
	focusType := reflect.TypeOf(app.GetFocus()).String()
	return focusType == "*tview.InputField" || focusType == "*tview.TextArea"
}

func IsDialog(pages *tview.Pages) bool {
	name, _ := pages.GetFrontPage()
	return strings.Contains(name, "DIALOG")
}
