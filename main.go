package main

import (
	"github.com/nesorrosen/MTGCollectionManager/db"
	"github.com/nesorrosen/MTGCollectionManager/ui"
)

func main() {

	db.DeleteSupertype()

	if err := ui.App.SetRoot(ui.Pages, true).Run(); err != nil {
		panic(err)
	}
}
