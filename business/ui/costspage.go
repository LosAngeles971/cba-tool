package ui

import (
	"fmt"

	"github.com/LosAngeles971/cba-tool/business/cba"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *CBAToolApp) updateCostsPage() {
	color := tcell.ColorWhite
	a.costsPage.SetCell(0, 0, tview.NewTableCell("Name").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.costsPage.SetCell(0, 1, tview.NewTableCell("Type").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.costsPage.SetCell(0, 2, tview.NewTableCell("Amount").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.costsPage.SetCell(0, 3, tview.NewTableCell("External").SetTextColor(color).SetAlign(tview.AlignCenter))
	if a.Data == nil {
		return
	}
	for i, cost := range a.Data.Costs {
		a.costsPage.SetCell(i+1, 0, tview.NewTableCell(cost.Name).SetTextColor(color).SetAlign(tview.AlignCenter))
		a.costsPage.SetCell(i+1, 1, tview.NewTableCell(cost.Type).SetTextColor(color).SetAlign(tview.AlignCenter))
		a.costsPage.SetCell(i+1, 2, tview.NewTableCell(fmt.Sprintf("%f %s", cost.Amount, cost.Currency)).SetTextColor(color).SetAlign(tview.AlignCenter))
		a.costsPage.SetCell(i+1, 3, tview.NewTableCell(fmt.Sprint(cost.External)).SetTextColor(color).SetAlign(tview.AlignCenter))
	}
}

func (a *CBAToolApp) buildCostPage() {
	a.costsPage = tview.NewTable().SetBorders(true)
	a.costsPage.SetTitle("List of all possible costs").SetBorder(true)
	a.costsPage.SetBordersColor(tcell.ColorYellow)
	a.costsPage.SetSelectable(true, false)
	a.costsPage.SetSelectedFunc(func(row int, column int) {
		if row == 0 {
			a.addNewCost()
		}
	})
	a.updateCostsPage()
}

func (a *CBAToolApp) addNewCost() {
	var form *tview.Form
	form = tview.NewForm().
		AddInputField("Name", "", 40, nil, nil).
		AddInputField("Description", "", 80, nil, nil).
		AddDropDown("Type", []string{"labor", "investment", "consulting", "others"}, 0, nil).
		AddDropDown("Currency", []string{"euro", "dollaro"}, 0, nil).
		AddInputField("Amount", "", 30, tview.InputFieldFloat, nil).
		AddCheckbox("External", false, nil).
		AddButton("Add", func() {
			_, t := form.GetFormItemByLabel("Type").(*tview.DropDown).GetCurrentOption()
			_, c := form.GetFormItemByLabel("Currency").(*tview.DropDown).GetCurrentOption()
			cost := cba.Cost{
				Name:        form.GetFormItemByLabel("Name").(*tview.InputField).GetText(),
				Description: form.GetFormItemByLabel("Description").(*tview.InputField).GetText(),
				Type:        t,
				Currency:    c,
				Amount:      0.0,
				Tags:        []string{},
				External:    false,
			}
			if form.GetFormItemByLabel("External").(*tview.Checkbox).IsChecked() {
				cost.External = true
			}
			a.Data.Costs = append(a.Data.Costs, cost)
			a.updateCostsPage()
			a.app.SetFocus(a.costsPage)
			a.app.SetRoot(a.layout, true)
		}).AddButton("Cancel", func() {
		a.app.SetFocus(a.costsPage)
		a.app.SetRoot(a.layout, true)
	})
	form.SetBorder(true).SetTitle("Enter new cost").SetTitleAlign(tview.AlignLeft)
	a.app.SetRoot(form, true)
}
