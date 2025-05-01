package resty

import (
	"github.com/saeedafzal/resty/api"
	"github.com/saeedafzal/tview"
	"github.com/saeedafzal/resty/tui/modal"
)

var methods = []string{"GET", "POST", "PUT", "DELETE"}

func (r Resty) requestForm() *tview.Form {
	form := tview.NewForm().
		AddDropDown("Method", methods, 0, r.requestFormDropDownSelected).
		AddInputField("URL", "", 0, nil, r.requestFormInputFieldChanged).
		AddButton("Add Headers", r.requestFormAddHeadersBtnSelected).
		AddButton("Send", r.requestFormSendBtnSelected)

	form.
		SetBorder(true).
		SetTitle("Request Form")

	r.components[0] = form
	return form
}

func (r Resty) requestFormDropDownSelected(option string, _ int) {
	r.requestData.Method = option
	r.updateRequestSummary()
}

func (r Resty) requestFormInputFieldChanged(text string) {
	r.requestData.Url = text
	r.updateRequestSummary()
}

func (r Resty) requestFormAddHeadersBtnSelected() {
	r.pages.AddPage(
		"header",
		modal.HeaderModal(r.pages, "Add Header", "", "", r.addHeaderAddSelected),
		true,
		true,
	)
}

func (r Resty) requestFormSendBtnSelected() {
	r.requestData.Body = r.requestBodyTextArea.GetText()

	res, err := api.DoRequest(r.requestData)
	if err != nil {
		r.updateResponseSummaryError(err)
		r.responseBodyTextView.Clear()
		return
	}

	r.updateResponseSummary(res)
	r.updateResponseBody(res)
}
