package environment

import (
	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/tui/dialog"
	"go.etcd.io/bbolt"
)

type Environment struct {
	db    *bbolt.DB
	app   *tview.Application
	pages *tview.Pages

	list  *tview.List
	table *tview.Table
}

func New(db *bbolt.DB, app *tview.Application, pages *tview.Pages) Environment {
	return Environment{
		db,
		app,
		pages,
		tview.NewList(),
		tview.NewTable(),
	}
}

func (e Environment) Root() tview.Primitive {
	e.initList()
	e.initTable()

	flex := tview.NewFlex().
		AddItem(e.list, 30, 1, true).
		AddItem(e.table, 0, 1, false)

	return dialog.NewPanel(flex, 0, 0)
}
