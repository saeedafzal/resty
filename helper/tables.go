package helper

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func HeaderCell(text string) *tview.TableCell {
	return tview.NewTableCell(text).
		SetSelectable(false).
		SetTextColor(tcell.ColorYellow).
		SetExpansion(1).
		SetAlign(tview.AlignCenter)
}

func Cell(text string) *tview.TableCell {
	return tview.NewTableCell(text).
		SetSelectable(true).
		SetExpansion(1).
		SetAlign(tview.AlignCenter)
}
