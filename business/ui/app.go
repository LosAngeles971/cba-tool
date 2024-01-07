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

const (
	page_cycles      = "cycles"
	page_costs       = "costs"
	page_allocations = "allocations"
	page_report      = "report"
)

type CBAToolApp struct {
	Data            *cba.CBA
	app             *tview.Application
	mainMenu        *tview.List
	pages           *tview.Pages
	cyclesPage      *tview.Table
	costsPage       *tview.Table
	allocationsPage *tview.Table
	reportPage      *tview.Table
	layout          *tview.Flex
}

func (a *CBAToolApp) eventHandler(eventKey *tcell.EventKey) *tcell.EventKey {
	if eventKey.Key() == tcell.KeyEscape {
		a.app.SetRoot(a.layout, true)
		a.app.SetFocus(a.mainMenu)
	}
	return eventKey
}

func Build() *CBAToolApp {
	cbaApp := &CBAToolApp{
		Data: cba.NewCBA(),
	}
	cbaApp.app = tview.NewApplication()
	cbaApp.app.EnableMouse(true)

	cbaApp.buildMainMenu()
	cbaApp.buildCyclesPage()
	cbaApp.buildCostPage()
	cbaApp.buildAllocationsPage()
	cbaApp.buildReportPage()

	cbaApp.pages = tview.NewPages()
	cbaApp.pages.AddPage(page_cycles, cbaApp.cyclesPage, true, true)
	cbaApp.pages.AddPage(page_costs, cbaApp.costsPage, true, false)
	cbaApp.pages.AddPage(page_allocations, cbaApp.allocationsPage, true, false)
	cbaApp.pages.AddPage(page_report, cbaApp.reportPage, true, false)

	cbaApp.layout = tview.NewFlex().AddItem(cbaApp.mainMenu, 0, 1, true)
	cbaApp.layout.AddItem(cbaApp.pages, 0, 3, false)

	cbaApp.app.SetInputCapture(cbaApp.eventHandler)
	cbaApp.app.SetRoot(cbaApp.layout, true)
	cbaApp.app.SetFocus(cbaApp.mainMenu)

	return cbaApp
}

func (a *CBAToolApp) Run() error {
	return a.app.Run()
}
