package model

import (
	"reflect"
	"strings"

	"github.com/rivo/tview"
)

type Model struct {
	App        *tview.Application
	Pages      *tview.Pages
	Components []tview.Primitive

	RequestData          RequestData
	UpdateResponsePanels func(res ResponseData)
}

func NewModel() *Model {
	return &Model{
		App:         tview.NewApplication().EnableMouse(true),
		Pages:       tview.NewPages(),
		Components:  make([]tview.Primitive, 5),
		RequestData: NewRequestData(),
	}
}

func (m Model) GetCurrentIndex() int {
	for i, v := range m.Components {
		if v.HasFocus() {
			return i
		}
	}

	return 0
}

func (m Model) IsInputField() bool {
	focusType := reflect.TypeOf(m.App.GetFocus()).String()
	return focusType == "*tview.InputField" || focusType == "*tview.TextArea"
}

func (m Model) IsDialog() bool {
	name, _ := m.Pages.GetFrontPage()
	return strings.Contains(strings.ToLower(name), "dialog")
}
