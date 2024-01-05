package ui

import (
	_ "embed"

	"github.com/LosAngeles971/cba-tool/business/cba"
	"github.com/rivo/tview"
)

//go:embed example.yaml
var example []byte

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

func (a *CBAToolApp) mainMenuCycles() {
	a.pages.SwitchToPage(page_cycles)
	a.app.SetFocus(a.cyclesPage)
}

func (a *CBAToolApp) mainMenuCosts() {
	a.pages.SwitchToPage(page_costs)
	a.app.SetFocus(a.costsPage)
}

func (a *CBAToolApp) mainMenuAllocations() {
	a.pages.SwitchToPage(page_allocations)
	a.app.SetFocus(a.allocationsPage)
}

func (a *CBAToolApp) loadProject() {
	a.Data = cba.NewCBA(example)
	a.updateCyclesPage()
	a.updateCostsPage()
	a.updateAllocationsPage()
	a.pages.SwitchToPage("cycles")
}

func (a *CBAToolApp) buildMainMenu() {
	a.mainMenu = tview.NewList().ShowSecondaryText(false)
	a.mainMenu.SetBorder(true).SetTitle("Main menù")
	a.mainMenu.AddItem("Load project", "Load", 'L', a.loadProject)
	a.mainMenu.AddItem("Save project", "Save", 'S', nil)
	a.mainMenu.AddItem("Quit", "Quit", 'Q', a.mainMenuQuit)
	a.mainMenu.AddItem("Cycle", "Project's cycle", ' ', a.mainMenuCycles)
	a.mainMenu.AddItem("Costs", "List of all costs", ' ', a.mainMenuCosts)
	a.mainMenu.AddItem("Allocations", "Costs allocations", ' ', a.mainMenuAllocations)
}
