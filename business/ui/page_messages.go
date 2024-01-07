package ui

import "github.com/rivo/tview"

func (a *CBAToolApp) callQuit() {
	modal := tview.NewModal().SetText("Do you want to quit the application?").AddButtons([]string{"Quit", "Cancel"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Quit" {
			a.app.Stop()
		} else {
			a.callMenuPage()
		}
	})
	a.app.SetRoot(modal, true)
}

func (a *CBAToolApp) callErrorMessage(e error) {
	modal := tview.NewModal()
	modal.SetTitle("Error message")
	modal.SetText(e.Error())
	modal.AddButtons([]string{"Continue"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		a.callMenuPage()
	})
	a.app.SetRoot(modal, true)
}