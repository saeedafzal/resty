package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	rootPage         = "root"
	addHeaderPage    = "add_header_dialog"
	updateHeaderPage = "update_header_dialog"
)

type TUI struct {
	model Model

	requestPanel  RequestPanel
	responsePanel ResponsePanel
}

func NewTUI(m Model) TUI {
	resPanel := NewResponsePanel(m)
	reqPanel := NewRequestPanel(m, resPanel)

	return TUI{
		model:         m,
		requestPanel:  reqPanel,
		responsePanel: resPanel,
	}
}

func (t TUI) Root() *tview.Pages {
	p := t.model.pages

	p.AddPage(rootPage, t.baseLayout(), true, true).
		AddPage(addHeaderPage, t.requestPanel.addHeaderDialog(), true, false)

	p.SetInputCapture(t.pageInputCapture)
	return p
}

func (t TUI) baseLayout() *tview.Flex {
	return tview.NewFlex().
		AddItem(t.requestPanel.Root(), 0, 1, true).
		AddItem(t.responsePanel.Root(), 0, 1, false)
}

// ---

func (t TUI) pageInputCapture(event *tcell.EventKey) *tcell.EventKey {
	m := t.model

	switch event.Key() {
	case tcell.KeyCtrlJ:
		m := t.model
		i := m.getCurrentIndex() + 1
		if i >= len(m.components) {
			i = 0
		}
		m.App.SetFocus(m.components[i])
		return nil
	case tcell.KeyCtrlK:
		m := t.model
		i := m.getCurrentIndex() - 1
		if i < 0 {
			i = len(m.components) - 1
		}
		m.App.SetFocus(m.components[i])
		return nil
	}

	if event.Rune() == 'q' && !m.isInputField() && !m.isDialog() {
		m.App.Stop()
		return nil
	}

	return event
}
