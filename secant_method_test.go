package roots

import (
	"math"
	"strconv"
	"testing"
)

func TestSecantMethodSolve(t *testing.T) {

	f := bf{}
	tol := 0.00001
	cases := []struct {
		p0, p1, expected float64
	}{
		{
			p0:       -5,
			p1:       -2,
			expected: -4,
		},
		{
			p0:       -10,
			p1:       -7,
			expected: -8,
		},
		{
			p0:       2,
			p1:       5,
			expected: 3,
		},
		{
			p0:       -1,
			p1:       2,
			expected: 1,
		},
		{
			p0:       2,
			p1:       -1,
			expected: 1,
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var sm SecantMethod
			var result float64
			err := sm.Solve(f, c.p0, c.p1, tol, &result)

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
