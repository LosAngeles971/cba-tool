package business

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func NewCBA(data []byte) *CBA {
	cba := &CBA{
		Report: Report{
			External: SubReport{
				Labor:      map[int]float64{},
				Investment: map[int]float64{},
				Consulting: map[int]float64{},
				Others:     map[int]float64{},
			},
			Internal: SubReport{
				Labor:      map[int]float64{},
				Investment: map[int]float64{},
				Consulting: map[int]float64{},
				Others:     map[int]float64{},
			},
		},
	}
	err := yaml.Unmarshal(data, cba)
	if err != nil {
		panic(err)
	}
	return cba
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

func (cba *CBA) getCost(name string) Cost {
	for _, c := range cba.Costs {
		if c.Name == name {
			return c
		}
	}
	panic(fmt.Errorf("missing cost ( %s )", name))
}

func (cba *CBA) Calc() {
	for _, c := range cba.Cycles {
		cba.Report.External.Consulting[c.Index] = 0
		cba.Report.External.Investment[c.Index] = 0
		cba.Report.External.Labor[c.Index] = 0
		cba.Report.External.Others[c.Index] = 0
		cba.Report.External.TotConsulting = 0
		cba.Report.External.TotInvestment = 0
		cba.Report.External.TotLabor = 0
		cba.Report.External.TotOthers = 0
		cba.Report.Internal.Consulting[c.Index] = 0
		cba.Report.Internal.Investment[c.Index] = 0
		cba.Report.Internal.Labor[c.Index] = 0
		cba.Report.Internal.Others[c.Index] = 0
		cba.Report.Internal.TotConsulting = 0
		cba.Report.Internal.TotInvestment = 0
		cba.Report.Internal.TotLabor = 0
		cba.Report.Internal.TotOthers = 0
	}
	for _, allocation := range cba.Allocations {
		cost := cba.getCost(allocation.Cost)
		value := cba.applyDiscount(allocation.Discount, cost.Amount*float64(allocation.Occurrence))
		for _, index := range allocation.Cycles {
			switch cost.Type {
			case COST_TYPE_CONSULTING:
				if cost.External {
					cba.Report.External.Consulting[index] = value
					cba.Report.External.TotConsulting += value
				} else {
					cba.Report.Internal.Consulting[index] = value
					cba.Report.Internal.TotConsulting += value
				}
			case COST_TYPE_LABOR:
				if cost.External {
					cba.Report.External.Labor[index] = value
					cba.Report.External.TotLabor += value
				} else {
					cba.Report.Internal.Labor[index] = value
					cba.Report.Internal.TotLabor += value
				}
			case COST_TYPE_INVESTMENT:
				if cost.External {
					cba.Report.External.Investment[index] = value
					cba.Report.External.TotInvestment += value
				} else {
					cba.Report.Internal.Investment[index] = value
					cba.Report.Internal.TotInvestment += value
				}
			default:
				if cost.Type != COST_TYPE_OTHERS {
					log.Warningf("unrecognized type ( %s ) for cost ( %s )", cost.Type, cost.Name)
				}
				if cost.External {
					cba.Report.External.Others[index] = value
					cba.Report.External.TotOthers += value
				} else {
					cba.Report.Internal.Others[index] = value
					cba.Report.Internal.TotOthers += value
				}
			}
		}
	}
}
