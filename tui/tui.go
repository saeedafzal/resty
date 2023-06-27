package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/model"
	"github.com/saeedafzal/resty/tui/request"
	"github.com/saeedafzal/resty/tui/response"
	"github.com/saeedafzal/resty/util"
)

type TUI struct {
	model *model.Model
}

func NewTUI(model *model.Model) {
	t := TUI{model}
	t.init()
}

// NOTE: Initialise the UI, this is the start point.
func (t TUI) init() {
	m := t.model
	p := m.Pages

	requestPanel := request.NewPanel(m)
	responsePanel := response.NewPanel(m)

	p.
		AddPage(util.BasePage, t.rootLayout(requestPanel, responsePanel), true, true).
		AddPage(util.AddHeaderDialogPage, requestPanel.AddHeaderDialog(), true, false)

	p.SetInputCapture(t.pageInputCapture)
}

// NOTE: Keybindings for the base page component. These are the global keybindings
// on the highest widget in the widget tree.
func (t TUI) pageInputCapture(event *tcell.EventKey) *tcell.EventKey {
	m := t.model

	if event.Rune() == 'q' && !m.IsInputField() && !m.IsDialog() {
		m.App.Stop()
		return nil
	}

	return event
}

// NOTE: Root layout.
func (t TUI) rootLayout(requestPanel request.Panel, responsePanel response.Panel) *tview.Flex {
	flex := tview.NewFlex().
		AddItem(requestPanel.Root(), 0, 1, true).
		AddItem(responsePanel.Root(), 0, 1, false)

	flex.SetInputCapture(t.rootInputCapture)
	return flex
}

// NOTE: Keybindings for the root flex layout.
func (t TUI) rootInputCapture(event *tcell.EventKey) *tcell.EventKey {
	m := t.model

	switch event.Key() {
	case tcell.KeyCtrlJ:
		i := m.GetCurrentIndex() + 1
		if i >= len(m.Components) {
			i = 0
		}
		m.App.SetFocus(m.Components[i])
		return nil
	case tcell.KeyCtrlK:
		i := m.GetCurrentIndex() - 1
		if i < 0 {
			i = len(m.Components) - 1
		}
		m.App.SetFocus(m.Components[i])
		return nil
	}

	return event
}
