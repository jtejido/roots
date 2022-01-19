package roots

import (
	"github.com/jtejido/roots/err"
)

type Function interface {
	Evaluate(x float64) float64
}

var (
	nm_iter     = 100
	brent_iter  = 500
	ridder_iter = 1000
)

func checkTolerance(tol float64) err.RootsError {
	if tol < 0 {
		return InvalidInputMsg(toleranceError)
	}

	return nil
}

func checkInterval(a, b float64) err.RootsError {
	if a == b {
		return InvalidInputMsg(sameIntervalError)
	}

	return nil
}
