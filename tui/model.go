package tui

import (
	"reflect"
	"strings"

	"github.com/rivo/tview"
)

// The Model shared between all components.
type Model struct {
	App   *tview.Application
	pages *tview.Pages

	// The components that can be cycled between
	components []tview.Primitive
	index      int
}

func NewModel() Model {
	return Model{
		App:   tview.NewApplication().EnableMouse(true),
		pages: tview.NewPages(),

		components: make([]tview.Primitive, 5),
		index:      0,
	}
}

func (m Model) setComponentIndex(i int) tview.Primitive {
	m.index = i
	return m.components[i]
}

func (m Model) getCurrentIndex() int {
	for i, v := range m.components {
		if v.HasFocus() {
			return i
		}
	}
	return 0
}

func (m Model) isInputField() bool {
	focusType := reflect.TypeOf(m.App.GetFocus()).String()
	return focusType == "*tview.InputField" || focusType == "*tview.TextArea"
}

func (m Model) isDialog() bool {
	name, _ := m.pages.GetFrontPage()
	return strings.Contains(name, "dialog")
}
