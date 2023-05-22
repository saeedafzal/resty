package tui

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (t TUI) requestPanel() *tview.Flex {
	t.requestBodyTextArea.SetBorder(true).SetTitle("Request Body")
	t.components[2] = t.requestBodyTextArea

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(t.requestForm(), 0, 1, true).
		AddItem(t.requestSummaryPanel(), 0, 2, false).
		AddItem(t.requestBodyTextArea, 0, 2, false)
	return flex
}

func (t TUI) requestForm() *tview.Form {
	form := tview.NewForm().
		AddDropDown("Method", []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}, 0, func(option string, _ int) {
			t.requestModel.Method = option
			t.updateRequestSummaryTextView()
		}).
		AddInputField("URL", "", 0, nil, func(text string) {
			t.requestModel.Url = text
			t.updateRequestSummaryTextView()
		}).
		AddButton("Add Header", func() { t.pages.ShowPage(addHeadersPage) }).
		AddButton("Send", func() { t.sendRequest() })
	form.SetBorder(true).SetTitle("Request Form")
	t.components[0] = form
	return form
}

func (t TUI) requestSummaryPanel() *tview.Flex {
	t.updateRequestSummaryTextView()

	t.requestHeadersTable.SetFocusFunc(func() {
		t.requestHeadersTable.SetSelectable(true, true)
	})
	t.requestHeadersTable.SetBlurFunc(func() {
		t.requestHeadersTable.SetSelectable(false, false)
	})

	t.requestHeadersTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'd':
			row, _ := t.requestHeadersTable.GetSelection()
			cell := t.requestHeadersTable.GetCell(row, 0)
			t.requestModel.Headers.Del(cell.Text)
			t.updateRequestHeadersTable()
			t.app.SetFocus(t.requestHeadersTable)
			t.requestHeadersTable.Select(0, 0)
			return nil
		}
		return event
	})
	t.updateRequestHeadersTable()
	t.components[1] = t.requestHeadersTable

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(t.requestSummaryTextView, 0, 1, false).
		AddItem(t.requestHeadersTable, 0, 5, false)
	flex.SetBorder(true).SetTitle("Request Summary")

	return flex
}

func (t TUI) addHeadersModal() *tview.Flex {
	width, height := 40, 10

	name, value := "", ""

	form := tview.NewForm().
		AddInputField("Name", "", 0, nil, func(text string) {
			name = text
		}).
		AddInputField("Value", "", 0, nil, func(text string) {
			value = text
		}).
		AddButton("Add Header", func() {
			t.requestModel.Headers.Set(name, value)
			t.updateRequestHeadersTable()
		}).
		AddButton("Cancel", func() {
			t.pages.HidePage(addHeadersPage)
		})
	form.SetBorder(true).SetTitle("Add New Header")

	flex := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(form, height, 1, false).
			AddItem(nil, 0, 1, false), width, 1, false).
		AddItem(nil, 0, 1, false)

	flex.SetFocusFunc(func() {
		t.app.SetFocus(form)
	})

	return flex
}

// ---

func (t TUI) updateRequestSummaryTextView() {
	doc := strings.Builder{}
	doc.WriteString("Method: " + t.requestModel.Method + "\n")
	doc.WriteString("URL:    " + t.requestModel.Url + "\n")
	doc.WriteString("\n")
	doc.WriteString("[yellow:-:bu]Headers[-:-:-]")

	t.requestSummaryTextView.SetText(doc.String())
}

func (t TUI) updateRequestHeadersTable() {
	t.requestHeadersTable.Clear()

	headerCell := func(name string) *tview.TableCell {
		return tview.NewTableCell(name).
			SetExpansion(1).
			SetSelectable(false).
			SetTextColor(tcell.ColorYellow)
	}

	cell := func(name string) *tview.TableCell {
		return tview.NewTableCell(name).
			SetExpansion(1).
			SetSelectable(true)
	}

	t.requestHeadersTable.
		SetBorders(true).
		SetSelectable(false, false).
		SetFixed(0, 0).
		SetCell(0, 0, headerCell("Name")).
		SetCell(0, 1, headerCell("Value"))

	i := 1
	for k := range t.requestModel.Headers {
		v := t.requestModel.Headers.Get(k)
		t.requestHeadersTable.
			SetCell(i, 0, cell(k)).
			SetCell(i, 1, cell(v))
		i++
	}
}

func (t TUI) sendRequest() {
	body := t.requestBodyTextArea.GetText()
	buffer := bytes.NewBuffer([]byte(body))
	t.requestModel.Body = buffer

	res, err := t.api.DoRequest(t.requestModel)
	if err != nil {
		// TODO: Display error modal
		t.responseSummaryTextView.SetText("")
	}

	t.updateResponseSummaryTextView(res)
	t.updateResponseBodyTextView(res.Headers.Get("Content-Type"), res.Body)
}
