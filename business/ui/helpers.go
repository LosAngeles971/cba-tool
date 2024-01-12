package ui

import (
	"strconv"

	"github.com/LosAngeles971/cba-tool/business/cba"
	"github.com/rivo/tview"
)

func (a *CBAToolApp) callPage(content tview.Primitive) {
	f := tview.NewFrame(content)
	f.SetBorder(true)
	f.SetBorderColor(t_colors[frame_border_color])
	f.SetBorders(2, 2, 1, 1, 1, 1)
	f.AddText("Cost-Benefit Analysis tool - 2024 - @LosAngeles971", true, tview.AlignLeft, t_colors[frame_header_color])
	f.AddText("Press (ESC) key to have the main menù", false, tview.AlignCenter, t_colors[frame_footer_color])
	a.app.SetRoot(f, true)
}

func getInteger(f *tview.InputField) int {
	a, err := strconv.Atoi(f.GetText())
	if err != nil {
		return 0
	}
	return a
}

func getFloat(f *tview.InputField) float64 {
	a, err := strconv.ParseFloat(f.GetText(), 64)
	if err != nil {
		return 0.0
	}
	return a
}

func getCurrencies() []string {
	return []string{cba.COST_TYPE_LABOR, cba.COST_TYPE_INVESTMENT, cba.COST_TYPE_CONSULTING, cba.COST_TYPE_OTHERS}
}

func getCurrencyIndex(c string) int {
	switch c {
	case cba.CURRENCY_DOLLAR:
		return 1
	default:
		return 0
	}
}

func getCostTypes() []string {
	return []string{cba.COST_TYPE_LABOR, cba.COST_TYPE_INVESTMENT, cba.COST_TYPE_CONSULTING, cba.COST_TYPE_OTHERS}
}

func (a *CBAToolApp) getCostIndex(name string) int {
	for i, c := range a.Data.Costs {
		if c.Name == name {
			return i
		}
	}
	return 0
}
