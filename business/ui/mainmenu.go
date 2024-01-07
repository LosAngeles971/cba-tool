package ui

import (
	_ "embed"

	"github.com/rivo/tview"
)

func (a *CBAToolApp) mainMenuQuit() {
	modal := tview.NewModal().SetText("Do you want to quit the application?").AddButtons([]string{"Quit", "Cancel"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Quit" {
			a.app.Stop()
		} else {
			a.app.SetRoot(a.layout, true)
		}
	})
	a.app.SetRoot(modal, true)
}

func (a *CBAToolApp) mainMenuPhases() {
	a.updateCyclesPage()
	a.pages.SwitchToPage(page_cycles)
	a.app.SetFocus(a.cyclesPage)
}

func (a *CBAToolApp) mainMenuCosts() {
	a.updateCostsPage()
	a.pages.SwitchToPage(page_costs)
	a.app.SetFocus(a.costsPage)
}

func (a *CBAToolApp) mainMenuAllocations() {
	a.updateAllocationsPage()
	a.pages.SwitchToPage(page_allocations)
	a.app.SetFocus(a.allocationsPage)
}

func (a *CBAToolApp) mainMenuReport() {
	a.updateReportPage()
	a.pages.SwitchToPage(page_report)
	a.app.SetFocus(a.reportPage)
}

func (a *CBAToolApp) loadProject() {
	//a.Data = cba.NewCBA()
	//a.mainMenuPhases()
	a.app.SetRoot(a.getBrowserPage("."), true)
}

func (a *CBAToolApp) buildMainMenu() {
	a.mainMenu = tview.NewList().ShowSecondaryText(false)
	a.mainMenu.SetBorder(true).SetTitle("Main menù")
	a.mainMenu.AddItem("Load project", "Load", 'L', a.loadProject)
	a.mainMenu.AddItem("Save project", "Save", 'S', nil)
	a.mainMenu.AddItem("Quit", "Quit", 'Q', a.mainMenuQuit)
	a.mainMenu.AddItem("Project settings", "Settings", ' ', a.callUpdateSettings)
	a.mainMenu.AddItem("Project's phases", "Project's phases", ' ', a.mainMenuPhases)
	a.mainMenu.AddItem("Costs", "List of all costs", ' ', a.mainMenuCosts)
	a.mainMenu.AddItem("Allocations", "Costs allocations", ' ', a.mainMenuAllocations)
	a.mainMenu.AddItem("Report", "CBA Report", ' ', a.mainMenuReport)
}
