package response

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

type Panel struct {
	model *model.Model

	responseSummaryTextView *tview.TextView
	responseBodyTextView    *tview.TextView
}

func NewPanel(m *model.Model) Panel {
	return Panel{
		model: m,

		responseSummaryTextView: tview.NewTextView(),
		responseBodyTextView:    tview.NewTextView(),
	}
}

func (r Panel) Root() *tview.Flex {
	r.initResponseSummaryTextView()
	r.initResponseBodyTextView()
	r.model.UpdateResponsePanels = r.updateResponsePanels

	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(r.responseSummaryTextView, 0, 1, false).
		AddItem(r.responseBodyTextView, 0, 2, false)
}

func (r Panel) initResponseSummaryTextView() {
	textView := r.responseSummaryTextView
	textView.SetDynamicColors(true)
	textView.SetBorder(true).SetTitle("Response Summary")
	r.model.Components[3] = textView
}

func (r Panel) initResponseBodyTextView() {
	textView := r.responseBodyTextView
	textView.SetDynamicColors(true)
	textView.SetBorder(true).SetTitle("Response Body")
	r.model.Components[4] = textView
}

func (r Panel) updateResponsePanels(res model.ResponseData) {
	r.responseSummaryTextView.Clear()

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

	r.responseSummaryTextView.SetText(doc.String())

	// Colorize the response body output
	r.responseBodyTextView.Clear()
	contentType := r.getLangFromContentType(res.Headers.Get("Content-Type"))
	if contentType != "plain" {
		w := tview.ANSIWriter(r.responseBodyTextView)
		buffer := &bytes.Buffer{}

		if contentType == "json" {
			_ = json.Indent(buffer, []byte(res.Body), "", "  ")
		} else {
			buffer = bytes.NewBufferString(res.Body)
		}

		err := quick.Highlight(
			w,
			buffer.String(),
			contentType,
			"terminal16m",
			"monokai", // TODO: Should be customisable
		)

		if err == nil {
			return
		}
	}

	// Default behaviour
	r.responseBodyTextView.SetText(res.Body)
}

// TODO: ?
func (_ Panel) getLangFromContentType(contentType string) string {
	if strings.Contains(contentType, "html") {
		return "html"
	} else if strings.Contains(contentType, "json") {
		return "json"
	} else {
		return "plain"
	}
}
