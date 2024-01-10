package resty

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/model"
)

func (r Resty) initResponseBody() {
	textview := r.responseBody.
		SetDynamicColors(true).
		SetText("...")

	textview.
		SetBorder(true).
		SetTitle("Response Body")

	r.components[4] = textview
}

func (r Resty) updateResponseBody(res *model.ResponseData) {
	r.responseBody.Clear()

	contentType := res.Headers.Get("Content-Type")
	lang := r.getLangFromContentType(contentType)

	if lang != "plain" {
		// Add syntax highlighting
		w := tview.ANSIWriter(r.responseBody)
		buffer := bytes.Buffer{}

		// Add indentation to json
		if contentType == "json" {
			if err := json.Indent(&buffer, []byte(res.Body), "", "  "); err != nil {
				buffer = *bytes.NewBufferString(res.Body)
			}
		} else {
			buffer = *bytes.NewBufferString(res.Body)
		}

		// Add highlight
		if err := quick.Highlight(
			w,
			buffer.String(),
			lang,
			"terminal16m",
			"monokai",
		); err == nil {
			return
		}
	}

	r.responseBody.SetText(res.Body)
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
