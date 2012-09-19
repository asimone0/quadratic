package quadratic

import (
	"math/cmplx"
)

//Solve returns the roots of a quadratic equation
//with coefficients a,b,c
func Solve(a, b, c complex128) (xpos, xneg complex128) {
	negB := -b
	twoA := 2 * a
	bSquared := b * b
	fourAC := 4 * a * c
	discrim := bSquared - fourAC
	sq := cmplx.Sqrt(complex128(discrim))
	xpos = (negB + sq) / twoA
	xneg = (negB - sq) / twoA
	return
}
