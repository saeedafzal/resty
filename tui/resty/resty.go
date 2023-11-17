package resty

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/model"
	"github.com/saeedafzal/resty/tui/resty/request"
	"github.com/saeedafzal/resty/tui/resty/response"
)

// Resty is the main page where you can make the actual HTTP requests.
type Resty struct {
	model *model.Model
}

func NewResty(model *model.Model) Resty {
	return Resty{model}
}

func (r Resty) Root() *tview.Flex {
	req := request.NewPanel(r.model)
	res := response.NewPanel(r.model)

	flex := tview.NewFlex().
		AddItem(req.Root(), 0, 1, true).
		AddItem(res.Root(), 0, 1, false)

	flex.SetInputCapture(r.restyInputCapture)
	return flex
}

func (r Resty) restyInputCapture(event *tcell.EventKey) *tcell.EventKey {
	m := r.model
	c := m.Components

	switch event.Key() {
	case tcell.KeyCtrlJ:
		i := r.getCurrentIndex() + 1
		if i >= len(m.Components) {
			i = 0
		}
		m.App.SetFocus(c[i])
		return nil
	case tcell.KeyCtrlK:
		i := r.getCurrentIndex() - 1
		if i < 0 {
			i = len(m.Components) - 1
		}
		m.App.SetFocus(m.Components[i])
		return nil
	}

	return event
}

func (r Resty) getCurrentIndex() int {
	for i, v := range r.model.Components {
		if v.HasFocus() {
			return i
		}
	}
	return 0
}
