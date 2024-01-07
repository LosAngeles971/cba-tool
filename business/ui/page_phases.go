package ui

import (
	"fmt"

	"github.com/LosAngeles971/cba-tool/business/cba"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *CBAToolApp) updateCyclesPage() {
	color := tcell.ColorWhite
	a.cyclesPage.Clear()
	a.cyclesPage.SetCell(0, 0, tview.NewTableCell("Index").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.cyclesPage.SetCell(0, 1, tview.NewTableCell("Phase").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.cyclesPage.SetCell(0, 2, tview.NewTableCell("Days").SetTextColor(color).SetAlign(tview.AlignCenter))
	if a.Data == nil {
		return
	}
	for i, cycle := range a.Data.Phases {
		a.cyclesPage.SetCell(i+1, 0, tview.NewTableCell(fmt.Sprint(cycle.Index)).SetTextColor(color).SetAlign(tview.AlignCenter))
		a.cyclesPage.SetCell(i+1, 1, tview.NewTableCell(cycle.Name).SetTextColor(color).SetAlign(tview.AlignCenter))
		a.cyclesPage.SetCell(i+1, 2, tview.NewTableCell(fmt.Sprint(cycle.Days)).SetTextColor(color).SetAlign(tview.AlignCenter))
	}
}

func (a *CBAToolApp) buildCyclesPage() {
	a.cyclesPage = tview.NewTable().SetBorders(true)
	a.cyclesPage.SetBorder(true).SetTitle("Project's phases")
	a.cyclesPage.SetSelectable(true, false)
	a.cyclesPage.SetSelectedFunc(func(row int, column int) {
		if row < 1 {
			a.addNewPhase()
		} else {
			a.updatePhase(a.Data.Phases[row-1].Index)
		}
	})
	a.updateCyclesPage()
}

func (a *CBAToolApp) getPhaseForm(phase *cba.Phase, update bool) *tview.Form {
	form := tview.NewForm()
	form.AddInputField("Name", phase.Name, 40, nil, nil)
	form.AddInputField("Index", fmt.Sprint(phase.Index), 3, tview.InputFieldInteger, nil)
	form.AddInputField("Days", fmt.Sprint(phase.Days), 10, tview.InputFieldInteger, nil)
	form.AddButton("Add/Update", func() {
		phase.Name =  form.GetFormItemByLabel("Name").(*tview.InputField).GetText()
		phase.Index = getInteger(form.GetFormItemByLabel("Index").(*tview.InputField))
		phase.Days = getInteger(form.GetFormItemByLabel("Days").(*tview.InputField))
		if !update {
			a.Data.Phases = append(a.Data.Phases, phase)
		}
		a.updateAllocationsPage()
		a.app.SetFocus(a.allocationsPage)
		a.app.SetRoot(a.layout, true)
	})
	form.AddButton("Cancel", func() {
		a.app.SetFocus(a.costsPage)
		a.app.SetRoot(a.layout, true)
	})
	if update {
		form.AddButton("Delete", func() {
			a.Data.DeletePhaseByIndex(phase.Index)
			//a.updateCostsPage()
			a.app.SetFocus(a.costsPage)
			a.app.SetRoot(a.layout, true)
		})
	}
	return form
}

func (a *CBAToolApp) updatePhase(i int) {
	phase := a.Data.FindPhaseByIndex(i)
	if phase == nil {
		return
	}
	form := a.getPhaseForm(phase, true)
	form.SetBorder(true).SetTitle("Update phase").SetTitleAlign(tview.AlignLeft)
	a.app.SetRoot(form, true)
}

func (a *CBAToolApp) addNewPhase() {
	phase := a.Data.NewPhase()
	form := a.getPhaseForm(phase, false)
	form.SetBorder(true).SetTitle("Enter new project's phase").SetTitleAlign(tview.AlignLeft)
	a.app.SetRoot(form, true)
}
