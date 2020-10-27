package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	// App is the app
	App *tview.Application

	// Pages represents the apps pages
	Pages *tview.Pages

	mainMenuPage *tview.Flex
	overViewPage *tview.Flex
	viewCardPage *tview.Flex
	addCardPage  *tview.Flex
)

func init() {
	App = tview.NewApplication()
	Pages = tview.NewPages()
	Pages.SetBackgroundColor(tcell.ColorBlack)

	mainMenuPage = tview.NewFlex()
	overViewPage = tview.NewFlex()
	viewCardPage = tview.NewFlex()
	addCardPage = tview.NewFlex()

	Pages.AddPage("mainMenuPage", mainMenuPage, true, true)
	Pages.AddPage("overViewPage", overViewPage, true, false)
	Pages.AddPage("viewCardPage", viewCardPage, true, false)
	Pages.AddPage("addCardPage", addCardPage, true, false)

	Pages.SwitchToPage("mainMenuPage")

	initMainMenu()
	initOverviewPage()
	initViewPage()
	initAddPage()
}

// Init mainmenu
func initMainMenu() {
	// Init logo
	logoWidth, logoHeight, logoBox := LogoArt("MTG", "isometric1")

	menuList := tview.NewList().
		ShowSecondaryText(false).
		AddItem("Add Card", "", '1', func() {
			Pages.SwitchToPage("addCardPage")
		}).
		AddItem("List Cards", "", '2', func() {
			Pages.SwitchToPage("overViewPage")
		}).
		AddItem("Quit", "", 'q', func() {
			App.Stop()
		})

	menuList.SetInputCapture(listVimNavigation)

	menuFrame := tview.NewFrame(menuList)

	mainFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(logoBox, logoHeight, 1, false).
		AddItem(menuFrame, 0, 1, true)

	width := logoWidth
	// height = logoheight + items in menu list + 1 * number of primitives
	height := logoHeight + 3 + 3

	endRes := Center(width, height, mainFlex)

	mainMenuPage.AddItem(endRes, 0, 1, true)
}

func initOverviewPage() {

}

func initViewPage() {

}

func initAddPage() {
	// Init logo
	logoWidth, logoHeight, logoBox := LogoArt("MTG", "puffy")
	logoFlex := Center(logoWidth, logoHeight, logoBox)

	addCardPage.AddItem(logoFlex, 0, 1, false)

}
