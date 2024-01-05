package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *CBAToolApp) updateAllocationsPage() {
	color := tcell.ColorWhite
	a.allocationsPage.SetCell(0, 0, tview.NewTableCell("Cost").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.allocationsPage.SetCell(0, 1, tview.NewTableCell("Item occurrences").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.allocationsPage.SetCell(0, 2, tview.NewTableCell("Allocated to").SetTextColor(color).SetAlign(tview.AlignCenter))
	a.allocationsPage.SetCell(0, 3, tview.NewTableCell("Applied discount").SetTextColor(color).SetAlign(tview.AlignCenter))
	if a.Data == nil {
		return
	}
	for i, alloc := range a.Data.Allocations {
		cycles := ""
		for _, c := range alloc.Cycles {
			cycles += fmt.Sprint(c) + ","
		}
		a.allocationsPage.SetCell(i+1, 0, tview.NewTableCell(alloc.Cost).SetTextColor(color).SetAlign(tview.AlignCenter))
		a.allocationsPage.SetCell(i+1, 1, tview.NewTableCell(fmt.Sprint(alloc.Occurrence)).SetTextColor(color).SetAlign(tview.AlignCenter))
		a.allocationsPage.SetCell(i+1, 2, tview.NewTableCell(cycles).SetTextColor(color).SetAlign(tview.AlignCenter))
		a.allocationsPage.SetCell(i+1, 3, tview.NewTableCell(alloc.Discount).SetTextColor(color).SetAlign(tview.AlignCenter))
	}
}

func (a *CBAToolApp) buildAllocationsPage() {
	a.allocationsPage = tview.NewTable().SetBorders(true)
	a.allocationsPage.SetBorder(true).SetTitle("Allocations of costs to project's cycles")
	a.allocationsPage.SetSelectable(true, false)
	a.updateAllocationsPage()
}
