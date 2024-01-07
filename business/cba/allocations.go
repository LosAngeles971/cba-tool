package cba

import (
	"github.com/google/uuid"
)

type Allocation struct {
	ID         string  `json:"id" yaml:"id"`
	Cost       string  `json:"cost" yaml:"cost"`
	Occurrence float64 `json:"occurrence" yaml:"occurrence"` // float64 because for labor cost you may have 1.3 FTE for example
	Phase      int     `json:"phase" yaml:"phase"`
	Discount   string  `json:"discount" yaml:"discount"`
}

type Phase struct {
	Name  string `json:"name" yaml:"name"`
	Index int    `json:"index" yaml:"index"`
	Days  int    `json:"days" yaml:"days"`
}

func (cba *CBA) NewAllocation() *Allocation {
	return &Allocation{
		ID:         uuid.NewString(),
		Occurrence: 0,
	}
}

func (cba *CBA) NewPhase() *Phase {
	cba.sortPhases()
	return &Phase{
		Name:  "",
		Index: len(cba.Phases),
		Days:  365,
	}
}

func (cba *CBA) FindPhaseByIndex(i int) *Phase {
	for _, p := range cba.Phases {
		if p.Index == i {
			return p
		}
	}
	return nil
}

func (cba *CBA) FindPhaseByName(name string) *Phase {
	for _, p := range cba.Phases {
		if p.Name == name {
			return p
		}
	}
	return nil
}

func (cba *CBA) DeletePhaseByIndex(i int) {
	buffer := []*Phase{}
	for _, p := range cba.Phases {
		if p.Index != i {
			buffer = append(buffer, p)
		}
	}
	cba.Phases = buffer
}

func (cba *CBA) ListPhases() []string {
	pp := []string{}
	for _, p := range cba.Phases {
		pp = append(pp, p.Name)
	}
	return pp
}

func (a *Allocation) getAllocatedCost(cost *Cost, phase *Phase) float64 {
	var value float64
	totamount := cost.Amount * float64(a.Occurrence)
	switch cost.Metric {
	case COST_METRIC_DAILY:
		value = totamount * float64(phase.Days)
	case COST_METRIC_YEARLY:
		value = (totamount * float64(phase.Days)) / 365
	default: // once
		value = totamount
	}
	return value
}

func (cba *CBA) DeleteAllocationByID(id string) {
	buffer := []*Allocation{}
	for _, a := range cba.Allocations {
		if a.ID != id {
			buffer = append(buffer, a)
		}
	}
	cba.Allocations = buffer
}

func (cba *CBA) FindAllocationByID(id string) *Allocation {
	for _, a := range cba.Allocations {
		if a.ID == id {
			return a
		}
	}
	return nil
}
