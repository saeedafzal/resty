package request

import (
	"net/http"

	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/util"
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

func (r Panel) requestForm() *tview.Form {
	primary := util.HexToColour("#6F00FF")

	form := tview.NewForm().
		AddDropDown("Method", methods, 0, r.requestFormDropdownHandler).
		AddInputField("URL", "", 0, nil, r.requestFormInputHandler).
		AddButton("Add Header", r.requestFormAddHeaderButtonHandler).
		AddButton("Send", r.sendRequest).
		// TODO: Colour from config
		SetFieldBackgroundColor(primary).
		SetButtonBackgroundColor(primary)

	form.
		SetBorder(true).
		SetTitle("Request Form")

	r.model.Components[0] = form
	return form
}

// NOTE: Handler for the method dropdown on the request form.
func (r Panel) requestFormDropdownHandler(option string, _ int) {
	r.model.RequestData.Method = option
	r.updateRequestSummaryTextView()
}

// NOTE: Handler for the url input field on the request form.
func (r Panel) requestFormInputHandler(text string) {
	r.model.RequestData.Url = text
	r.updateRequestSummaryTextView()
}

// NOTE: Handler for the "Add Header" button on the request form.
func (r Panel) requestFormAddHeaderButtonHandler() {
	r.model.Pages.ShowPage(util.AddHeaderDialogPage)
}

func (r Panel) sendRequest() {
	m := r.model

	m.RequestData.Body = r.requestBodyTextArea.GetText()
	res, err := r.api.DoRequest(m.RequestData)
	if err != nil {
		// TODO: Have some UI feedback on failed request
	}

	m.UpdateResponsePanels(res)
}
