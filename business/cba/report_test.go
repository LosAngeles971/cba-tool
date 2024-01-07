package cba

import(
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcReport(t *testing.T) {
	cba := NewCBA()
	cba.Load(test_project)
	r := cba.CalcReport()
	require.Greater(t, r.Total, 0.0)
}