package cba

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"text/template"

	log "github.com/sirupsen/logrus"
)

type SubReport struct {
	Labor         map[int]float64
	Investment    map[int]float64
	Consulting    map[int]float64
	Others        map[int]float64
	TotLabor      float64
	TotInvestment float64
	TotConsulting float64
	TotOthers     float64
	Total         float64
}

type Report struct {
	External SubReport
	Internal SubReport
	Total    float64
}

func newReport() *Report {
	r := &Report{
		External: SubReport{
			Labor:         map[int]float64{},
			Investment:    map[int]float64{},
			Consulting:    map[int]float64{},
			Others:        map[int]float64{},
			TotLabor:      0.0,
			TotInvestment: 0.0,
			TotConsulting: 0.0,
			TotOthers:     0.0,
			Total:         0.0,
		},
		Internal: SubReport{
			Labor:         map[int]float64{},
			Investment:    map[int]float64{},
			Consulting:    map[int]float64{},
			Others:        map[int]float64{},
			TotLabor:      0.0,
			TotInvestment: 0.0,
			TotConsulting: 0.0,
			TotOthers:     0.0,
			Total:         0.0,
		},
		Total: 0.0,
	}
	return r
}

func (cba *CBA) applyDiscount(name string, value float64) float64 {
	for _, d := range cba.Discounts {
		if d.Name == name {
			if d.Percentage != 0.0 {
				value = value * d.Percentage
			}
			if d.Absolute != 0.0 {
				value = value - d.Absolute
			}
		}
	}
	if value < 0.0 {
		return 0.0
	}
	return value
}

func (cba *CBA) CalcReport() Report {
	r := newReport()
	for _, p := range cba.Phases {
		r.External.Consulting[p.Index] = 0
		r.External.Investment[p.Index] = 0
		r.External.Labor[p.Index] = 0
		r.External.Others[p.Index] = 0
		r.Internal.Consulting[p.Index] = 0
		r.Internal.Investment[p.Index] = 0
		r.Internal.Labor[p.Index] = 0
		r.Internal.Others[p.Index] = 0
	}
	for _, a := range cba.Allocations {
		cost := cba.FindCostByName(a.Cost)
		phase := cba.FindPhaseByIndex(a.Phase)
		if cost != nil && phase != nil {
			value := a.getAllocatedCost(cost, phase)
			switch cost.Type {
			case COST_TYPE_CONSULTING:
				if cost.External {
					r.External.Consulting[phase.Index] += value
					r.External.TotConsulting += value
				} else {
					r.Internal.Consulting[phase.Index] += value
					r.Internal.TotConsulting += value
				}		
			case COST_TYPE_LABOR:
				if cost.External {
					r.External.Labor[phase.Index] += value
					r.External.TotLabor += value
				} else {
					r.Internal.Labor[phase.Index] += value
					r.Internal.TotLabor += value
				}
			case COST_TYPE_INVESTMENT:
				if cost.External {
					r.External.Investment[phase.Index] += value
					r.External.TotInvestment += value
				} else {
					r.Internal.Investment[phase.Index] += value
					r.Internal.TotInvestment += value
				}
			default:
				if cost.Type != COST_TYPE_OTHERS {
					log.Warningf("unrecognized type ( %s ) for cost ( %s )", cost.Type, cost.Name)
				}
				if cost.External {
					r.External.Others[phase.Index] += value
					r.External.TotOthers += value
				} else {
					r.Internal.Others[phase.Index] += value
					r.Internal.TotOthers += value
				}
			}
		}
	}
	r.External.Total = r.External.TotConsulting + r.External.TotInvestment + r.External.TotLabor + r.External.TotOthers
	r.Internal.Total = r.Internal.TotConsulting + r.Internal.TotInvestment + r.Internal.TotLabor + r.Internal.TotOthers
	r.Total = r.External.Total + r.Internal.Total
	return *r
}

func (cba *CBA) RenderingReport(report Report, templateFile string, reportFile string) error {
	log.Debugf("rendering template ( %s )", templateFile)
	name := path.Base(templateFile)
	t, err := template.New(name).ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse template ( %s ) - %v", templateFile, err)
	}
	f, err := os.Create(reportFile)
	if err != nil {
		return fmt.Errorf("failed to create target report file ( %s ) - %v", reportFile, err)
	}
	defer f.Close()
	return t.Execute(f, report)
}
