package formula

import (
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

func calculateExponentOnNumber(z complex128, n int) complex128 {
	return cmplx.Pow(z, complex(float64(n), 0))
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
	PlusNNoConjugate = "PlusNNoConjugate"
	PlusMNoConjugate = "PlusMNoConjugate"
)

type CalculationInstruction struct {
	firstExponentPower string
	secondExponentPower string
	useComplexConjugate bool
}

var instructionsByCoefficientRelationship = map[CoefficientRelationship]CalculationInstruction{
	PlusNPlusM: {
		firstExponentPower: "n",
		secondExponentPower: "m",
		useComplexConjugate: true,
	},
	PlusMPlusN: {
		firstExponentPower: "m",
		secondExponentPower: "n",
		useComplexConjugate: true,
	},
	PlusNNoConjugate: {
		firstExponentPower: "n",
		secondExponentPower: "",
		useComplexConjugate: false,
	},
	PlusMNoConjugate: {
		firstExponentPower: "m",
		secondExponentPower: "",
		useComplexConjugate: false,
	},
}

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

			instructions := instructionsByCoefficientRelationship[relationship]
			firstPower := setCoefficientBasedOnInstruction(coeffs.PowerN, coeffs.PowerM, instructions.firstExponentPower)
			secondPower := setCoefficientBasedOnInstruction(coeffs.PowerN, coeffs.PowerM, instructions.secondExponentPower)

			if instructions.useComplexConjugate {
				sum += CalculateExponentPairOnNumberAndConjugate(z, firstPower, secondPower) * coeffs.Scale
			} else {
				sum += calculateExponentOnNumber(z, firstPower) * coeffs.Scale
			}
		}
	}
	return sum
}

func setCoefficientBasedOnInstruction(powerN, powerM int, exponentPowerSetting string) int {
	switch exponentPowerSetting {
	case "n":
		return powerN
	case "m":
		return powerM
	default:
		return 0
	}
}
