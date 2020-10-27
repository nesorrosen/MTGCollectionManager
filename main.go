package main

import "github.com/nesorrosen/MTGCollectionManager/ui"

func main() {

	if err := ui.App.SetRoot(ui.Pages, true).Run(); err != nil {
		panic(err)
	}
}
