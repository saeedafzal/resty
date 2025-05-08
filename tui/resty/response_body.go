package resty

import (
	"encoding/json"
	"strings"
	"bytes"
	"log/slog"
	"io"
	"errors"

	"github.com/saeedafzal/resty/model"

	"github.com/saeedafzal/tview"

	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/alecthomas/chroma/v2/formatters"
)

func (r Resty) initResponseBodyTextView() {
	textview := r.responseBodyTextView.
		SetDynamicColors(true)

	textview.
		SetBorder(true).
		SetTitle("Response Body")

	r.components[4] = textview
}

func (r Resty) updateResponseBody(res *model.ResponseData) {
	displayData := res.Body
	contentType := res.Headers.Get("Content-Type")

	lexer := lexers.Fallback

	// If json, indent it properly for display.
	if strings.Contains(contentType, "application/json") {
		lexer = lexers.Get("json")
		var dst bytes.Buffer
		if err := json.Indent(&dst, []byte(displayData), "", "  "); err != nil {
			slog.Warn("Failed to json indent response:", "content-type", contentType, "data", res.Body)
		} else {
			displayData = dst.String()
		}
	}

	// Syntax highlighting on response body text.
	style := styles.Get("monokai")
	if style == nil {
		style = styles.Fallback
	}

	formatter := formatters.Get("terminal16m")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	content, err1 := io.ReadAll(strings.NewReader(displayData))
	iterator, err2 := lexer.Tokenise(nil, string(content))

	r.responseBodyTextView.Clear()
	w := tview.ANSIWriter(r.responseBodyTextView)
	if err3 := formatter.Format(w, style, iterator); errors.Join(err1, err2, err3) != nil {
		r.responseBodyTextView.SetText(displayData)
	}
}
