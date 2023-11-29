package resty

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/api"
	"github.com/saeedafzal/resty/model"
)

type Resty struct {
	app         *tview.Application
	pages       *tview.Pages
	requestData model.RequestData
	api         api.API

	requestSummaryTextView *tview.TextView
	requestHeadersTable    *tview.Table
	requestBodyTextArea    *tview.TextArea

	responseSummaryTextView *tview.TextView
	responseBodyTextView    *tview.TextView

	components []tview.Primitive
}

func NewResty(app *tview.Application, pages *tview.Pages) *tview.Flex {
	r := Resty{
		app,
		pages,
		model.NewRequestData(),
		api.NewAPI(),

		tview.NewTextView().SetDynamicColors(true),
		tview.NewTable(),
		tview.NewTextArea(),

		tview.NewTextView().SetDynamicColors(true),
		tview.NewTextView().SetDynamicColors(true),

		make([]tview.Primitive, 5),
	}

	return r.Root()
}

func (r *Resty) Root() *tview.Flex {
	flex := tview.NewFlex().
		AddItem(r.requestPanel(), 0, 1, true).
		AddItem(r.responsePanel(), 0, 1, false)

	flex.SetInputCapture(r.restyInputCapture)

	return flex
}

func (r *Resty) requestPanel() *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(r.requestForm(), 9, 1, true).
		AddItem(r.requestBody(), 0, 1, false).
		AddItem(r.requestSummary(), 0, 1, false)
}

func (r *Resty) responsePanel() *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(r.responseSummary(), 0, 1, false).
		AddItem(r.responseBody(), 0, 2, false)
}

func (r *Resty) restyInputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlJ:
		i := r.getCurrentIndex() + 1
		if i >= len(r.components) {
			i = 0
		}
		r.app.SetFocus(r.components[i])
		return nil
	case tcell.KeyCtrlK:
		i := r.getCurrentIndex() - 1
		if i < 0 {
			i = len(r.components) - 1
		}
		r.app.SetFocus(r.components[i])
		return nil
	}
	return event
}

func (r *Resty) getCurrentIndex() int {
	for i, v := range r.components {
		if v.HasFocus() {
			return i
		}
	}
	return 0
}
