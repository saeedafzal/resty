package tui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/app/model"
)

func (t TUI) responsePanel() *tview.Flex {
	t.responseSummaryTextView.SetBorder(true).SetTitle("Response Summary")
	t.responseBodyTextView.SetBorder(true).SetTitle("Response Body")

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(t.responseSummaryTextView, 0, 1, false).
		AddItem(t.responseBodyTextView, 0, 2, false)
	return flex
}

// ---

func (t TUI) updateResponseSummaryTextView(res model.ResponseData) {
	t.responseSummaryTextView.SetText("")

	colour := "-"
	if res.StatusCode >= 500 {
		colour = "red"
	} else if res.StatusCode >= 400 {
		colour = "yellow"
	} else if res.StatusCode >= 200 {
		colour = "green"
	}

	doc := strings.Builder{}
	doc.WriteString(fmt.Sprintf("Status Code:   [%s::b]%d[-:-:-]\n", colour, res.StatusCode))
	doc.WriteString(fmt.Sprintf("Response Time: [::b]%d[-:-:-]ms\n", res.Time))
	doc.WriteString("\n[yellow::bu]Headers[-:-:-]\n")

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

	t.responseSummaryTextView.SetText(doc.String())
	t.responseSummaryTextView.ScrollToBeginning()
}

func (t TUI) updateResponseBodyTextView(contentType string, body string) {
	t.responseBodyTextView.SetText("")
	w := tview.ANSIWriter(t.responseBodyTextView)

	if err := quick.Highlight(w, body, getLangFromContentType(contentType), "terminal16m", "monokai"); err != nil {
		// TODO: Error modal? Default colour?
	}

	t.responseBodyTextView.ScrollToBeginning()
}

// ---

func getLangFromContentType(contentType string) string {
	if strings.Contains(contentType, "html") {
		return "html"
	} else if strings.Contains(contentType, "json") {
		return "json"
	} else {
		return "plain"
	}
}
