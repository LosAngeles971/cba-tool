package ui

import (
	"fmt"

	"github.com/LosAngeles971/cba-tool/business/cba"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *CBAToolApp) callCostsPage() {
	costsPage := tview.NewTable().SetBorders(true)
	costsPage.SetTitle("List of all possible costs").SetBorder(true)
	costsPage.SetBordersColor(tcell.ColorYellow)
	costsPage.SetSelectable(true, false)
	costsPage.SetSelectedFunc(func(row int, column int) {
		if row < 1 {
			a.callAddNewCost()
		} else {
			a.callUpdateCost(a.Data.Costs[row-1].Name)
		}
	})
	costsPage.Clear()
	costsPage.SetCell(0, 0, tview.NewTableCell("Name").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	costsPage.SetCell(0, 1, tview.NewTableCell("Type").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	costsPage.SetCell(0, 2, tview.NewTableCell("Amount").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	costsPage.SetCell(0, 3, tview.NewTableCell("External").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	if a.Data != nil {
		for i, cost := range a.Data.Costs {
			costsPage.SetCell(i+1, 0, tview.NewTableCell(cost.Name).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
			costsPage.SetCell(i+1, 1, tview.NewTableCell(cost.Type).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
			costsPage.SetCell(i+1, 2, tview.NewTableCell(fmt.Sprintf("%f %s", cost.Amount, cost.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
			costsPage.SetCell(i+1, 3, tview.NewTableCell(fmt.Sprint(cost.External)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
		}
	}
	a.app.SetRoot(costsPage, true)
}

// func (a *CBAToolApp) updateCostsPage() {
// 	color := tcell.ColorWhite
// 	a.costsPage.Clear()
// 	a.costsPage.SetCell(0, 0, tview.NewTableCell("Name").SetTextColor(color).SetAlign(tview.AlignCenter))
// 	a.costsPage.SetCell(0, 1, tview.NewTableCell("Type").SetTextColor(color).SetAlign(tview.AlignCenter))
// 	a.costsPage.SetCell(0, 2, tview.NewTableCell("Amount").SetTextColor(color).SetAlign(tview.AlignCenter))
// 	a.costsPage.SetCell(0, 3, tview.NewTableCell("External").SetTextColor(color).SetAlign(tview.AlignCenter))
// 	if a.Data == nil {
// 		return
// 	}
// 	for i, cost := range a.Data.Costs {
// 		a.costsPage.SetCell(i+1, 0, tview.NewTableCell(cost.Name).SetTextColor(color).SetAlign(tview.AlignCenter))
// 		a.costsPage.SetCell(i+1, 1, tview.NewTableCell(cost.Type).SetTextColor(color).SetAlign(tview.AlignCenter))
// 		a.costsPage.SetCell(i+1, 2, tview.NewTableCell(fmt.Sprintf("%f %s", cost.Amount, cost.Currency)).SetTextColor(color).SetAlign(tview.AlignCenter))
// 		a.costsPage.SetCell(i+1, 3, tview.NewTableCell(fmt.Sprint(cost.External)).SetTextColor(color).SetAlign(tview.AlignCenter))
// 	}
// }

// func (a *CBAToolApp) buildCostPage() {
// 	a.costsPage = tview.NewTable().SetBorders(true)
// 	a.costsPage.SetTitle("List of all possible costs").SetBorder(true)
// 	a.costsPage.SetBordersColor(tcell.ColorYellow)
// 	a.costsPage.SetSelectable(true, false)
// 	a.costsPage.SetSelectedFunc(func(row int, column int) {
// 		if row < 1 {
// 			a.addNewCost()
// 		} else {
// 			a.updateCost(a.Data.Costs[row-1].Name)
// 		}
// 	})
// 	a.updateCostsPage()
// }

func (a *CBAToolApp) getCostForm(cost *cba.Cost, update bool) *tview.Form {
	form := tview.NewForm()
	form.AddInputField("Name", cost.Name, 40, nil, nil)
	form.AddInputField("Metric", cost.Metric, 80, nil, nil)
	form.AddDropDown("Type", getCostTypes(), 0, nil)
	form.AddDropDown("Currency", getCurrencies(), getCurrencyIndex(cost.Currency), nil)
	form.AddInputField("Amount", fmt.Sprint(cost.Amount), 30, tview.InputFieldFloat, nil)
	bText := "Add"
	if update {
		bText = "Update"
	}
	form.AddCheckbox("External", cost.External, nil)
	form.AddButton(bText, func() {
		_, t := form.GetFormItemByLabel("Type").(*tview.DropDown).GetCurrentOption()
		_, c := form.GetFormItemByLabel("Currency").(*tview.DropDown).GetCurrentOption()
		cost.Name = form.GetFormItemByLabel("Name").(*tview.InputField).GetText()
		cost.Metric = form.GetFormItemByLabel("Metric").(*tview.InputField).GetText()
		cost.Type = t
		cost.Currency = c
		cost.Amount = getFloat(form.GetFormItemByLabel("Amount").(*tview.InputField))
		cost.External = form.GetFormItemByLabel("External").(*tview.Checkbox).IsChecked()
		if !update {
			a.Data.Costs = append(a.Data.Costs, cost)
		}
		//a.updateCostsPage()
		//a.app.SetFocus(a.costsPage)
		//a.app.SetRoot(a.layout, true)
		a.callCostsPage()
	})
	form.AddButton("Cancel", func() {
		//a.app.SetFocus(a.costsPage)
		//a.app.SetRoot(a.layout, true)
		a.callCostsPage()
	})
	if update {
		form.AddButton("Delete", func() {
			a.Data.DeleteCostByName(cost.Name)
			//a.updateCostsPage()
			//a.app.SetFocus(a.costsPage)
			//a.app.SetRoot(a.layout, true)
			a.callCostsPage()
		})
	}
	return form
}

func (a *CBAToolApp) callUpdateCost(name string) {
	cost := a.Data.FindCostByName(name)
	if cost == nil {
		return
	}
	form := a.getCostForm(cost, true)
	form.SetBorder(true).SetTitle("Update cost item").SetTitleAlign(tview.AlignLeft)
	a.app.SetRoot(form, true)
}

func (a *CBAToolApp) callAddNewCost() {
	cost := a.Data.NewCost()
	form := a.getCostForm(cost, false)
	form.SetBorder(true).SetTitle("Add cost item").SetTitleAlign(tview.AlignLeft)
	a.app.SetRoot(form, true)
}