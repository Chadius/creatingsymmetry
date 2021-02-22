package formula

import "math/cmplx"

// CalculateExponentPairOnNumberAndConjugate calculates (z^n * ~z^m)
//   where z is a complex number, ~z is the complex conjugate
//   and n and m are integers.
func CalculateExponentPairOnNumberAndConjugate(z complex128, n, m int) complex128 {
	complexConjugate := complex(real(z), -1 * imag(z))

	zRaisedToN := cmplx.Pow(z, complex(float64(n), 0))
	complexConjugateRaisedToM := cmplx.Pow(complexConjugate, complex(float64(m), 0))
	return zRaisedToN * complexConjugateRaisedToM
}