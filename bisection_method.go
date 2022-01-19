package roots

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/roots/err"
	"math"
)

type BisectionMethod struct{}

func (bm BisectionMethod) Solve(f Function, a, b, tol float64, result *float64) err.RootsError {

	if err := bm.validate(f, a, b, tol); err != nil {
		return err
	}

	var dif float64

	var p float64
	for {
		fa := f.Evaluate(a)

		p = float64((a + b) / 2)

		fp := f.Evaluate(p)
		dif = math.Abs(fp)
		if gsl.Sign(fp) != gsl.Sign(fa) {
			b = p
		} else {
			a = p
		}

		if !(dif > tol) {
			break
		}
	}

	*result = p

	return nil
}

func (bm BisectionMethod) validate(f Function, a, b, tol float64) err.RootsError {

	if err := checkTolerance(tol); err != nil {
		return err
	}

	if err := checkInterval(a, b); err != nil {
		return err
	}

	fa := f.Evaluate(a)
	fb := f.Evaluate(b)

	if gsl.Sign(fa) == gsl.Sign(fb) {
		return OutOfBoundsErrorMsg(sameSignError)
	}

	return nil
}
