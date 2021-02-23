package formula

import "math/cmplx"

// CalculateExponentPairOnNumberAndConjugate calculates (z^n * ~z^m)
//   where z is a complex number, ~z is the complex conjugate
//   and n and m are integers.
//   This computation is used in almost every symmetry calculation.
func CalculateExponentPairOnNumberAndConjugate(z complex128, n, m int) complex128 {
	complexConjugate := complex(real(z), -1 * imag(z))

	zRaisedToN := cmplx.Pow(z, complex(float64(n), 0))
	complexConjugateRaisedToM := cmplx.Pow(complexConjugate, complex(float64(m), 0))
	return zRaisedToN * complexConjugateRaisedToM
}

// CoefficientPairs holds two coefficients n and m and a scale.
type CoefficientPairs struct {
	Scale complex128
	PowerN int
	PowerM int
}

// Calculate calculates (z^n * ~z^m)
//   where z is a complex number, ~z is the complex conjugate
//   and n and m are integers.
func (p CoefficientPairs) Calculate(z complex128) complex128 {
	return CalculateExponentPairOnNumberAndConjugate(z, p.PowerN, p.PowerM) * p.Scale
}

// SymmetryFormula is a mathematical formula that works on
//   complex numbers.
type SymmetryFormula struct {
	PairedCoefficients []*CoefficientPairs
}

// Calculate takes the complex number z and applies the formula to it,
//     returning another complex number.
func (f SymmetryFormula) Calculate(z complex128) complex128 {
	sum := complex(0,0)
	for _, coeffs := range f.PairedCoefficients {
		sum += coeffs.Calculate(z)
	}

	return sum
}