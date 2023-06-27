package request

import (
	"strings"

	"github.com/rivo/tview"
)

// NOTE: Base flex layout for the request summary section.
func (r Panel) requestSummaryFlex() *tview.Flex {
	r.initRequestHeadersTable()

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(r.requestSummaryTextView, 4, 1, false).
		AddItem(r.requestHeadersTable, 0, 1, false)

	flex.
		SetBorder(true).
		SetTitle("Request Summary")

	return flex
}

// NOTE: Update the content within the request summary from request data.
func (r Panel) updateRequestSummaryTextView() {
	req := r.model.RequestData

	doc := strings.Builder{}
	doc.WriteString("Method: " + req.Method + "\n")
	doc.WriteString("URL:    " + req.Url + "\n\n")
	doc.WriteString("[yellow:-:b]Headers[-:-:-]")

	r.requestSummaryTextView.SetText(doc.String())
}
