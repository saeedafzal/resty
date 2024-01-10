package resty

import (
	"fmt"
	"sort"
	"strings"

	"github.com/saeedafzal/resty/model"
)

func (r Resty) initResponseSummary() {
	textview := r.responseSummary.
		SetDynamicColors(true).
		SetText("...")

	textview.
		SetBorder(true).
		SetTitle("Response Summary")

	r.components[3] = textview
}

func (r Resty) updateResponseSummaryError(err error) {
	doc := strings.Builder{}
	doc.WriteString("[red::bu]ERROR[-:-:-]\n")
	doc.WriteString(fmt.Sprintf("API call failed: %s", err))
	r.responseSummary.SetText(doc.String())
}

func (r Resty) updateResponseSummary(res *model.ResponseData) {
	// Set the colour of the response status
	colour := "-"
	if res.StatusCode >= 500 {
		colour = "red"
	} else if res.StatusCode >= 400 {
		colour = "yellow"
	} else if res.StatusCode >= 200 {
		colour = "green"
	}

	// Build the response summary text
	doc := strings.Builder{}
	doc.WriteString(fmt.Sprintf("Status Code:   [%s::b]%d[-:-:-]\n", colour, res.StatusCode))
	doc.WriteString(fmt.Sprintf("Response Time: [::b]%d[-:-:-]ms\n", res.Time))
	doc.WriteString("\n[yellow::bu]Headers[-:-:-]\n")

	// Response headers
	keys := make([]string, 0, len(res.Headers))
	for k := range res.Headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := range keys {
		k := keys[i]
		doc.WriteString(fmt.Sprintf("%s: %s\n", k, res.Headers.Get(k)))
	}

	r.responseSummary.SetText(doc.String())
}
