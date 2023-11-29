package business

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func (cba *CBA) PrintAnalysis() {
	t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
	header := table.Row{""}
	for _, c := range cba.Cycles {
		header = append(header, c.Name)
	}
	header = append(header, "Total")
    t.AppendHeader(header)
	labor := table.Row{"Labor"}
	investment := table.Row{"Investment"}
	consulting := table.Row{"Consulting"}
	others := table.Row{"Others"}
	for _, c := range cba.Cycles {
		labor = append(labor, fmt.Sprintf("%.2f %s", cba.Report.External.Labor[c.Index], cba.Currency))
		investment = append(investment, fmt.Sprintf("%.2f %s", cba.Report.External.Investment[c.Index], cba.Currency))
		consulting = append(consulting, fmt.Sprintf("%.2f %s", cba.Report.External.Consulting[c.Index], cba.Currency))
		others = append(others, fmt.Sprintf("%.2f %s", cba.Report.External.Others[c.Index], cba.Currency))
	}
	labor = append(labor, fmt.Sprintf("%.2f %s", cba.Report.External.TotLabor, cba.Currency))
	investment = append(investment, fmt.Sprintf("%.2f %s", cba.Report.External.TotInvestment, cba.Currency))
	consulting = append(consulting, fmt.Sprintf("%.2f %s", cba.Report.External.TotConsulting, cba.Currency))
	others = append(others, fmt.Sprintf("%.2f %s", cba.Report.External.TotOthers, cba.Currency))
    t.AppendRows([]table.Row{labor, investment, consulting, others})
    t.AppendSeparator()
    labor = table.Row{"Labor"}
	investment = table.Row{"Investment"}
	consulting = table.Row{"Consulting"}
	others = table.Row{"Others"}
	for _, c := range cba.Cycles {
		labor = append(labor, fmt.Sprintf("%.2f %s", cba.Report.Internal.Labor[c.Index], cba.Currency))
		investment = append(investment, fmt.Sprintf("%.2f %s", cba.Report.Internal.Investment[c.Index], cba.Currency))
		consulting = append(consulting, fmt.Sprintf("%.2f %s", cba.Report.Internal.Consulting[c.Index], cba.Currency))
		others = append(others, fmt.Sprintf("%.2f %s", cba.Report.Internal.Others[c.Index], cba.Currency))
	}
	labor = append(labor, fmt.Sprintf("%.2f %s", cba.Report.Internal.TotLabor, cba.Currency))
	investment = append(investment, fmt.Sprintf("%.2f %s", cba.Report.Internal.TotInvestment, cba.Currency))
	consulting = append(consulting, fmt.Sprintf("%.2f %s", cba.Report.Internal.TotConsulting, cba.Currency))
	others = append(others, fmt.Sprintf("%.2f %s", cba.Report.Internal.TotOthers, cba.Currency))
	t.AppendRows([]table.Row{labor, investment, consulting, others})
	t.Render()
}