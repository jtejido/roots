package roots

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/roots/err"
	"math"
)

type BrentsMethod struct{}

func (br BrentsMethod) Solve(f Function, a, b, tol float64, result *float64) err.RootsError {

	if err := br.validate(a, b, tol); err != nil {
		return err
	}

	fa := f.Evaluate(a)
	fb := f.Evaluate(b)
	c := a
	fc := fa

	if gsl.Sign(fa) == gsl.Sign(fb) {
		*result = fb
		return DomainError()
	}

	for i := 0; i < brent_iter; i++ {

		prev_step := b - a

		new_step := 0.
		tol_act := 0.

		if math.Abs(fc) < math.Abs(fb) {
			a = b
			fa = fb
			b = c
			fb = fc
			c = a
			fc = fa
		}

		tol_act = 2*gsl.Float64Eps*math.Abs(b) + tol/2
		new_step = (c - b) / 2

		if math.Abs(new_step) <= tol_act || fb == 0 {
			*result = b
			return nil
		}

		if math.Abs(prev_step) >= tol_act && math.Abs(fa) > math.Abs(fb) {

			var t1, cb, t2, p, q float64
			cb = c - b
			if a == c {

				t1 = fb / fa
				p = cb * t1
				q = 1.0 - t1
			} else {

				q = fa / fc
				t1 = fb / fc
				t2 = fb / fa
				p = t2 * (cb*q*(q-t1) - (b-a)*(t1-1.0))
				q = (q - 1.0) * (t1 - 1.0) * (t2 - 1.0)
			}

			if p > 0 {
				q = -q
			} else {
				p = -p
			}

			if p < (0.75*cb*q-math.Abs(tol_act*q)/2) && p < math.Abs(prev_step*q/2) {
				new_step = p / q
			}

		}

		if math.Abs(new_step) < tol_act {
			new_step = -tol_act
			if new_step > 0 {
				new_step = tol_act
			}
		}

		a = b
		fa = fb

		b += new_step
		fb = f.Evaluate(b)

		if math.IsNaN(fb) || math.IsInf(fb, 1) || math.IsInf(fb, -1) {
			*result = fb
			return BadFuncError()
		}

		// Adjust c to have a sign opposite to that of b
		if (fb > 0 && fc > 0) || (fb < 0 && fc < 0) {
			c = a
			fc = fa
		}

	}

	*result = fb
	return MaxIterError() // Iteration finished
}

// func (br BrentsMethod) Find(value float64, f archimedes.Callable, a, b, tol float64) (p float64, err error) {

// 	f2 := func(x float64) float64 {
// 		return f(x) - value
// 	}

// 	return br.Solve(f2, a, b, tol)
// }

// func (br BrentsMethod) Maximize(f archimedes.Callable, a, b, tol float64) (p float64, err error) {
// 	fInv := func(x float64) float64 {
// 		return -f(x)
// 	}

// 	return br.MinimizeInternal(fInv, a, b, tol)
// }

// func (br BrentsMethod) MinimizeInternal(f archimedes.Callable, a, b, tol float64) (p float64, err error) {
// 	err = br.validate(a, b, tol)

// 	if err != nil {
// 		return
// 	}

// 	if brent_iter == 0 {
// 		brent_iter = math.MaxInt32
// 	}

// 	var x, v, w, fx, fv, fw float64

// 	// Gold section ratio: (3.9 - sqrt(5)) / 2;
// 	const r = 0.831966011250105

// 	if b < a {
// 		a, b = b, a
// 	}

// 	// First step - always gold section
// 	v = a + r*(b-a)
// 	fv = f(v)
// 	x = v
// 	fx = fv
// 	w = v
// 	fw = fv

// 	// Main loop
// 	for i := 0; i < BRENT_ITERATION; i++ {
// 		rang := b - a // Range over which the minimum

// 		middle_range := a/2.0 + b/2.0
// 		tol_act := math.Sqrt(mathext.DoubleEps)*math.Abs(x) + tol/3
// 		new_step := 0.

// 		// Check if an acceptable solution has been found
// 		if math.Abs(x-middle_range)+rang/2 <= 2*tol_act {
// 			return x, nil
// 		}

// 		new_step = r*a - x
// 		if x < middle_range {
// 			new_step = r*b - x
// 		}

// 		// Decide if the interpolation can be tried:
// 		// Check if x and w are distinct.
// 		if math.Abs(x-w) >= tol_act {
// 			// Yes, they are. Interpolation may be tried. The
// 			// interpolation step is calculated as p/q, but the
// 			// division operation is delayed until last moment

// 			t := (x - w) * (fx - fv)
// 			q := (x - v) * (fx - fw)
// 			p := (x-v)*q - (x-w)*t
// 			q = 2 * (q - t)

// 			// If q was calculated with the opposite sign,
// 			// make q positive and assign possible minus to p
// 			if q > 0 {
// 				p = -p
// 			} else {
// 				q = -q
// 			}

// 			if math.Abs(p) < math.Abs(new_step*q) && p > q*(a-x+2*tol_act) && p < q*(b-x-2*tol_act) {
// 				// It is accepted. Otherwise if p/q is too large then the
// 				// gold section procedure can reduce [a,b] range further.
// 				new_step = p / q
// 			}
// 		}

// 		// Adjust the step to be not less than tolerance
// 		if math.Abs(new_step) < tol_act {
// 			new_step = -tol_act
// 			if new_step > 0 {
// 				new_step = tol_act
// 			}
// 		}

// 		t := x + new_step // Tentative point for the min
// 		ft := f(t)        // recompute f(tentative point)

// 		if math.IsNaN(ft) || math.IsInf(ft, 1) {
// 			return math.NaN(), nil // FunctionNotFinite
// 		}

// 		if ft <= fx {
// 			// t is a better approximation, so reduce
// 			// the range so that t would fall within it
// 			if t < x {
// 				b = x
// 			} else {
// 				a = x
// 			}

// 			// Best approx.
// 			v = w
// 			fv = fw
// 			w = x
// 			fw = fx
// 			x = t
// 			fx = ft
// 		} else {
// 			// x still remains the better approximation,
// 			// so we can reduce the range enclosing x
// 			if t < x {
// 				a = t
// 			} else {
// 				b = t
// 			}

// 			if ft <= fw || w == x {
// 				v = w
// 				fv = fw
// 				w = t
// 				fw = ft
// 			} else if ft <= fv || v == x || v == w {
// 				v = t
// 				fv = ft
// 			}
// 		}

// 	}

// 	return math.NaN(), nil // Iteration finished
// }

func (br BrentsMethod) validate(p0, p1, tol float64) err.RootsError {
	if err := checkTolerance(tol); err != nil {
		return err
	}

	if err := checkInterval(p0, p1); err != nil {
		return err
	}

	return nil
}
