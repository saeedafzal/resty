package model

import "github.com/rivo/tview"

type Model struct {
	RequestData RequestData
	App         *tview.Application
	Pages       *tview.Pages
	Components  []tview.Primitive

	// Callbacks
	UpdateRequestSummary  func()
	UpdateResponseSummary func(res ResponseData)
}

func NewModel() *Model {
	return &Model{
		RequestData: NewRequestData(),
		App:         tview.NewApplication().EnableMouse(true),
		Pages:       tview.NewPages(),
		Components:  make([]tview.Primitive, 5),
	}
}
