package tui

import "github.com/rivo/tview"

type responsePanel struct {
	Model
}

func NewResponsePanel(model Model) responsePanel {
	return responsePanel{model}
}

func (t responsePanel) Root() *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(t.responseSummaryTextView(), 0, 1, false).
		AddItem(t.responseBodyTextView(), 0, 1, false)
}

func (t responsePanel) responseSummaryTextView() *tview.TextView {
	textView := tview.NewTextView().SetDynamicColors(true)

	textView.SetBorder(true).SetTitle("Response Summary")
	t.Model.components[3] = textView
	return textView
}

func (t responsePanel) responseBodyTextView() *tview.TextView {
	textView := tview.NewTextView().SetDynamicColors(true)

	textView.SetBorder(true).SetTitle("Response Body")
	t.Model.components[4] = textView
	return textView
}
