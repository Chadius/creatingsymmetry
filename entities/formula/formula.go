package formula

import (
	"fmt"
	"log"
	"math/cmplx"
)

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

// CoefficientRelationship TODO
type CoefficientRelationship string

const (
	PlusNPlusM CoefficientRelationship = "PlusNPlusM"
	PlusMPlusN                         = "PlusMPlusN"
)


// RecipeFormula TODO
type RecipeFormula struct {
	Coefficients  []*CoefficientPairs
	Relationships []CoefficientRelationship
}

// Calculate takes the complex number z and applies the formula to it,
//     returning another complex number.
func (f RecipeFormula) Calculate(z complex128) complex128 {
	sum := complex(0,0)
	for _, coeffs := range f.Coefficients {
		for _, relationship := range f.Relationships {
			firstPower, secondPower, err := changeCoefficientsBasedOnRelationship(coeffs.PowerN, coeffs.PowerM, relationship)
			if err != nil {
				log.Fatal(err)
			}
			sum += CalculateExponentPairOnNumberAndConjugate(z, firstPower, secondPower) * coeffs.Scale
		}
	}
	return sum
}

// changeCoefficientsBasedOnRelationship uses the relationship to determine how to calculate the given
//   powers of n and m.
func changeCoefficientsBasedOnRelationship(powerN, powerM int, relationship CoefficientRelationship) (int, int, error) {
	switch relationship {
	case PlusNPlusM:
		return powerN, powerM, nil
	case PlusMPlusN:
		return powerM, powerN, nil
	}
	return 0, 0, fmt.Errorf("unknown relationship %s", relationship)
}