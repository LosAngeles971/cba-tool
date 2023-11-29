package business

import(
	"testing"
	_ "embed"

	"github.com/stretchr/testify/require"
)

//go:embed example.yaml
var example []byte

func TestCalc(t *testing.T) {
	cba := NewCBA(example)
	cba.Calc()
	require.Equal(t, "EUR", cba.Currency)
}