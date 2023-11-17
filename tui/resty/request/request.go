package request

import (
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/api"
	"github.com/saeedafzal/resty/model"
)

type Panel struct {
	model                  *model.Model
	api                    api.API
	requestSummaryTextView *tview.TextView
	requestBodyTextArea    *tview.TextArea
}

func NewPanel(model *model.Model) Panel {
	return Panel{
		model,
		api.NewAPI(),
		tview.NewTextView().SetDynamicColors(true),
		tview.NewTextArea(),
	}
}

func (p Panel) Root() *tview.Flex {
	p.initRequestBodyTextArea()
	p.initRequestSummaryTextView()

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(p.requestForm(), 10, 1, true).
		AddItem(p.requestBodyTextArea, 0, 1, false).
		AddItem(p.requestSummaryTextView, 0, 1, false)
}
