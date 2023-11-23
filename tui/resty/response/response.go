package response

import (
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/model"
)

type Panel struct {
	model                   *model.Model
	responseSummaryTextView *tview.TextView
	responseBodyTextView    *tview.TextView
}

func NewPanel(model *model.Model) Panel {
	return Panel{
		model,
		tview.NewTextView().SetDynamicColors(true),
		tview.NewTextView().SetDynamicColors(true),
	}
}

func (p Panel) Root() *tview.Flex {
	p.initResponseSummaryTextView()
	p.initResponseBodyTextView()

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(p.responseSummaryTextView, 0, 1, false).
		AddItem(p.responseBodyTextView, 0, 2, false)
}
