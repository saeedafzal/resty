package tui

import "github.com/rivo/tview"

var (
	responseSummaryTextView = tview.NewTextView()
	responseBodyTextView    = tview.NewTextView()
)

func (u UI) responseView() *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(u.responseSummaryTextView(), 0, 1, false).
		AddItem(u.responseBodyTextView(), 0, 2, false)
}

func (_ UI) responseSummaryTextView() *tview.TextView {
	responseSummaryTextView.SetDynamicColors(true)
	responseSummaryTextView.SetBorder(true).SetTitle("Response Summary")
	return responseSummaryTextView
}

func (_ UI) responseBodyTextView() *tview.TextView {
	responseBodyTextView.SetBorder(true).SetTitle("Response Body")
	return responseBodyTextView
}
