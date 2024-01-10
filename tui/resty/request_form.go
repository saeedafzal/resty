package resty

import (
	"net/http"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/api"
	"github.com/saeedafzal/resty/helper"
)

var methods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
	http.MethodHead,
	http.MethodOptions,
	http.MethodTrace,
	http.MethodPatch,
	http.MethodConnect,
}

func (r Resty) requestForm() *tview.Form {
	req := r.requestData

	form := tview.NewForm().
		SetButtonBackgroundColor(tcell.ColorIndigo).
		SetFieldBackgroundColor(tcell.ColorIndigo).
		AddDropDown("Method", methods, 0, func(option string, _ int) {
			req.Method = option
			r.updateRequestSummary()
		}).
		AddInputField("URL", "", 0, nil, func(text string) {
			req.Url = text
			r.updateRequestSummary()
		}).
		AddButton("Headers", func() {
			r.pages.AddAndSwitchToPage(helper.HEADER_DIALOG, r.headersDialog(), true)
		}).
		AddButton("Send", func() {
			r.requestData.Body = r.requestBody.GetText()
			res, err := api.DoRequest(r.requestData)
			if err != nil {
				r.updateResponseSummaryError(err)
				r.responseBody.SetText("")
				return
			}
			r.updateResponseSummary(res)
			r.updateResponseBody(res)
		})

	form.
		SetBorder(true).
		SetTitle("Request Form")

	r.components[0] = form
	return form
}
