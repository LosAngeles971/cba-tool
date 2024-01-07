package ui

import (
	"fmt"

	"github.com/LosAngeles971/cba-tool/business/cba"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *CBAToolApp) updateAllocationsPage() {
	color := tcell.ColorWhite
	a.allocationsPage.Clear()
	a.allocationsPage.SetCell(0, 0, tview.NewTableCell("Cost").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.allocationsPage.SetCell(0, 1, tview.NewTableCell("Item occurrences").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.allocationsPage.SetCell(0, 2, tview.NewTableCell("Allocated to").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.allocationsPage.SetCell(0, 3, tview.NewTableCell("Applied discount").SetTextColor(color).SetAlign(tview.AlignCenter))
	if a.Data == nil {
		return
	}
	for i, alloc := range a.Data.Allocations {
		a.allocationsPage.SetCell(i+1, 0, tview.NewTableCell(alloc.Cost).SetTextColor(color).SetAlign(tview.AlignCenter))
		a.allocationsPage.SetCell(i+1, 1, tview.NewTableCell(fmt.Sprint(alloc.Occurrence)).SetTextColor(color).SetAlign(tview.AlignCenter))
		a.allocationsPage.SetCell(i+1, 2, tview.NewTableCell(fmt.Sprint(alloc.Phase)).SetTextColor(color).SetAlign(tview.AlignCenter))
		a.allocationsPage.SetCell(i+1, 3, tview.NewTableCell(alloc.Discount).SetTextColor(color).SetAlign(tview.AlignCenter))
	}
}

func (a *CBAToolApp) buildAllocationsPage() {
	a.allocationsPage = tview.NewTable().SetBorders(true)
	a.allocationsPage.SetBorder(true).SetTitle("Allocations of costs to project's cycles")
	a.allocationsPage.SetSelectable(true, false)
	a.allocationsPage.SetSelectedFunc(func(row int, column int) {
		if row < 1 {
			a.addNewAllocation()
		} else {
			a.updateAllocation(a.Data.Costs[row-1].Name)
		}
	})
	a.updateAllocationsPage()
}

func (a *CBAToolApp) getAllocationForm(alloc *cba.Allocation, update bool) *tview.Form {
	form := tview.NewForm()
	form.AddDropDown("Associated cost", a.Data.ListCosts(), a.getCostIndex(alloc.Cost), nil)
	form.AddInputField("Cost occurrences", fmt.Sprint(alloc.Occurrence), 15, tview.InputFieldInteger, nil)
	form.AddDropDown("Associated phase", a.Data.ListPhases(), alloc.Phase, nil)
	//form.AddDropDown("Applied discount", getCurrencies(), getCurrencyIndex(cost.Currency), nil)
	form.AddButton("Add/Update", func() {
		_, c := form.GetFormItemByLabel("Associated cost").(*tview.DropDown).GetCurrentOption()
		_, n := form.GetFormItemByLabel("Associated phase").(*tview.DropDown).GetCurrentOption()
		p := a.Data.FindPhaseByName(n)
		if p != nil {
			alloc.Cost = c
			alloc.Phase = p.Index
			alloc.Occurrence = getFloat(form.GetFormItemByLabel("Cost occurrences").(*tview.InputField))
			a.updateAllocationsPage()
			a.app.SetFocus(a.allocationsPage)
			a.app.SetRoot(a.layout, true)
		}
	})
	form.AddButton("Cancel", func() {
		a.app.SetFocus(a.allocationsPage)
		a.app.SetRoot(a.layout, true)
	})
	if update {
		form.AddButton("Delete", func() {
			a.Data.DeleteAllocationByID(alloc.ID)
			a.updateAllocationsPage()
			a.app.SetFocus(a.allocationsPage)
		a.app.SetRoot(a.layout, true)
		})
	}
	return form
}

func (a *CBAToolApp) updateAllocation(id string) {
	alloc := a.Data.FindAllocationByID(id)
	if alloc == nil {
		return
	}
	form := a.getAllocationForm(alloc, true)
	form.SetBorder(true).SetTitle("Update cost allocation").SetTitleAlign(tview.AlignLeft)
	a.app.SetRoot(form, true)
}

func (a *CBAToolApp) addNewAllocation() {
	alloc := a.Data.NewAllocation()
	form := a.getAllocationForm(alloc, false)
	form.SetBorder(true).SetTitle("Add cost allocation").SetTitleAlign(tview.AlignLeft)
	a.app.SetRoot(form, true)
}