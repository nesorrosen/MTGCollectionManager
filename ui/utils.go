package ui

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// Center returns a new primitive which shows the provided primitive in its
// center, given the provided primitive's size.
func Center(width, height int, p tview.Primitive) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

// LogoArt returns a new primitive, and it's size, which contains an
// Ascii art of the input text
func LogoArt(text, font string) (width, height int, logoBox *tview.TextView) {
	logo := figure.NewFigure(text, font, true)
	logoLines := logo.Slicify()

	width = 0
	height = len(logoLines) + 1
	for _, line := range logoLines {
		if len(line) > width {
			width = len(line)
		}
	}

	logoBox = tview.NewTextView().SetTextColor(tcell.ColorGreen)
	fmt.Fprint(logoBox, logo.String())

	return
}

func listVimNavigation(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'h':
		return tcell.NewEventKey(tcell.KeyBackspace, 0, tcell.ModNone)
	case 'j':
		return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
	case 'k':
		return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
	case 'l':
		return tcell.NewEventKey(tcell.KeyEnter, 0, tcell.ModNone)
	}
	return event
}
