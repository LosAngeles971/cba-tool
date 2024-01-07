package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *CBAToolApp) updateReportPage() {
	r := a.Data.CalcReport()
	if r.Total == 0 && len(a.Data.Allocations) >0 {
		panic(fmt.Errorf("%v", r))
	}
	// HEADER
	a.reportPage.SetCell(0, 0, tview.NewTableCell("").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	for i, c := range a.Data.Phases {
		a.reportPage.SetCell(0, 1 + i, tview.NewTableCell(c.Name).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	}
	a.reportPage.SetCell(0, 1 + len(a.Data.Phases), tview.NewTableCell("Total").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	// end of HEADER
	// External LABOR
	a.reportPage.SetCell(1, 0, tview.NewTableCell("External - labor").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	for i, c := range a.Data.Phases {
		a.reportPage.SetCell(1, 1 + i, tview.NewTableCell(fmt.Sprintf("%.2f %s", r.External.Labor[c.Index], a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	}
	a.reportPage.SetCell(1, 1 + len(a.Data.Phases), tview.NewTableCell(fmt.Sprintf("%.2f %s", r.External.TotLabor, a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	// end of External LABOR
	// External INVESTMENT
	a.reportPage.SetCell(2, 0, tview.NewTableCell("External - investment").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	for i, c := range a.Data.Phases {
		a.reportPage.SetCell(2, 1 + i, tview.NewTableCell(fmt.Sprintf("%.2f %s", r.External.Investment[c.Index], a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	}
	a.reportPage.SetCell(2, 1 + len(a.Data.Phases), tview.NewTableCell(fmt.Sprintf("%.2f %s", r.External.TotInvestment, a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	// end of External INVESTMENT
	// External CONSULTING
	a.reportPage.SetCell(3, 0, tview.NewTableCell("External - consulting").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	for i, c := range a.Data.Phases {
		a.reportPage.SetCell(3, 1 + i, tview.NewTableCell(fmt.Sprintf("%.2f %s", r.External.Consulting[c.Index], a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	}
	a.reportPage.SetCell(3, 1 + len(a.Data.Phases), tview.NewTableCell(fmt.Sprintf("%.2f %s", r.External.TotConsulting, a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	// end of External CONSULTING
	// External OTHERS
	a.reportPage.SetCell(4, 0, tview.NewTableCell("External - others").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	for i, c := range a.Data.Phases {
		a.reportPage.SetCell(4, 1 + i, tview.NewTableCell(fmt.Sprintf("%.2f %s", r.External.Others[c.Index], a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	}
	a.reportPage.SetCell(4, 1 + len(a.Data.Phases), tview.NewTableCell(fmt.Sprintf("%.2f %s", r.External.TotOthers, a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	// end of External OTHERS
	// Internal LABOR
	a.reportPage.SetCell(5, 0, tview.NewTableCell("Internal - labor").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	for i, c := range a.Data.Phases {
		a.reportPage.SetCell(5, 1 + i, tview.NewTableCell(fmt.Sprintf("%.2f %s", r.Internal.Labor[c.Index], a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	}
	a.reportPage.SetCell(5, 1 + len(a.Data.Phases), tview.NewTableCell(fmt.Sprintf("%.2f %s", r.Internal.TotLabor, a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	// end of Internal LABOR
	// Internal INVESTMENT
	a.reportPage.SetCell(6, 0, tview.NewTableCell("Internal - investment").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	for i, c := range a.Data.Phases {
		a.reportPage.SetCell(6, 1 + i, tview.NewTableCell(fmt.Sprintf("%.2f %s", r.Internal.Investment[c.Index], a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	}
	a.reportPage.SetCell(6, 1 + len(a.Data.Phases), tview.NewTableCell(fmt.Sprintf("%.2f %s", r.Internal.TotInvestment, a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	// end of Internal INVESTMENT
	// Internal CONSULTING
	a.reportPage.SetCell(7, 0, tview.NewTableCell("Internal - consulting").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	for i, c := range a.Data.Phases {
		a.reportPage.SetCell(7, 1 + i, tview.NewTableCell(fmt.Sprintf("%.2f %s", r.Internal.Consulting[c.Index], a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	}
	a.reportPage.SetCell(7, 1 + len(a.Data.Phases), tview.NewTableCell(fmt.Sprintf("%.2f %s", r.Internal.TotConsulting, a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	// end of Internal CONSULTING
	// Internal OTHERS
	a.reportPage.SetCell(8, 0, tview.NewTableCell("Internal - others").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	for i, c := range a.Data.Phases {
		a.reportPage.SetCell(8, 1 + i, tview.NewTableCell(fmt.Sprintf("%.2f %s", r.Internal.Others[c.Index], a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	}
	a.reportPage.SetCell(8, 1 + len(a.Data.Phases), tview.NewTableCell(fmt.Sprintf("%.2f %s", r.Internal.TotOthers, a.Data.Currency)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter))
	// end of Internal OTHERS
}

func (a *CBAToolApp) buildReportPage() {
	a.reportPage = tview.NewTable().SetBorders(true)
	a.reportPage.SetBorder(true).SetTitle("Cost-Benefit Analysis report")
	a.reportPage.SetSelectable(false, false)
	a.updateReportPage()
}