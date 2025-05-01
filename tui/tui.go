package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/saeedafzal/tview"
	"github.com/saeedafzal/resty/tui/resty"
)

type Tui struct {
	app   *tview.Application
	pages *tview.Pages
}

func New(app *tview.Application) Tui {
	return Tui{app, tview.NewPages()}
}

func (t Tui) Root() *tview.Pages {
	resty := resty.New(t.app, t.pages)

	t.pages.
		AddPage("resty", resty.Root(), true, true)

	t.pages.SetInputCapture(t.rootInputCapture)

	return t.pages
}

func (t Tui) rootInputCapture(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'q' {
		t.app.Stop()
		return nil
	}

	return event
}
