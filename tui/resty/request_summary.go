package resty

import (
	"strings"
)

func (r Resty) initRequestSummary() {
	textview := r.requestSummary.
		SetDynamicColors(true)

	textview.
		SetBorder(true).
		SetTitle("Request Summary")

	r.updateRequestSummary()
}

func (r Resty) updateRequestSummary() {
	req := r.requestData

	b := strings.Builder{}
	b.WriteString("Method: [-:-:b]" + req.Method + "[-:-:-]\n")
	b.WriteString("URL:    [-:-:b]" + req.Url + "[-:-:-]\n\n")
	b.WriteString("[yellow:-:b]Headers:[-:-:-]")

	for k, v := range req.Headers {
		b.WriteString("\n" + k + ": " + strings.Join(v, ","))
	}

	r.requestSummary.Clear()
	r.requestSummary.SetText(b.String())

	r.components[2] = r.requestSummary
}
