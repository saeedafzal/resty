package environment

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/saeedafzal/resty/helper"
	"github.com/saeedafzal/resty/tui/dialog"
	"go.etcd.io/bbolt"
)

func (e Environment) initTable() {
	e.table.SetFixed(0, 0)

	e.table.
		SetBorder(true).
		SetTitle("Environment Variables").
		SetFocusFunc(func() { e.table.SetSelectable(true, false) }).
		SetBlurFunc(func() { e.table.SetSelectable(false, false) }).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Rune() {
			case 'a':
				e.pages.AddPage(helper.FORM_DIALOG, dialog.InputTwoDialog(e.pages, "Add Variable", "", "", e.addEnvironmentVarCallback), true, true)
				return nil
			case 'd':
				e.removeEnvironmentVar()
				return nil
			case 'e':
				row, _ := e.table.GetSelection()
				key := e.table.GetCell(row, 0)
				value := e.table.GetCell(row, 1)
				e.pages.AddPage(helper.FORM_DIALOG, dialog.InputTwoDialog(e.pages, "Edit Variable", key.Text, value.Text, e.addEnvironmentVarCallback), true, true)
				return nil
			}

			if event.Key() == tcell.KeyESC || event.Key() == tcell.KeyTAB {
				e.app.SetFocus(e.list)
				return nil
			}

			return event
		})
}

func (e Environment) displayEnvrionmentVariables(env string) {
	e.table.Clear().
		SetCell(0, 0, helper.HeaderCell("Name")).
		SetCell(0, 1, helper.HeaderCell("Value"))

	index := e.list.GetCurrentItem()
	name, _ := e.list.GetItemText(index)
	index = 1

	if err := e.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(helper.ENV_BUCKET)
		b = b.Bucket([]byte(name))

		return b.ForEach(func(k, v []byte) error {
			e.table.
				SetCell(index, 0, helper.Cell(string(k))).
				SetCell(index, 1, helper.Cell(string(v)))
			index = index + 1
			return nil
		})
	}); err != nil {
		// TODO: ?
		log.Panic(err)
	}
}

func (e Environment) addEnvironmentVarCallback(key, value string) {
	index := e.list.GetCurrentItem()
	name, _ := e.list.GetItemText(index)

	if err := e.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(helper.ENV_BUCKET)
		b = b.Bucket([]byte(name))
		return b.Put([]byte(key), []byte(value))
	}); err != nil {
		// TODO: ?
		log.Panic(err)
	}

	i := e.table.GetRowCount()
	e.table.
		SetCell(i, 0, helper.Cell(key)).
		SetCell(i, 1, helper.Cell(value))

	e.pages.RemovePage(helper.FORM_DIALOG)
}

func (e Environment) removeEnvironmentVar() {
	index := e.list.GetCurrentItem()
	name, _ := e.list.GetItemText(index)

	row, _ := e.table.GetSelection()
	cell := e.table.GetCell(row, 0)

	if err := e.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(helper.ENV_BUCKET)
		b = b.Bucket([]byte(name))
		return b.Delete([]byte(cell.Text))
	}); err != nil {
		// TODO: ?
		log.Panic(err)
	}

	e.table.RemoveRow(row)
	e.pages.RemovePage(helper.FORM_DIALOG)
}

func (e Environment) editEnvironmentVarCallback(key, value string) {
	index := e.list.GetCurrentItem()
	name, _ := e.list.GetItemText(index)

	row, _ := e.table.GetSelection()
	keyCell := e.table.GetCell(row, 0)
	valueCell := e.table.GetCell(row, 1)

	if err := e.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(helper.ENV_BUCKET)
		b = b.Bucket([]byte(name))

		if err := b.Delete([]byte(keyCell.Text)); err != nil {
			return err
		}

		return b.Put([]byte(key), []byte(value))
	}); err != nil {
		// TODO: ?
		log.Panic(err)
	}

	keyCell.SetText(key)
	valueCell.SetText(value)

	e.pages.RemovePage(helper.FORM_DIALOG)
}
