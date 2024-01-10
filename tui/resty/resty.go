package resty

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/model"
)

type Resty struct {
	app   *tview.Application
	pages *tview.Pages

	requestSummary  *tview.TextView
	headersTable    *tview.Table
	requestBody     *tview.TextArea
	responseSummary *tview.TextView
	responseBody    *tview.TextView

	components []tview.Primitive
	index      int

	requestData *model.RequestData
}

func New(app *tview.Application, pages *tview.Pages) Resty {
	return Resty{
		app,
		pages,

		tview.NewTextView(),
		tview.NewTable(),
		tview.NewTextArea(),
		tview.NewTextView(),
		tview.NewTextView(),

		make([]tview.Primitive, 5),
		0,

		model.NewRequestData(),
	}
}

func (r Resty) Root() *tview.Flex {
	flex := tview.NewFlex().
		AddItem(r.request(), 0, 1, true).
		AddItem(r.response(), 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlJ:
			r.index = r.index + 1
			if r.index == len(r.components) {
				r.index = 0
			}
			r.app.SetFocus(r.components[r.index])
			return nil
		case tcell.KeyCtrlK:
			r.index = r.index - 1
			if r.index < 0 {
				r.index = len(r.components) - 1
			}
			r.app.SetFocus(r.components[r.index])
			return nil
		}

		return event
	})

	return flex
}

func (r Resty) request() *tview.Flex {
	r.initRequestSummary()
	r.initRequestBody()

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(r.requestForm(), 9, 1, true).
		AddItem(r.requestBody, 0, 1, false).
		AddItem(r.requestSummary, 0, 1, false)
}

func (r Resty) response() *tview.Flex {
	r.initResponseSummary()
	r.initResponseBody()

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(r.responseSummary, 0, 1, false).
		AddItem(r.responseBody, 0, 2, false)
}
