package environment

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/saeedafzal/resty/helper"
	"github.com/saeedafzal/resty/tui/dialog"
	"go.etcd.io/bbolt"
)

func (e Environment) initList() {
	e.list = e.list.
		ShowSecondaryText(false).
		SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
			e.displayEnvrionmentVariables(mainText)
		})

	e.list.
		SetBorder(true).
		SetTitle("Environments").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Rune() {
			case 'a':
				e.pages.AddPage(helper.FORM_DIALOG, dialog.InputDialog(e.pages, "Add Environment", "", e.addEnvironmentCallback), true, true)
				return nil
			case 'd':
				e.removeEnvironment()
				return nil
			case 'e':
				index := e.list.GetCurrentItem()
				name, _ := e.list.GetItemText(index)
				e.pages.AddPage(helper.FORM_DIALOG, dialog.InputDialog(e.pages, "Edit Environment", name, e.editEnvironmentCallback), true, true)
				return nil
			// Navigation (TODO: Wrap around)
			case 'j':
				e.list.SetCurrentItem(e.list.GetCurrentItem() + 1)
				return nil
			case 'k':
				e.list.SetCurrentItem(e.list.GetCurrentItem() - 1)
				return nil
			}

			if event.Key() == tcell.KeyEnter || event.Key() == tcell.KeyTAB {
				e.app.SetFocus(e.table)
				return nil
			}

			return event
		})

	// Render environments in DB
	if err := e.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(helper.ENV_BUCKET)
		return b.ForEachBucket(func(k []byte) error {
			e.list.AddItem(string(k), "", 0, nil)
			return nil
		})
	}); err != nil {
		// TODO: ?
		log.Panic(err)
	}
}

func (e Environment) addEnvironmentCallback(name string) {
	if err := e.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(helper.ENV_BUCKET)
		_, err := b.CreateBucketIfNotExists([]byte(name))
		return err
	}); err != nil {
		// TODO: ?
		log.Panic(err)
	}

	e.list.AddItem(name, "", 0, nil)
	e.pages.RemovePage(helper.FORM_DIALOG)
}

func (e Environment) removeEnvironment() {
	index := e.list.GetCurrentItem()
	name, _ := e.list.GetItemText(index)

	if err := e.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(helper.ENV_BUCKET)
		return b.DeleteBucket([]byte(name))
	}); err != nil {
		// TODO: ?
		log.Panic(err)
	}

	e.list.RemoveItem(index)
}

func (e Environment) editEnvironmentCallback(name string) {
	index := e.list.GetCurrentItem()
	oldName, _ := e.list.GetItemText(index)

	if err := e.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(helper.ENV_BUCKET)

		_, err := b.CreateBucket([]byte(name))
		if err != nil {
			return err
		}

		return b.DeleteBucket([]byte(oldName))
	}); err != nil && err != bbolt.ErrBucketExists {
		// TODO: ?
		log.Panic(err)
	}

	e.list.SetItemText(index, name, "")
	e.pages.RemovePage(helper.FORM_DIALOG)
}
