package tui

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/helper"
	"github.com/saeedafzal/resty/tui/resty"
)

type TUI struct {
	app   *tview.Application
	pages *tview.Pages
}

func New(app *tview.Application) TUI {
	return TUI{app, tview.NewPages()}
}

func (t TUI) Root() *tview.Pages {
	resty := resty.New(t.app, t.pages)

	t.pages.
		AddPage(helper.RESTY_PAGE, resty.Root(), true, true)

	t.pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' && !t.isDialog() {
			t.app.Stop()
			return nil
		}

		return event
	})

	return t.pages
}

func (t TUI) isDialog() bool {
	name, _ := t.pages.GetFrontPage()
	return strings.Contains(name, "DIALOG")
}
