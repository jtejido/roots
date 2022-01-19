package roots

import (
	"github.com/jtejido/roots/err"
	"math"
)

type SecantMethod struct{}

func (sm SecantMethod) Solve(f Function, p0, p1, tol float64, result *float64) err.RootsError {

	if err := sm.validate(p0, p1, tol); err != nil {
		return err
	}

	var p float64
	for {
		q0 := f.Evaluate(p0)
		q1 := f.Evaluate(p1)
		slope := (q1 - q0) / (p1 - p0)

		p = p1 - (q1 / slope)

		dif := math.Abs(p - p1)
		p0 = p1
		p1 = p

		if !(dif > tol) {
			break
		}
	}

	*result = p
	return nil
}

func (sm SecantMethod) validate(p0, p1, tol float64) err.RootsError {
	if err := checkTolerance(tol); err != nil {
		return err
	}

	if err := checkInterval(p0, p1); err != nil {
		return err
	}

	return nil
}
