package request

import (
	"net/http"

	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/tui/dialog"
)

var methods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
	http.MethodHead,
	http.MethodOptions,
	http.MethodPatch,
	http.MethodConnect,
	http.MethodTrace,
}

func (p Panel) requestForm() *tview.Form {
	form := tview.NewForm().
		AddDropDown("Method", methods, 0, p.dropdownHandler).
		AddInputField("URL", "", 0, nil, p.inputHandler).
		AddButton("Headers", p.headersBtnHandler).
		AddButton("Send", p.doRequest)

	form.
		SetBorder(true).
		SetTitle("Request Form")

	p.model.Components[0] = form
	return form
}

func (p Panel) dropdownHandler(option string, _ int) {
	p.model.RequestData.Method = option
	p.updateRequestSummary()
}

func (p Panel) inputHandler(text string) {
	p.model.RequestData.Url = text
	p.updateRequestSummary()
}

func (p Panel) headersBtnHandler() {
	headersDialog := dialog.NewHeaderDialog(p.model)
	p.model.Pages.AddAndSwitchToPage("HEADERS_DIALOG", headersDialog.Root(), true)
}

func (p Panel) doRequest() {
	p.model.RequestData.Body = p.requestBodyTextArea.GetText()
	res, err := p.api.DoRequest(p.model.RequestData)
	p.model.UpdateResponseSummary(res, err)
}
