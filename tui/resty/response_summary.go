package resty

import (
	"strings"
	"fmt"
	"maps"
	"slices"

	"github.com/saeedafzal/resty/model"
)

func (r Resty) initResponseSummaryTextView() {
	textview := r.responseSummaryTextView.
		SetDynamicColors(true)

	textview.
		SetBorder(true).
		SetTitle("Response Summary")

	r.components[3] = textview
}

func (r Resty) updateResponseSummaryError(err error) {
	var doc strings.Builder
	doc.WriteString("[red::bu]ERROR[-:-:-]\n")
	doc.WriteString(fmt.Sprintf("API call failed: %s", err))

	r.responseSummaryTextView.SetText(doc.String())
}

func (r Resty) updateResponseSummary(res *model.ResponseData) {
	colour := "-"
	if res.StatusCode >= 500 {
		colour = "red"
	} else if res.StatusCode >= 400 {
		colour = "yellow"
	} else if res.StatusCode >= 200 {
		colour = "green"
	}

	var doc strings.Builder
	doc.WriteString(fmt.Sprintf("Status Code:   [%s::b]%d[-:-:-]\n", colour, res.StatusCode))
	doc.WriteString(fmt.Sprintf("Response Time: [::b]%d[-:-:-]ms\n", res.Time))
	doc.WriteString("\n[yellow::bu]Headers[-:-:-]\n")

	keys := slices.Sorted(maps.Keys(res.Headers))
	for i := range keys {
		key := keys[i]
		doc.WriteString(fmt.Sprintf("%s: %s\n", key, res.Headers.Get(key)))
	}

	r.responseSummaryTextView.SetText(doc.String())
}
