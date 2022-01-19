package roots

import (
	"math"
	"strconv"
	"testing"
)

type bf struct{}

func (f bf) Evaluate(x float64) float64 {
	return math.Pow(x, 4) + 8*math.Pow(x, 3) - 13*math.Pow(x, 2) - 92*x + 96
}

func TestBisectionMethodSolve(t *testing.T) {

	f := bf{}
	tol := 0.00001

	cases := []struct {
		a        float64
		b        float64
		expected float64
	}{
		{
			a:        -7,
			b:        0,
			expected: -4,
		},
		{
			a:        2,
			b:        5,
			expected: 3,
		},
		{
			a:        -10,
			b:        -5,
			expected: -8,
		},
		{
			a:        0,
			b:        2,
			expected: 1,
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var bm BisectionMethod
			var result float64
			err := bm.Solve(&f, c.a, c.b, tol, &result)

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
