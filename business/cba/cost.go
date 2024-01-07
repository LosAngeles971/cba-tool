package cba

const (
	COST_TYPE_LABOR      = "labor"
	COST_TYPE_INVESTMENT = "investment"
	COST_TYPE_CONSULTING = "consulting"
	COST_TYPE_OTHERS     = "others"

	COST_METRIC_YEARLY = "yearly"
	COST_METRIC_DAILY  = "daily"
	COST_METRIC_ONCE   = "once"
)

type Cost struct {
	Name     string   `json:"name" yaml:"name"`
	Metric   string   `json:"metric" yaml:"metric"`
	Type     string   `json:"type" yaml:"type"`
	Amount   float64  `json:"amount" yaml:"amount"`
	Currency string   `json:"currency" yaml:"currency"`
	Tags     []string `json:"tags" yaml:"tags"`
	External bool     `json:"external" yaml:"external"`
}

func (cba *CBA) NewCost() *Cost {
	return  &Cost{
		Name: "",
		Metric: COST_METRIC_ONCE,
		Type: COST_TYPE_INVESTMENT,
		Amount: 0.0,
		Currency: CURRENCY_EURO,
		Tags: []string{},
		External: true,
	}
}

func (cba *CBA) FindCostByName(name string) *Cost {
	for _, c := range cba.Costs {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func (cba *CBA) DeleteCostByName(name string) {
	buffer := []*Cost{}
	for _, c := range cba.Costs {
		if c.Name != name {
			buffer = append(buffer, c)
		}
	}
	cba.Costs = buffer
}

func (cba *CBA) ListCosts() []string {
	cc := []string{}
	for _, c := range cba.Costs {
		cc = append(cc, c.Name)
	}
	return cc
}