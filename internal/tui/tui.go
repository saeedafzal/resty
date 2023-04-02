package tui

import (
	"fmt"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/internal/api"
)

type UI struct {
	app     *tview.Application
	request api.Request
	api     api.Api
	pages   *tview.Pages
}

func NewUI(app *tview.Application) UI {
	return UI{
		app:     app,
		request: api.NewDefaultRequest(),
		api:     api.NewApi(),
		pages:   tview.NewPages(),
	}
}

func (u UI) Layout() *tview.Pages {
	u.pages.
		AddPage(basePage, u.baseView(), true, true).
		AddPage(sendingDialogPage, u.sendingDialog(), true, false).
		AddPage(addHeaderDialog, u.addHeaderDialog(), true, false)
	return u.pages
}

func (u UI) doRequest() {
	res, err := u.api.DoRequest(u.request)
	if err != nil {
		// TODO: Show error modal
	}

	doc := strings.Builder{}
	doc.WriteString(fmt.Sprintf("Status Code:   %s\n", u.formatStatusCode(res.StatusCode)))
	doc.WriteString(fmt.Sprintf("Response Time: %dms\n", res.ResponseTime))
	doc.WriteString("\n[-:-:bu]Headers[-:-:-]\n")
	for k := range res.Headers {
		doc.WriteString(fmt.Sprintf("%s: %s\n", k, res.Headers.Get(k)))
	}

	responseSummaryTextView.SetText(doc.String())
	u.formatBody(res.Body, res.Headers.Get("Content-Type"))
}

func (u UI) formatStatusCode(code int) string {
	colour := "-"

	if code >= 500 {
		colour = "red"
	} else if code >= 400 {
		colour = "yellow"
	} else if code >= 200 {
		colour = "green"
	}

	return fmt.Sprintf("[%s:-:b]%d[-:-:-]", colour, code)
}

func (u UI) formatBody(body string, contentType string) {
	responseBodyTextView.SetText("")
	w := tview.ANSIWriter(responseBodyTextView)
	if err := quick.Highlight(w, body, getLangFromContentType(contentType), "terminal16m", "monokai"); err != nil {
		// TODO: Display error modal? Use default colour or just set the text?
	}
}

// TODO: Move to another file
func getLangFromContentType(contentType string) string {
	if strings.Contains(contentType, "html") {
		return "html"
	}
	if strings.Contains(contentType, "json") {
		return "json"
	}
	if strings.Contains(contentType, "javascript") {
		return "javascript"
	}
	return "plaintext"
}
