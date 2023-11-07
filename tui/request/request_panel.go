package request

import (
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/api"
	"github.com/saeedafzal/resty/model"
)

type Panel struct {
	model *model.Model

	requestSummaryTextView *tview.TextView
	requestHeadersTable    *tview.Table
	requestBodyTextArea    *tview.TextArea

	api             api.API
	editHeaderValue string
}

func NewPanel(model *model.Model) Panel {
	return Panel{
		model:                  model,
		requestSummaryTextView: tview.NewTextView().SetDynamicColors(true),
		requestHeadersTable:    tview.NewTable(),
		requestBodyTextArea:    tview.NewTextArea(),

		api:             api.NewAPI(),
		editHeaderValue: "",
	}
}

// Root builds the main layout for the [Panel].
func (r Panel) Root() *tview.Flex {
	r.requestBodyInit()

	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(r.requestForm(), 10, 1, true).
		AddItem(r.requestBodyTextArea, 0, 1, false).
		AddItem(r.requestSummaryFlex(), 0, 1, false)
}
