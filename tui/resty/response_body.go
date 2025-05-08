package resty

import (
	"encoding/json"
	"strings"
	"bytes"
	"log/slog"

	"github.com/saeedafzal/resty/model"
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

	// If json, indent it properly for display.
	if strings.Contains(contentType, "application/json") {
		var dst bytes.Buffer
		if err := json.Indent(&dst, []byte(displayData), "", "  "); err != nil {
			slog.Warn("Failed to json indent response:", "content-type", contentType, "data", res.Body)
		} else {
			displayData = dst.String()
		}
	}

	r.responseBodyTextView.SetText(displayData)
}
