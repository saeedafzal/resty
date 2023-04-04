package tui

import (
	"reflect"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/app/api"
	"github.com/saeedafzal/resty/app/model"
)

const (
	basePage       = "basePage"
	addHeadersPage = "addHeadersPage"
)

type TUI struct {
	app        *tview.Application
	pages      *tview.Pages
	components []tview.Primitive

	requestModel model.RequestModel
	api          api.API

	requestSummaryTextView *tview.TextView
	requestHeadersTable    *tview.Table
	requestBodyTextArea    *tview.TextArea

	responseSummaryTextView *tview.TextView
	responseBodyTextView    *tview.TextView
}

func NewTUI(app *tview.Application) TUI {
	return TUI{
		app:        app,
		pages:      tview.NewPages(),
		components: make([]tview.Primitive, 5),

		requestModel: model.NewRequestModel(),
		api:          api.NewAPI(),

		requestSummaryTextView: tview.NewTextView().SetDynamicColors(true),
		requestHeadersTable:    tview.NewTable(),
		requestBodyTextArea:    tview.NewTextArea(),

		responseSummaryTextView: tview.NewTextView().SetDynamicColors(true),
		responseBodyTextView:    tview.NewTextView().SetDynamicColors(true),
	}
}

func (t TUI) Pages() *tview.Pages {
	t.pages.
		AddPage(basePage, t.layout(), true, true).
		AddPage(addHeadersPage, t.addHeadersModal(), true, false)

	// TODO: move out?
	t.pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlK:
			for i, v := range t.components {
				if v.HasFocus() {
					index := i + 1
					if index >= len(t.components) {
						index = 0
					}
					c := t.components[index]
					t.app.SetFocus(c)
					return nil
				}
			}

			return event
		case tcell.KeyCtrlJ:
			for i, v := range t.components {
				if v.HasFocus() {
					index := i - 1
					if index < 0 {
						index = len(t.components) - 1
					}
					c := t.components[index]
					t.app.SetFocus(c)
					return nil
				}
			}

			return event
		}

		focusType := reflect.TypeOf(t.app.GetFocus()).String()
		if focusType == "*tview.InputField" || focusType == "*tview.TextArea" {
			return event
		}

		switch event.Rune() {
		case 'q':
			t.app.Stop()
			return nil
		}

		return event
	})

	return t.pages
}

func (t TUI) layout() *tview.Flex {
	flex := tview.NewFlex().
		AddItem(t.requestPanel(), 0, 1, true).
		AddItem(t.responsePanel(), 0, 1, false)
	return flex
}
