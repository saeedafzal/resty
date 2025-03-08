package resty

import "github.com/saeedafzal/resty/model"

func (r Resty) initResponseBodyTextView() {
	textview := r.responseBodyTextView.
		SetDynamicColors(true)

	textview.
		SetBorder(true).
		SetTitle("Response Body")

	r.components[4] = textview
}

func (r Resty) updateResponseBody(res *model.ResponseData) {
	r.responseBodyTextView.SetText(res.Body)
}
