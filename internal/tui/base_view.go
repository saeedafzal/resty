package tui

import "github.com/rivo/tview"

func (u UI) baseView() *tview.Flex {
	flex := tview.NewFlex().
		AddItem(u.requestView(), 0, 1, true).
		AddItem(u.responseView(), 0, 1, false)

	return flex
}

func (_ UI) createTable() *tview.Table {
	cell := func(text string) *tview.TableCell {
		return tview.NewTableCell(text).
			SetSelectable(false).
			SetExpansion(1)
	}

	table := tview.NewTable().
		SetSelectable(true, true).
		SetBorders(true).
		SetFixed(1, 1).
		SetCell(0, 0, cell("Name")).
		SetCell(0, 1, cell("Value"))

	return table
}
