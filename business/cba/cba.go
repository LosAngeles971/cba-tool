package cba

import (
	_ "embed"
	"os"
	"sort"

	"gopkg.in/yaml.v3"
)

//go:embed example.yaml
var example []byte

//go:embed common_costs.yaml
var catalog []byte

const (
	CURRENCY_EURO          = "euro"
	CURRENCY_DOLLAR        = "dollar"
	CURRENCY_EURO_SYMBOL   = "€"
	CURRENCY_DOLLAR_SYMBOL = "$"
)

type Discount struct {
	Name       string  `json:"name" yaml:"name"`
	Percentage float64 `json:"percentage" yaml:"percentage"`
	Absolute   float64 `json:"absolute" yaml:"absolute"`
}

type CBA struct {
	Phases        []*Phase      `json:"phases" yaml:"phases"`
	Discounts     []*Discount   `json:"discounts" yaml:"discounts"`
	Costs         []*Cost       `json:"costs" yaml:"costs"`
	Allocations   []*Allocation `json:"allocations" yaml:"allocations"`
	Currency      string        `json:"currency" yaml:"currency"`
	ValueAddedTax float64       `json:"vat" yaml:"vat"`
	VAT           bool          `json:"vat_enable" yaml:"vat_enable"`
}

func NewCBA() *CBA {
	cba := &CBA{
		Phases:        []*Phase{},
		Discounts:     []*Discount{},
		Costs:         []*Cost{},
		Allocations:   []*Allocation{},
		Currency:      CURRENCY_EURO,
		ValueAddedTax: 0,
		VAT:           false,
	}
	// TEMPORARY PRE-LOAD
	cba.Load(example)
	cba.LoadCostsFromCatalog(catalog)
	return cba
}

func (cba *CBA) sortPhases() {
	sort.Slice(cba.Phases, func(i, j int) bool {
		return cba.Phases[i].Index < cba.Phases[j].Index
	})
}

func (cba *CBA) Load(data []byte) {
	err := yaml.Unmarshal(data, cba)
	if err != nil {
		panic(err)
	}
	cba.sortPhases()
}

func (cba *CBA) LoadFile(path string) error {
	data, err := os.ReadFile(path)
	if err == nil {
		cba.Load(data)
	}
	return err
}

func (cba *CBA) LoadCostsFromCatalog(data []byte) error {
	cc := CatalogOfCosts{}
	err := yaml.Unmarshal(data, &cc)
	if err != nil {
		return err
	}
	for _, c := range cc.Catalog {
		c.ReadOnly = true
		e := cba.FindCostByName(c.Name)
		if e == nil {
			cba.Costs = append(cba.Costs, c)
		}
	}
	return nil
}

func (cba *CBA) LoadCostsFromCatalogFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return cba.LoadCostsFromCatalog(data)
}
