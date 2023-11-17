package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/model"
	"github.com/saeedafzal/resty/tui/resty"
	"reflect"
	"strings"
)

// TUI is the root layout of the entire application.
type TUI struct {
	model *model.Model
}

func NewTUI(model *model.Model) TUI {
	return TUI{model}
}

func (t TUI) Root() *tview.Pages {
	// Pages
	r := resty.NewResty(t.model)

	pages := t.model.Pages
	pages.AddPage("RESTY_PAGE", r.Root(), true, true)

	pages.SetInputCapture(t.tuiInputCapture)
	return pages
}

func (t TUI) tuiInputCapture(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'q' && !t.isInputField() && !t.IsDialog() {
		t.model.App.Stop()
	}

	return event
}

// Checks if an input field such as a [tview.TextArea] or
// [tview.InputField] currently has focus.
func (t TUI) isInputField() bool {
	focusType := reflect.TypeOf(t.model.App.GetFocus()).String()
	return focusType == "*tview.InputField" || focusType == "*tview.TextArea"
}

func (t TUI) IsDialog() bool {
	name, _ := t.model.Pages.GetFrontPage()
	return strings.Contains(name, "DIALOG")
}
