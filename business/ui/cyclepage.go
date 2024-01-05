package ui

import (
	"fmt"
	"strconv"

	"github.com/LosAngeles971/cba-tool/business/cba"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *CBAToolApp) updateCyclesPage() {
	color := tcell.ColorWhite
	a.cyclesPage.SetCell(0, 0, tview.NewTableCell("Index").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.cyclesPage.SetCell(0, 1, tview.NewTableCell("Phase").SetTextColor(color).SetAlign(tview.AlignCenter))
	if a.Data == nil {
		return
	}
	for i, cycle := range a.Data.Cycles {
		a.cyclesPage.SetCell(i+1, 0, tview.NewTableCell(fmt.Sprint(cycle.Index)).SetTextColor(color).SetAlign(tview.AlignCenter))
		a.cyclesPage.SetCell(i+1, 1, tview.NewTableCell(cycle.Name).SetTextColor(color).SetAlign(tview.AlignCenter))
	}
}

func (a *CBAToolApp) buildCyclesPage() {
	a.cyclesPage = tview.NewTable().SetBorders(true)
	a.cyclesPage.SetBorder(true).SetTitle("Project's cycles")
	a.cyclesPage.SetSelectable(true, false)
	a.cyclesPage.SetSelectedFunc(func(row int, column int) {
		if row == 0 {
			a.addNewPhase()
		}
	})
	a.updateCyclesPage()
}

func (a *CBAToolApp) addNewPhase() {
	var form *tview.Form
	form = tview.NewForm().
		AddInputField("Name", "", 40, nil, nil).
		AddInputField("Index", "", 3, tview.InputFieldInteger, nil).
		AddButton("Add", func() {
			i, _ := strconv.Atoi(form.GetFormItemByLabel("Index").(*tview.InputField).GetText())
			cy := cba.Cycle{
				Name:  form.GetFormItemByLabel("Name").(*tview.InputField).GetText(),
				Index: i,
			}
			a.Data.Cycles = append(a.Data.Cycles, cy)
			a.updateAllocationsPage()
			a.app.SetFocus(a.allocationsPage)
			a.app.SetRoot(a.layout, true)
		}).AddButton("Cancel", func() {
		a.app.SetFocus(a.allocationsPage)
		a.app.SetRoot(a.layout, true)
	})
	form.SetBorder(true).SetTitle("Enter new project's phase").SetTitleAlign(tview.AlignLeft)
	a.app.SetRoot(form, true)
}
