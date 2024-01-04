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
	app *tview.Application
	mainMenu *tview.List
	pages *tview.Pages
	cyclesPage *tview.Table
}

func (a *CBAToolApp) eventHandler(eventKey *tcell.EventKey) *tcell.EventKey {
	if eventKey.Rune() == 'q' {
		a.app.Stop()
		return nil
	}
	return eventKey
}

func (a *CBAToolApp) buildMainMenu() {
	a.mainMenu = tview.NewList().ShowSecondaryText(false)
	a.mainMenu.SetBorder(true).SetTitle("Main menù")
	a.mainMenu.AddItem("Load project", "Load", 'L', nil)
	a.mainMenu.AddItem("Save project", "Save", 'S', nil)
	a.mainMenu.AddItem("Quit", "Quit", 'Q', nil)
	a.mainMenu.AddItem("Cycles", "Cycles", 'C', a.mainMenuCycles)
}

func (a *CBAToolApp) buildCyclesPage() {
	a.cyclesPage = tview.NewTable().SetBorders(true)
	a.cyclesPage.SetBorder(true).SetTitle("Project's cycles")
	color := tcell.ColorWhite
	a.cyclesPage.SetCell(0, 0, tview.NewTableCell("Index").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.cyclesPage.SetCell(0, 1, tview.NewTableCell("Phase").SetTextColor(color).SetAlign(tview.AlignCenter))
}

func Build() *CBAToolApp {
	cbaApp := &CBAToolApp{}
	cbaApp.app = tview.NewApplication()
	cbaApp.app.EnableMouse(true)

	cbaApp.buildMainMenu()
	cbaApp.buildCyclesPage()

	cbaApp.pages = tview.NewPages()
	cbaApp.pages.AddPage("cycles", cbaApp.cyclesPage, true, true)

	flex := tview.NewFlex().AddItem(cbaApp.mainMenu, 0, 1, true)
	flex.AddItem(cbaApp.pages, 0, 3, false)

	cbaApp.app.SetInputCapture(cbaApp.eventHandler)
	cbaApp.app.SetRoot(flex, true)
	cbaApp.app.SetFocus(cbaApp.mainMenu)

	return cbaApp
}

func (a *CBAToolApp) Run() error {
	return a.app.Run()
}