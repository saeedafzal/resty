package tui

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/helper"
	"github.com/saeedafzal/resty/tui/environment"
	"github.com/saeedafzal/resty/tui/navigation"
	"github.com/saeedafzal/resty/tui/resty"
	"go.etcd.io/bbolt"
)

type TUI struct {
	app   *tview.Application
	pages *tview.Pages

	db *bbolt.DB
}

func New(app *tview.Application, db *bbolt.DB) TUI {
	return TUI{app, tview.NewPages(), db}
}

func (t TUI) Root() *tview.Pages {
	resty := resty.New(t.app, t.pages)
	navigation := navigation.New(t.pages)
	environment := environment.New(t.db, t.app, t.pages)

	t.pages.
		AddPage(helper.RESTY_PAGE, resty.Root(), true, true).
		AddPage(helper.NAVIGATION_PAGE, navigation.Root(), true, false).
		AddPage(helper.ENVIRONMENT_PAGE, environment.Root(), true, false)

	t.pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' && !t.isDialog() {
			t.app.Stop()
			return nil
		}

		if event.Key() == tcell.KeyCtrlN {
			t.pages.SwitchToPage(helper.NAVIGATION_PAGE)
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
