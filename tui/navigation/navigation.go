package navigation

import (
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/helper"
	"github.com/saeedafzal/resty/tui/dialog"
)

type Navigation struct {
	pages *tview.Pages
}

func New(pages *tview.Pages) Navigation {
	return Navigation{pages}
}

func (n Navigation) Root() tview.Primitive {
	list := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Resty", "", 0, func() { n.pages.SwitchToPage(helper.RESTY_PAGE) }).
		AddItem("Environment", "", 0, func() { n.pages.SwitchToPage(helper.ENVIRONMENT_PAGE) })

	list.
		SetBorder(true).
		SetTitle("Navigation")

	return dialog.NewDialog(list, 30, 0)
}
