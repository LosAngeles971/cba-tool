package ui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/LosAngeles971/cba-tool/business/cba"
)

type CBAToolApp struct {
	Data *cba.CBA
	mainWindow fyne.Window
}

func getTableLabel() *widget.Label {
	l := widget.NewLabel("wide content")
	l.Truncation = fyne.TextTruncateClip
	return l
}

func (a *CBAToolApp) getCyclesTab() *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(a.Data.Cycles) + 1, 2
		},
		func() fyne.CanvasObject {
			return getTableLabel()
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			if i.Row < 1 {
				label.TextStyle = fyne.TextStyle{Bold: true}
				if i.Col == 0 {
					label.SetText("Index")
				} else {
					label.SetText("Phase")	
				}
			} else {
				cycle := a.Data.Cycles[i.Row - 1]
				if i.Col == 0 {
					label.SetText(fmt.Sprint(cycle.Index))
				} else {
					label.SetText(cycle.Name)	
				}
			}
		})
	table.SetColumnWidth(0, widget.NewLabel("Index 999").MinSize().Width)
	table.SetColumnWidth(1, widget.NewLabel("Year number 1").MinSize().Width)
	return table
}

func (a *CBAToolApp) getCostsTab() *widget.Table {
	return widget.NewTable(
		func() (int, int) {
			return len(a.Data.Costs), 5
		},
		func() fyne.CanvasObject {
			return getTableLabel()
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			cost := a.Data.Costs[i.Row]
			switch i.Col {
			case 0:
				o.(*widget.Label).SetText(fmt.Sprint(cost.Name))
			case 1:
				o.(*widget.Label).SetText(fmt.Sprint(cost.Amount))
			case 2:
				o.(*widget.Label).SetText(fmt.Sprint(cost.Type))
			case 3:
				o.(*widget.Label).SetText(fmt.Sprint(cost.External))
			case 4:
				o.(*widget.Label).SetText(fmt.Sprint(cost.Description))
			}
		})
}

func (a *CBAToolApp) Build() {
	mainApp := app.New()
	a.mainWindow = mainApp.NewWindow("CBA tool - LosAngeles971@2023")
	a.mainWindow.Resize(fyne.NewSize(1000, 640))
	tabs := container.NewAppTabs(
		container.NewTabItem("Cycles", a.getCyclesTab()),
		container.NewTabItem("Discounts", widget.NewLabel("World!")),
		container.NewTabItem("Costs", a.getCostsTab()),
		container.NewTabItem("Allocations", widget.NewLabel("World!")),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)
	border := container.NewBorder(toolbar, nil, nil, nil, tabs)
	a.mainWindow.SetContent(border)
}

func (a *CBAToolApp) Run() {
	a.mainWindow.ShowAndRun()
}