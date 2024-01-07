package ui

import(
	_ "embed"

	"github.com/rivo/tview"
)

//go:embed about.txt
var about_text string
//go:embed help.txt
var help_text string

func (a *CBAToolApp) callMenuPage() {
	flex := tview.NewFlex()
	help := tview.NewTextView()
	help.SetText(about_text)
	help.SetTextAlign(tview.AlignLeft)
	help.SetTitle("Help")
	help.SetBorder(true)
	help.SetWordWrap(true)
	help.SetDynamicColors(true)
	mainMenu := tview.NewList().ShowSecondaryText(false)
	mainMenu.SetBorder(true).SetTitle("Main menù")
	mainMenu.AddItem("About", "", ' ', func() {
		help.SetText(about_text)
	})
	mainMenu.AddItem("Help", "", ' ', func() {
		help.SetText(help_text)
	})
	mainMenu.AddItem("Load project", "", 'L', func() {
		a.callBrowserPage()
	})
	mainMenu.AddItem("Save project", "", 'S', nil)
	mainMenu.AddItem("Quit", "", 'Q', func() {
		a.callQuit()
	})
	mainMenu.AddItem("Settings", "", ' ', a.callUpdateSettings)
	mainMenu.AddItem("Phases", "", ' ', func() {
		a.callPhasesPage()
	})
	mainMenu.AddItem("Costs", "", ' ', func() {
		a.callCostsPage()
	})
	mainMenu.AddItem("Costs allocation", "", ' ', func() {
		a.callAllocationsPage()
	})
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