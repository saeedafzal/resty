package tui

import (
	"net/http"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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

type RequestModel struct {
	method  string
	url     string
	headers http.Header
}

type RequestPanel struct {
	model             Model
	requestModel      RequestModel
	updateHeaderValue string

	requestSummaryTextView *tview.TextView
	requestHeadersTable    *tview.Table
	addHeaderForm          *tview.Form
}

func NewRequestPanel(model Model) RequestPanel {
	return RequestPanel{
		model: model,
		requestModel: RequestModel{
			method:  methods[0],
			url:     "",
			headers: make(http.Header),
		},

		requestSummaryTextView: tview.NewTextView().SetDynamicColors(true),
		requestHeadersTable:    tview.NewTable(),
		addHeaderForm:          tview.NewForm(),
	}
}

func (t RequestPanel) Root() *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(t.requestForm(), 0, 1, true).
		AddItem(t.requestBodyTextArea(), 0, 2, false).
		AddItem(t.requestSummaryPanel(), 0, 2, false)
}

func (t RequestPanel) requestForm() *tview.Form {
	form := tview.NewForm().
		AddDropDown("Method", methods, 0, func(option string, _ int) {
			t.requestModel.method = option
			t.updateRequestSummary()
		}).
		AddInputField("URL", "", 0, nil, func(text string) {
			t.requestModel.url = text
			t.updateRequestSummary()
		}).
		AddButton("Add Headers", func() {
			m := t.model
			m.pages.ShowPage(addHeaderPage)
			m.App.SetFocus(t.addHeaderForm)
		}).
		AddButton("Send Request", nil)

	/* c, _ := strconv.ParseInt("FF0000", 16, 32)
	form.SetFieldBackgroundColor(tcell.NewHexColor(int32(c))) */

	form.SetBorder(true).SetTitle("Request Form")
	t.model.components[0] = form
	return form
}

func (t RequestPanel) requestBodyTextArea() *tview.TextArea {
	textArea := tview.NewTextArea()

	textArea.SetBorder(true).SetTitle("Request Body")
	t.model.components[1] = textArea
	return textArea
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
			t.requestModel.headers[name] = []string{value}
			t.updateRequestHeadersTable()
		}).
		AddButton("Cancel", func() { t.model.pages.HidePage(addHeaderPage) })

	form.SetBorder(true).
		SetTitle("Headers").
		SetInputCapture(t.dialogInputHandler)

	return dialog(t.addHeaderForm)
}

func (t RequestPanel) showUpdateHeaderDialog() {
	form := tview.NewForm().
		AddInputField("Value", t.updateHeaderValue, 0, nil, nil).
		AddButton("Update", nil).
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
		SetSelectable(true)
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
		m.pages.RemovePage(addHeaderPage)
		return nil
	}
	return event
}

func (t RequestPanel) updateRequestSummary() {
	m := t.requestModel

	doc := strings.Builder{}
	doc.WriteString("Method: " + m.method + "\n")
	doc.WriteString("URL:    " + m.url + "\n\n")
	doc.WriteString("[yellow:-:bu]Headers[-:-:-]")

	t.requestSummaryTextView.SetText(doc.String())
}

func (t RequestPanel) updateRequestHeadersTable() {
	t.requestHeadersTable.
		Clear().
		SetCell(0, 0, t.headerCell("Name")).
		SetCell(0, 1, t.headerCell("Value"))

	i := 1
	for k, v := range t.requestModel.headers {
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
		delete(t.requestModel.headers, cell.Text)
		table.RemoveRow(row)
		return nil
	case 'e':
		row, col := table.GetSelection()
		cell := table.GetCell(row, col)
		t.updateHeaderValue = cell.Text
		t.showUpdateHeaderDialog()
		// t.model.App.SetFocus(t.updateHeaderForm)
	}

	return event
}
