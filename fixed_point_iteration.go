package roots

import (
	"github.com/jtejido/roots/err"
	"math"
)

type FixedPointIteration struct{}

func (fsi FixedPointIteration) Solve(f Function, a, b, init, tol float64, result *float64) err.RootsError {

	if err := fsi.validate(&a, &b, init, tol); err != nil {
		return err
	}

	p := init

	for {
		gp := f.Evaluate(p)
		dif := math.Abs(gp - p)
		p = gp

		if !(dif > tol) {
			break
		}
	}

	*result = p
	return nil
}

func (fsi FixedPointIteration) validate(a, b *float64, p, tol float64) err.RootsError {
	if err := checkTolerance(tol); err != nil {
		return err
	}

	if err := checkInterval(*a, *b); err != nil {
		return err
	}

	if *a > *b {
		*a, *b = *b, *a
	}

	if p < *a || p > *b {
		return OutOfBoundsErrorMsg(initialGuessError)
	}

	return nil
}
