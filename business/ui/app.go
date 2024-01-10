package ui

/*
https://github.com/rivo/tview/wiki/Table
https://gist.github.com/rivo/2893c6740a6c651f685b9766d1898084
https://github.com/rivo/tview/wiki/Postgres
https://github.com/rivo/tview/wiki
https://github.com/destinmoulton/pixi/blob/master/gui/gui.go
https://github.com/rivo/tview
*/
import (
	"github.com/LosAngeles971/cba-tool/business/cba"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CBAToolApp struct {
	Data *cba.CBA
	app  *tview.Application
}

func (a *CBAToolApp) eventHandler(eventKey *tcell.EventKey) *tcell.EventKey {
	if eventKey.Key() == tcell.KeyEscape {
		a.callMenuPage()
	}
	return eventKey
}

func Build() *CBAToolApp {
	cbaApp := &CBAToolApp{
		Data: cba.NewCBA(),
	}
	cbaApp.app = tview.NewApplication()

	cbaApp.app.EnableMouse(true)
	cbaApp.app.SetInputCapture(cbaApp.eventHandler)

	cbaApp.callMenuPage()

	return cbaApp
}

func (a *CBAToolApp) Run() error {
	return a.app.Run()
}
