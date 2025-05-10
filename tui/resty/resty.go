package resty

import (
	"github.com/gdamore/tcell/v2"
	"github.com/saeedafzal/resty/model"
	"github.com/saeedafzal/tview"
)

type Resty struct {
	app                     *tview.Application
	pages                   *tview.Pages
	components              []tview.Primitive
	index                   int
	requestBodyTextArea     *tview.TextArea
	requestSummaryTextView  *tview.TextView
	requestHeaderTable      *tview.Table
	responseSummaryTextView *tview.TextView
	responseBodyTextView    *tview.TextView
	requestData             *model.RequestData
}

func New(app *tview.Application, pages *tview.Pages) Resty {
	return Resty{
		app,
		pages,
		make([]tview.Primitive, 5),
		0,
		tview.NewTextArea(),
		tview.NewTextView(),
		tview.NewTable(),
		tview.NewTextView(),
		tview.NewTextView(),
		model.NewRequestData(),
	}
}

func (r Resty) Root() *tview.Flex {
	r.initRequestBodyTextArea()
	r.initResponseSummaryTextView()
	r.initResponseBodyTextView()

	request := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(r.requestForm(), 9, 0, true).
		AddItem(r.requestBodyTextArea, 0, 1, false).
		AddItem(r.requestSummary(), 0, 1, false)

	response := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(r.responseSummaryTextView, 0, 1, false).
		AddItem(r.responseBodyTextView, 0, 2, false)

	flex := tview.NewFlex().
		AddItem(request, 0, 1, true).
		AddItem(response, 0, 1, false)

	flex.SetInputCapture(r.rootInputCapture)

	return flex
}

func (r *Resty) rootInputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlJ:
		r.index = (r.index + 1) % len(r.components)
		r.app.SetFocus(r.components[r.index])
		return nil
	case tcell.KeyCtrlK:
		length := len(r.components)
		r.index = (r.index - 1 + length) % length
		r.app.SetFocus(r.components[r.index])
		return nil
	}

	return event
}
