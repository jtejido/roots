package roots

import (
	"github.com/jtejido/roots/err"
)

var (
	toleranceError        = "Tolerance must be greater than zero."
	sameIntervalError     = "Start point and end point of interval cannot be the same."
	sameSignError         = "Input function has the same sign at the start and end of the interval. Choose start and end points such that the function evaluated at those points has a different sign (one positive, one negative)."
	initialGuessError     = "Initial guess p must be in [a, b]."
	errRootIsNotBracketed = "No root in the given interval"
	iterError             = "Too many iterations"
	badFuncErr            = "Bad func error"
)

func OutOfBoundsErrorMsg(msg string) err.RootsError {
	return err.Error(msg, err.ERANGE)
}

func InvalidInputMsg(msg string) err.RootsError {
	return err.Error(msg, err.EINVAL)
}

func DomainError() err.RootsError {
	return err.Error(errRootIsNotBracketed, err.EDOM)
}

func MaxIterError() err.RootsError {
	return err.Error(iterError, err.EMAXITER)
}

func BadFuncError() err.RootsError {
	return err.Error(badFuncErr, err.EBADFUNC)
}
