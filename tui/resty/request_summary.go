package resty

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

func (r *Resty) requestSummary() *tview.TextView {
	r.updateRequestSummary()

	textview := r.requestSummaryTextView

	textview.
		SetBorder(true).
		SetTitle("Request Summary Meh")

	r.components[2] = textview
	return textview
}

func (r *Resty) updateRequestSummary() {
	req := r.requestData

	doc := strings.Builder{}
	doc.WriteString("Method: " + req.Method + "\n")
	doc.WriteString("URL:    " + req.Url + "\n\n")

	doc.WriteString("[yellow:-:b]Headers[-:-:-]")
	for k, v := range req.Headers {
		doc.WriteString(fmt.Sprintf("\n%s: %s", k, v[0]))
	}

	// Set the text of the textview
	textview := r.requestSummaryTextView
	textview.SetText(doc.String())
}
