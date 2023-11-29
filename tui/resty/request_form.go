package resty

import (
	"net/http"

	"github.com/rivo/tview"
)

var httpMethods = []string{
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

func (r *Resty) requestForm() *tview.Form {
	form := tview.NewForm().
		AddDropDown("Method", httpMethods, 0, r.dropdownHandler).
		AddInputField("URL", "", 0, nil, r.inputHandler).
		AddButton("Headers", r.headersBtnHandler).
		AddButton("Send", r.sendRequest)

	form.
		SetBorder(true).
		SetTitle("Request Form")

	r.components[0] = form
	return form
}

func (r *Resty) dropdownHandler(option string, _ int) {
	r.requestData.Method = option
	r.updateRequestSummary()
}

func (r *Resty) inputHandler(text string) {
	r.requestData.Url = text
	r.updateRequestSummary()
}

func (r *Resty) headersBtnHandler() {
	r.pages.AddAndSwitchToPage("HEADERS_DIALOG", r.headersDialog(), true)
}

func (r *Resty) sendRequest() {
	r.requestData.Body = r.requestBodyTextArea.GetText()
	res, err := r.api.DoRequest(r.requestData)
	r.responseBodyTextView.Clear()
	r.UpdateResponseSummaryTextView(res, err)
	if err == nil {
		r.updateResponseBody(res)
	}
}
