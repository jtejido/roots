package roots

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/roots/err"
	"math"
)

type RiddersMethod struct{}

func (rm RiddersMethod) Solve(f Function, a, b, tol float64, result *float64) err.RootsError {

	if err := rm.validate(a, b, tol); err != nil {
		return err
	}

	x1 := a
	x2 := b
	fx1 := f.Evaluate(x1)
	fx2 := f.Evaluate(x2)
	halfEps := tol * 0.5

	if fx1*fx2 >= 0 {
		*result = math.NaN()
		return err.Error("The given interval does not appear to bracket the root", err.EDOM)
	}

	dif := 1.
	maxIterations := ridder_iter
	for math.Abs(x1-x2) > tol && maxIterations > 0 {
		x3 := (x1 + x2) * 0.5

		fx3 := f.Evaluate(x3)

		x4 := x3 + (x3-x1)*float64(gsl.Sign(fx1-fx2))*fx3/math.Sqrt(fx3*fx3-fx1*fx2)

		fx4 := f.Evaluate(x4)
		if fx3*fx4 < 0 {
			x1 = x3
			fx1 = fx3
			x2 = x4
			fx2 = fx4
		} else if fx1*fx4 < 0 {
			dif = math.Abs(x4 - x2)
			if dif <= halfEps {
				*result = x4
				return nil
			}
			x2 = x4
			fx2 = fx4
		} else {
			dif = math.Abs(x4 - x1)
			if dif <= halfEps {
				*result = x4
				return nil
			}
			x1 = x4
			fx1 = fx4
		}
		maxIterations--
	}

	*result = x2
	return nil
}

func (rm RiddersMethod) validate(p0, p1, tol float64) err.RootsError {
	if err := checkTolerance(tol); err != nil {
		return err
	}

	if err := checkInterval(p0, p1); err != nil {
		return err
	}

	return nil
}
