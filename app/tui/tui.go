package tui

import (
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/app/api"
	"github.com/saeedafzal/resty/app/model"
)

const (
	basePage       = "basePage"
	addHeadersPage = "addHeadersPage"
)

type TUI struct {
	app   *tview.Application
	pages *tview.Pages

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
		app:   app,
		pages: tview.NewPages(),

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
	return t.pages
}

func (t TUI) layout() *tview.Flex {
	flex := tview.NewFlex().
		AddItem(t.requestPanel(), 0, 1, true).
		AddItem(t.responsePanel(), 0, 1, false)
	return flex
}
