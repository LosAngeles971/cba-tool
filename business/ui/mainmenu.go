package ui

func (a *CBAToolApp) mainMenuCycles() {
	a.pages.SwitchToPage("cycles")
	a.app.SetFocus(a.cyclesPage)
}