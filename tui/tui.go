package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/tui/helper"
	"github.com/saeedafzal/resty/tui/resty"
)

// Base layout for the entire UI.
type TUI struct {
	app   *tview.Application
	pages *tview.Pages
}

func NewTUI(app *tview.Application) tview.Primitive {
	return TUI{
		app,
		tview.NewPages(),
	}.Root()
}

func (t TUI) Root() *tview.Pages {
	pages := t.pages

	pages.SetInputCapture(t.pagesInputCapture)

	return t.pages.
		AddPage("RESTY", resty.NewResty(t.app, t.pages), true, true)
}

func (t TUI) pagesInputCapture(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'q' && !helper.IsInput(t.app) && !helper.IsDialog(t.pages) {
		t.app.Stop()
		return nil
	}

	return event
}
