package main

import (
	"fmt"
	"log"

	"github.com/rivo/tview"
	"github.com/saeedafzal/resty/helper"
	"github.com/saeedafzal/resty/tui"
	"github.com/spf13/pflag"
	"go.etcd.io/bbolt"
)

var version string

func main() {
	// Handle CLI flags and exit if necessary
	if flags() {
		return
	}

	// Start bbolt
	db, err := bbolt.Open("dev.db", 0600, nil)
	if err != nil || createBuckets(db) != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Debug logs

	// Start terminal ui
	app := tview.NewApplication()
	tui := tui.New(app, db)

	app.
		EnableMouse(true).
		SetRoot(tui.Root(), true)

	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}

func flags() bool {
	var v, h bool
	pflag.BoolVarP(&v, "version", "v", false, "Display application version.")
	pflag.BoolVarP(&h, "help", "h", false, "Usage of Resty.")
	pflag.Parse()

	if v {
		fmt.Println("Resty", version)
		return true
	}

	if h {
		pflag.Usage()
		return true
	}

	return false
}

func createBuckets(db *bbolt.DB) error {
	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(helper.ENV_BUCKET)
		return err
	})
}
