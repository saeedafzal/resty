package resty

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"strings"

	"github.com/saeedafzal/resty/core"
	"github.com/saeedafzal/resty/model"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/gdamore/tcell/v2"
	"github.com/saeedafzal/tview"
)

func (r Resty) initResponseBodyTextView() {
	textview := r.responseBodyTextView.
		SetDynamicColors(true)

	textview.
		SetBorder(true).
		SetTitle("Response Body").
		SetInputCapture(r.responseBodyTextViewInputCapture)

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

func (r Resty) responseBodyTextViewInputCapture(event *tcell.EventKey) *tcell.EventKey {
	text := r.responseBodyTextView.GetText(true)

	if event.Key() == tcell.KeyCtrlE && text != "" {
		core.OpenEditor(r.app, text, true)
		return nil
	}

	return event
}
