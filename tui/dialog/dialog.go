package dialog

import (
	"github.com/rivo/tview"
)

// Dialog defines the default layout and size for a modal component.
// TODO: Be able to specify dimensions.
func Dialog(p tview.Primitive) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, 10, 1, true).
			AddItem(nil, 0, 1, false), 40, 1, true).
		AddItem(nil, 0, 1, false)
}
