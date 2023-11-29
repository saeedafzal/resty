package resty

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/model"
)

func (r *Resty) responseBody() *tview.TextView {
	textview := r.responseBodyTextView

	textview.
		SetBorder(true).
		SetTitle("Response Summary")

	r.components[4] = textview
	return textview
}

func (r *Resty) updateResponseBody(res model.ResponseData) {
	contentType := res.Headers.Get("Content-Type")
	lang := r.getLangFromContentType(contentType)

	if lang != "plain" {
		// Add syntax highlighting
		w := tview.ANSIWriter(r.responseBodyTextView)
		buffer := bytes.Buffer{}

		// If json, make sure response is indented
		if contentType == "json" {
			_ = json.Indent(&buffer, []byte(res.Body), "", "  ")
		} else {
			buffer = *bytes.NewBufferString(res.Body)
		}

		err := quick.Highlight(
			w,
			buffer.String(),
			lang,
			"terminal16m",
			"monokai",
		)
		if err == nil {
			return
		}
	}

	r.responseBodyTextView.SetText(res.Body)
}

func (_ *Resty) getLangFromContentType(contentType string) string {
	if strings.Contains(contentType, "html") {
		return "html"
	} else if strings.Contains(contentType, "json") {
		return "json"
	} else {
		return "plain"
	}
}
