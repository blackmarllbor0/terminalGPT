package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"terminalGPT/internal/pkg/api/interfaces"
	"terminalGPT/pkg/ui/pages"
	"terminalGPT/pkg/ui/pages/chat"
	"terminalGPT/pkg/ui/pages/lauch"
	"terminalGPT/pkg/ui/pages/login"
	"terminalGPT/pkg/ui/pages/register"
)

type UserInterface struct {
	app   *tview.Application
	pages *tview.Pages

	gptModelAPI interfaces.GPTGenerateText
}

func NewUserInterface(gptModelAPI interfaces.GPTGenerateText) *UserInterface {
	return &UserInterface{
		gptModelAPI: gptModelAPI,
	}
}

func (ui *UserInterface) Start() error {
	ui.app = tview.NewApplication()
	ui.pages = tview.NewPages()

	launchPage := lauch.NewLaunch(ui.pages)
	registerPage := register.NewRegister(ui.pages)
	loginPage := login.NewLogin(ui.pages)
	chatPage := chat.NewChat(ui.gptModelAPI)

	ui.pages.
		AddPage(pages.LAUNCH, launchPage.Page(), true, true).
		AddPage(pages.REGISTER, registerPage.Page(), true, false).
		AddPage(pages.LOGIN, loginPage.Page(), true, false).
		AddPage(pages.CHAT, chatPage.Page(), true, false)

	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			ui.app.Stop()
			fmt.Println("Goodbye! 👋")
		}

		return event
	})

	if err := ui.app.SetRoot(ui.pages, true).EnableMouse(true).Run(); err != nil {
		return err
	}

	return nil
}
