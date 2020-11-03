package ui

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/nesorrosen/MTGCollectionManager/db"
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
	card := &db.CardComplete{}

	// Init logo
	logoWidth, logoHeight, logoBox := LogoArt("MTG", "isometric1")
	logoFlex := Center(logoWidth, logoHeight, logoBox)

	cardForm := tview.NewForm()

	nameField := tview.NewInputField().
		SetFieldWidth(20).
		SetLabel("Card Name").
		SetChangedFunc(func(text string) {
			card.Name = text
		})

	amountField := tview.NewInputField().
		SetFieldWidth(20).
		SetLabel("Amount").
		SetAcceptanceFunc(func(textToCheck string, lastChar rune) bool {
			num, err := strconv.Atoi(textToCheck)
			if err != nil {
				return false
			}
			if num > 255 {
				return false
			}
			return true
		}).
		SetChangedFunc(func(text string) {
			num, _ := strconv.Atoi(text)
			card.Amount = num
		})

	cardForm.AddFormItem(nameField)
	cardForm.AddFormItem(amountField)
	cardForm.
		AddButton("Add Card", func() {
			db.AddCard(card)
		}).
		AddButton("Cancel", func() {
			addCardPage.Clear()
			Pages.SwitchToPage("mainMenuPage")
		})

	formFrame := tview.NewFrame(cardForm).
		AddText("Add Card Form", true, tview.AlignCenter, tcell.ColorWhite)

	frameFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(logoFlex, logoHeight, 1, false).
		AddItem(formFrame, 0, 1, true)

	width := logoWidth
	// height = logoHeight + number of primitives + (1 + number of textrows in the frame) + 2 * number of formitems + 1 because idk
	height := logoHeight + 2 + (1 + 1) + 2*3 + 1

	resFlex := Center(width, height, frameFlex)

	addCardPage.AddItem(resFlex, 0, 1, true)
}
