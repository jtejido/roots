package roots

import (
	"math"
	"strconv"
	"testing"
)

type fpif struct{}

func (f fpif) Evaluate(x float64) float64 {
	return (math.Pow(x, 4) + 8*math.Pow(x, 3) - 13*math.Pow(x, 2) + 96) / 2
}

func TestFixedPointIterationMethodSolve(t *testing.T) {

	f := fpif{}
	tol := 0.00001
	cases := []struct {
		a, b, p, expected float64
	}{
		{
			a:        0,
			b:        2,
			p:        0,
			expected: 1,
		},
		{
			a:        2,
			b:        0,
			p:        0,
			expected: 1,
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var fpi FixedPointIteration
			var result float64
			err := fpi.Solve(&f, c.a, c.b, c.p, tol, &result)

			if err == nil {
				if math.Abs(result-c.expected) > tol {
					t.Errorf("Mismatch. Case %d, want: %v, got: %v", i, c.expected, result)
				}
			} else {
				t.Errorf("Mismatch. got error: %s", err.Error())
			}
		})
	}

}
