package lauch

import (
	"github.com/rivo/tview"
	"terminalGPT/pkg/ui/pages"
)

type Launch struct {
	pages *tview.Pages
}

func NewLaunch(pages *tview.Pages) *Launch {
	return &Launch{pages: pages}
}

func (l Launch) Page() *tview.Grid {
	grid := tview.NewGrid().
		SetRows(0, 3, 0).
		SetColumns(0, 0, 0, 0).
		SetGap(0, 5)

	loginButton := tview.NewButton("Log-In").SetSelectedFunc(func() {
		l.pages.SwitchToPage(pages.LOGIN)
	})

	registerButton := tview.NewButton("Register").SetSelectedFunc(func() {
		l.pages.SwitchToPage(pages.REGISTER)
	})

	grid.AddItem(loginButton, 1, 1, 1, 1, 0, 0, false).
		AddItem(registerButton, 1, 2, 1, 1, 0, 0, false)

	return grid
}
