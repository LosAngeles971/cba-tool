package cba

const (
	COST_TYPE_LABOR      = "labor"
	COST_TYPE_INVESTMENT = "investment"
	COST_TYPE_CONSULTING = "consulting"
	COST_TYPE_OTHERS     = "others"
)

type Discount struct {
	Name       string  `json:"name" yaml:"name"`
	Percentage float64 `json:"percentage" yaml:"percentage"`
	Absolute   float64 `json:"absolute" yaml:"absolute"`
}

type Cost struct {
	Name        string   `json:"name" yaml:"name"`
	Description string   `json:"description" yaml:"description"`
	Type        string   `json:"type" yaml:"type"`
	Amount      float64  `json:"amount" yaml:"amount"`
	Currency    string   `json:"currency" yaml:"currency"`
	Tags        []string `json:"tags" yaml:"tags"`
	External    bool     `json:"external" yaml:"external"`
}

type Allocation struct {
	Cost       string  `json:"cost" yaml:"cost"`
	Occurrence float64 `json:"occurrence" yaml:"occurrence"`
	Cycles     []int   `json:"cycles" yaml:"cycles"`
	Discount   string  `json:"discount" yaml:"discount"`
}

type Cycle struct {
	Name  string `json:"name" yaml:"name"`
	Index int    `json:"index" yaml:"index"`
}

type SubReport struct {
	Labor         map[int]float64
	Investment    map[int]float64
	Consulting    map[int]float64
	Others        map[int]float64
	TotLabor      float64
	TotInvestment float64
	TotConsulting float64
	TotOthers     float64
}

type Report struct {
	External SubReport
	Internal SubReport
}

type CBA struct {
	Cycles      []Cycle      `json:"cycles" yaml:"cycles"`
	Discounts   []Discount   `json:"discounts" yaml:"discounts"`
	Costs       []Cost       `json:"costs" yaml:"costs"`
	Allocations []Allocation `json:"allocations" yaml:"allocations"`
	Currency    string       `json:"currency" yaml:"currency"`
	Report      Report
}
