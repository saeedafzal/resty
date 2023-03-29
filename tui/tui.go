package tui

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/api"
)

type UI struct {
	app     *tview.Application
	request api.Request
	api     api.Api
}

func NewUI(app *tview.Application) UI {
	return UI{
		app:     app,
		request: api.NewDefaultRequest(),
		api:     api.NewApi(),
	}
}

func (u UI) Layout() *tview.Pages {
	pages := tview.NewPages().
		AddPage(basePage, u.baseView(), true, true)

	return pages
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
	for k, _ := range res.Headers {
		doc.WriteString(fmt.Sprintf("%s: %s\n", k, res.Headers.Get(k)))
	}

	responseSummaryTextView.SetText(doc.String())
	responseBodyTextView.SetText(res.Body)
}

func (u UI) formatStatusCode(code int) string {
	colour := "-"

	if code >= 200 || code < 300 {
		// NOTE: Green for 200 response codes
		colour = "green"
	} else if code >= 200 || code < 300 {
		// NOTE: Amber for 400 response codes
		colour = "amber"
	} else if code >= 500 || code <= 599 {
		colour = "red"
	}

	return fmt.Sprintf("[%s:-:b]%d[-:-:-]", colour, code)
}
