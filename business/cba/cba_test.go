package cba

import(
	"testing"
	_ "embed"

	"github.com/stretchr/testify/require"
)

//go:embed example.yaml
var test_project []byte

func TestCBALoad(t *testing.T) {
	cba := NewCBA()
	cba.Load(test_project)
	require.Equal(t, "EUR", cba.Currency)
}