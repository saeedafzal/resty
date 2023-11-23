package response

import (
	"fmt"
	"sort"
	"strings"

	"github.com/saeedafzal/resty/model"
)

func (p Panel) initResponseSummaryTextView() {
	p.model.UpdateResponseSummary = p.updateResponseSummary

	p.responseSummaryTextView.
		SetBorder(true).
		SetTitle("Response Summary")

	p.model.Components[3] = p.responseSummaryTextView
}

func (p Panel) updateResponseSummary(res model.ResponseData, err error) {
	if err != nil {
		doc := strings.Builder{}
		doc.WriteString("[red::bu]ERROR[-:-:-]\n")
		doc.WriteString(fmt.Sprintf("API call failed: %s", err))
		p.responseSummaryTextView.SetText(doc.String())
		return
	}

	p.updateResponseSummaryTextView(res)
	p.updateResponseBodyTextView(res)
}

func (p Panel) updateResponseSummaryTextView(res model.ResponseData) {
	p.responseSummaryTextView.Clear()

	// Set the colour of the response status value
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

	// Display response headers
	keys := make([]string, len(res.Headers))
	i := 0
	for k := range res.Headers {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for i := range keys {
		k := keys[i]
		doc.WriteString(fmt.Sprintf("%s: %s\n", k, res.Headers.Get(k)))
	}

	p.responseSummaryTextView.SetText(doc.String())
}
