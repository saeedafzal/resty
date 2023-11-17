package request

import (
	"fmt"
	"strings"
)

func (p Panel) initRequestSummaryTextView() {
	p.model.UpdateRequestSummary = p.updateRequestSummary

	p.requestSummaryTextView.
		SetBorder(true).
		SetTitle("Request Summary")

	p.updateRequestSummary()
	p.model.Components[2] = p.requestSummaryTextView
}

func (p Panel) updateRequestSummary() {
	req := p.model.RequestData

	doc := strings.Builder{}
	doc.WriteString("Method: " + req.Method + "\n")
	doc.WriteString("URL:    " + req.Url + "\n\n")
	doc.WriteString("[yellow:-:b]Headers[-:-:-]")

	for k, v := range req.Headers {
		doc.WriteString(fmt.Sprintf("\n%s: %s", k, v[0]))
	}

	p.requestSummaryTextView.SetText(doc.String())
}
