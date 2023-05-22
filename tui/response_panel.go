package tui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/model"
)

type ResponsePanel struct {
	model Model

	responseSummaryTextView *tview.TextView
	responseBodyTextView    *tview.TextView
}

func NewResponsePanel(model Model) ResponsePanel {
	return ResponsePanel{
		model: model,

		responseSummaryTextView: tview.NewTextView().SetDynamicColors(true),
		responseBodyTextView:    tview.NewTextView().SetDynamicColors(true),
	}
}

func (t ResponsePanel) Root() *tview.Flex {
	t.createResponseSummaryTextView()
	t.createResponseBodyTextView()

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(t.responseSummaryTextView, 0, 1, false).
		AddItem(t.responseBodyTextView, 0, 1, false)
}

func (t ResponsePanel) createResponseSummaryTextView() {
	t.responseSummaryTextView.SetBorder(true).SetTitle("Response Summary")
	t.model.components[3] = t.responseSummaryTextView
}

func (t ResponsePanel) createResponseBodyTextView() {
	t.responseBodyTextView.SetBorder(true).SetTitle("Response Body")
	t.model.components[4] = t.responseBodyTextView
}

// ---

func (t ResponsePanel) updateResponsePanels(res model.ResponseData) {
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

	t.responseBodyTextView.SetText("")
	w := tview.ANSIWriter(t.responseBodyTextView)

	var buffer bytes.Buffer
	if err := json.Indent(&buffer, []byte(res.Body), "", "  "); err != nil {
		// TODO: Handle this error
	}

	contentType := res.Headers.Get("Content-Type")
	if err := quick.Highlight(
		w,
		buffer.String(),
		getLangFromContentType(contentType),
		"terminal16m",
		"monokai",
	); err != nil {
		// TODO: Default colour?
	}

	t.responseBodyTextView.ScrollToBeginning()
}

func getLangFromContentType(contentType string) string {
	if strings.Contains(contentType, "html") {
		return "html"
	} else if strings.Contains(contentType, "json") {
		return "json"
	} else {
		return "plain"
	}
}
