package tui

import (
	"net/http"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/api"
	"github.com/saeedafzal/resty/model"
)

var methods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
	http.MethodHead,
	http.MethodPatch,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}

type RequestPanel struct {
	model             Model
	requestData       model.RequestData
	updateHeaderValue string

	api           api.API
	responsePanel ResponsePanel

	requestSummaryTextView *tview.TextView
	requestHeadersTable    *tview.Table
	addHeaderForm          *tview.Form
	requestBodyTextArea    *tview.TextArea
}

func NewRequestPanel(m Model, responsePanel ResponsePanel) RequestPanel {
	return RequestPanel{
		model: m,
		requestData: model.RequestData{
			Method:  methods[0],
			Url:     "",
			Headers: make(http.Header),
			Body:    "",
		},

		api:           api.NewAPI(),
		responsePanel: responsePanel,

		requestSummaryTextView: tview.NewTextView().SetDynamicColors(true),
		requestHeadersTable:    tview.NewTable(),
		addHeaderForm:          tview.NewForm(),
		requestBodyTextArea:    tview.NewTextArea(),
	}
}

func (t RequestPanel) Root() *tview.Flex {
	t.createRequestBodyTextArea()

	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(t.requestForm(), 0, 1, true).
		AddItem(t.requestBodyTextArea, 0, 2, false).
		AddItem(t.requestSummaryPanel(), 0, 2, false)
}

func (t RequestPanel) requestForm() *tview.Form {
	form := tview.NewForm().
		AddDropDown("Method", methods, 0, func(option string, _ int) {
			t.requestData.Method = option
			t.updateRequestSummary()
		}).
		AddInputField("URL", "", 0, nil, func(text string) {
			t.requestData.Url = text
			t.updateRequestSummary()
		}).
		AddButton("Add Headers", func() {
			m := t.model
			m.pages.ShowPage(addHeaderPage)
			m.App.SetFocus(t.addHeaderForm)
		}).
		AddButton("Send Request", func() { t.sendRequest() })

	/* c, _ := strconv.ParseInt("FF0000", 16, 32)
	form.SetFieldBackgroundColor(tcell.NewHexColor(int32(c))) */

	form.SetBorder(true).SetTitle("Request Form")
	t.model.components[0] = form
	return form
}

func (t RequestPanel) createRequestBodyTextArea() {
	t.requestBodyTextArea.
		SetBorder(true).
		SetTitle("Request Body")

	t.model.components[1] = t.requestBodyTextArea
}

func (t RequestPanel) requestSummaryPanel() *tview.Flex {
	t.buildRequestHeadersTable()

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(t.requestSummaryTextView, 0, 1, false).
		AddItem(t.requestHeadersTable, 0, 4, false)

	flex.SetBorder(true).SetTitle("Request Summary")
	return flex
}

func (t RequestPanel) addHeaderDialog() tview.Primitive {
	form := t.addHeaderForm
	name, value := "", ""

	form.
		AddInputField("Name", name, 0, nil, func(text string) { name = text }).
		AddInputField("Value", value, 0, nil, func(text string) { value = text }).
		AddButton("Add Header", func() {
			t.requestData.Headers[name] = []string{value}
			t.updateRequestHeadersTable()
		}).
		AddButton("Cancel", func() { t.model.pages.HidePage(addHeaderPage) })

	form.SetBorder(true).
		SetTitle("Headers").
		SetInputCapture(t.dialogInputHandler)

	return dialog(t.addHeaderForm)
}

func (t RequestPanel) showUpdateHeaderDialog() {
	value := ""

	form := tview.NewForm().
		AddInputField("Value", t.updateHeaderValue, 0, nil, func(text string) { value = text }).
		AddButton("Update", func() { t.updateRequestHeadersHandler(value) }).
		AddButton("Cancel", func() { t.model.pages.RemovePage(updateHeaderPage) })

	form.SetBorder(true).
		SetTitle("Update Header Value").
		SetInputCapture(t.updateDialogInputHandler)

	t.model.pages.AddPage(updateHeaderPage, dialog(form), true, true)
	t.model.App.SetFocus(form)
}

func (t RequestPanel) buildRequestHeadersTable() {
	table := t.requestHeadersTable
	t.model.components[2] = table

	table.SetBorders(true).
		SetFixed(0, 0).
		SetSelectable(false, false)

	table.SetFocusFunc(func() { table.SetSelectable(true, true) }).
		SetBlurFunc(func() { table.SetSelectable(false, false) }).
		SetInputCapture(t.requestHeadersTableInputHandler)

	t.updateRequestHeadersTable()
}

func (t RequestPanel) headerCell(name string) *tview.TableCell {
	return tview.NewTableCell(name).
		SetExpansion(1).
		SetSelectable(false).
		SetTextColor(tcell.ColorYellow)
}

func (t RequestPanel) cell(name string) *tview.TableCell {
	return tview.NewTableCell(name).
		SetExpansion(1).
		SetSelectable(true).
		SetMaxWidth(1)
}

// ---

func (t RequestPanel) dialogInputHandler(event *tcell.EventKey) *tcell.EventKey {
	m := t.model
	if event.Key() == tcell.KeyEsc || (event.Rune() == 'q' && !m.isInputField()) {
		m.pages.HidePage(addHeaderPage)
		return nil
	}
	return event
}

func (t RequestPanel) updateDialogInputHandler(event *tcell.EventKey) *tcell.EventKey {
	m := t.model
	if event.Key() == tcell.KeyEsc || (event.Rune() == 'q' && !m.isInputField()) {
		m.pages.RemovePage(updateHeaderPage)
		return nil
	}
	return event
}

func (t RequestPanel) updateRequestSummary() {
	r := t.requestData

	doc := strings.Builder{}
	doc.WriteString("Method: " + r.Method + "\n")
	doc.WriteString("URL:    " + r.Url + "\n\n")
	doc.WriteString("[yellow:-:bu]Headers[-:-:-]")

	t.requestSummaryTextView.SetText(doc.String())
}

func (t RequestPanel) updateRequestHeadersTable() {
	t.requestHeadersTable.
		Clear().
		SetCell(0, 0, t.headerCell("Name")).
		SetCell(0, 1, t.headerCell("Value"))

	i := 1
	for k, v := range t.requestData.Headers {
		t.requestHeadersTable.
			SetCell(i, 0, t.cell(k)).
			SetCell(i, 1, t.cell(v[0]))
		i++
	}
}

func (t RequestPanel) requestHeadersTableInputHandler(event *tcell.EventKey) *tcell.EventKey {
	table := t.requestHeadersTable

	switch event.Rune() {
	case 'd':
		row, _ := table.GetSelection()
		cell := table.GetCell(row, 0)
		delete(t.requestData.Headers, cell.Text)
		table.RemoveRow(row)
		return nil
	case 'e':
		row, col := table.GetSelection()
		cell := table.GetCell(row, col)
		t.updateHeaderValue = cell.Text
		t.showUpdateHeaderDialog()
	}

	return event
}

func (t RequestPanel) updateRequestHeadersHandler(value string) {
	table := t.requestHeadersTable
	headers := t.requestData.Headers

	row, col := table.GetSelection()

	if col == 0 { // Header key
		v := headers[t.updateHeaderValue]
		delete(headers, t.updateHeaderValue)
		headers[value] = v
	} else { // Value
		cell := table.GetCell(row, 0)
		headers[cell.Text] = []string{value}
	}

	t.updateRequestHeadersTable()
	t.model.pages.RemovePage(updateHeaderPage)
}

func (t RequestPanel) sendRequest() {
	t.requestData.Body = t.requestBodyTextArea.GetText()

	res, err := t.api.DoRequest(t.requestData)
	if err != nil {
		// TODO: Some type of error feedback
	}
	t.responsePanel.updateResponsePanels(res)
}
