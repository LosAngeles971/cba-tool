package ui

import(
	_ "embed"

	"github.com/rivo/tview"
)

//go:embed about.txt
var about_test string

func (a *CBAToolApp) callMenuPage() {
	flex := tview.NewFlex()
	help := tview.NewTextView()
	help.SetText(about_test)
	help.SetTextAlign(tview.AlignLeft)
	help.SetTitle("Help")
	help.SetBorder(true)
	help.SetWordWrap(true)
	mainMenu := tview.NewList().ShowSecondaryText(false)
	mainMenu.SetBorder(true).SetTitle("Main menù")
	mainMenu.AddItem("About", "", ' ', func() {
		help.SetText(about_test)
	})
	mainMenu.AddItem("Load project", "", 'L', a.loadProject)
	mainMenu.AddItem("Save project", "", 'S', nil)
	mainMenu.AddItem("Quit", "", 'Q', a.mainMenuQuit)
	mainMenu.AddItem("Settings", "", ' ', a.callUpdateSettings)
	mainMenu.AddItem("Phases", "", ' ', a.mainMenuPhases)
	mainMenu.AddItem("Costs", "", ' ', func() {
		a.callCostsPage()
	})
	mainMenu.AddItem("Allocations", "", ' ', a.mainMenuAllocations)
	mainMenu.AddItem("Report", "", ' ', func() {
		a.callReportPage()
	})
	mainMenu.AddItem("Summary", "A summary of the phase by phase total costs", ' ', func() {
		a.callSummaryPage()
	})
	flex.AddItem(mainMenu, 0, 1, true)
	flex.AddItem(help, 0, 3, false)
	a.app.SetRoot(flex, true)
}