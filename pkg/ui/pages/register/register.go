package register

import (
	"github.com/rivo/tview"
	"terminalGPT/pkg/ui/pages"
)

type Register struct {
	pages *tview.Pages
}

func NewRegister(pages *tview.Pages) *Register {
	return &Register{pages: pages}
}

func (r Register) Page() *tview.Grid {
	form := tview.NewForm().
		AddInputField("username: ", "", 30, nil, nil).
		AddPasswordField("password: ", "", 30, '*', nil).
		AddButton("register new user", func() {
			r.pages.SwitchToPage(pages.CHAT)
		}).
		AddButton("come back", func() {
			r.pages.SwitchToPage(pages.LAUNCH)
		}).
		SetButtonsAlign(tview.AlignCenter)

	grid := tview.NewGrid().
		SetRows(0, 3, 0).
		SetColumns(0, 50, 0).
		AddItem(tview.NewBox(), 0, 0, 1, 3, 0, 0, false).
		AddItem(tview.NewBox(), 1, 0, 1, 1, 0, 0, false).
		AddItem(form, 1, 1, 2, 1, 0, 0, true)

	return grid
}
