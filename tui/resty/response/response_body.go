package response

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/model"
)

func (p Panel) initResponseBodyTextView() {
	p.responseBodyTextView.
		SetBorder(true).
		SetTitle("Response Body")

	p.model.Components[4] = p.responseBodyTextView
}

func (p Panel) updateResponseBodyTextView(res model.ResponseData) {
	p.responseBodyTextView.Clear()

	contentType := p.getLangFromContentType(res.Headers.Get("Content-Type"))
	if contentType != "plain" {
		w := tview.ANSIWriter(p.responseBodyTextView)
		buffer := &bytes.Buffer{}

		// If json, make sure response string is indented
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
			"monokai",
		)

		if err == nil {
			return
		}
	}

	p.responseBodyTextView.SetText(res.Body)
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
