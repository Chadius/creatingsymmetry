package formula

import (
	"math/cmplx"
)

// CoefficientRelationship relates how a pair of coordinates should be applied.
type CoefficientRelationship string

// CoefficientRelationship s determine the order and sign of powers n and m.
//   Plus means *1, Minus means *-1
//   If N appears first the powers then power N is applied to the number and power M to the complex conjugate.
//   If M appears first the powers then power M is applied to the number and power N to the complex conjugate.
const (
	// PlusNPlusM Apply N to the first and M to the second complex number.
	PlusNPlusM CoefficientRelationship = "+N+M"
	// PlusMPlusN Apply M to the first and N to the second complex number.
	PlusMPlusN                         = "+M+N"
	MinusNMinusM = "-N-M"
	MinusMMinusN = "-M-N"
)

// SetCoefficientsBasedOnRelationship will rearrange powerN and powerM according to their relationship.
func SetCoefficientsBasedOnRelationship(powerN, powerM int, relationship CoefficientRelationship) (int, int) {
	var power1, power2 int
	switch relationship {
	case PlusNPlusM:
		power1 = powerN
		power2 = powerM
	case PlusMPlusN:
		power1 = powerM
		power2 = powerN
	case MinusMMinusN:
		power1 = -1 * powerM
		power2 = -1 * powerN
	case MinusNMinusM:
		power1 = -1 * powerN
		power2 = -1 * powerM
	}
	return power1, power2
}

// ZExponentialFormulaElement describes a formula of the form Scale * z^PowerN * zConjugate^PowerM.
type ZExponentialFormulaElement struct {
	Scale                  complex128
	PowerN                 int
	PowerM                 int
	// IgnoreComplexConjugate will make sure zConjugate is not used in this calculation
	//    (effectively setting it to 1 + 0i)
	IgnoreComplexConjugate bool
	// LockedCoefficientPairs will create similar terms to add to this one when calculating.
	//    This is useful when trying to force symmetry by adding another term with swapped
	//    PowerN & PowerM, or multiplying by -1.
	LockedCoefficientPairs []*LockedCoefficientPair
}

// Calculate returns the result of using the formula on the given complex number.
func (element ZExponentialFormulaElement) Calculate(z complex128) complex128 {
	sum := CalculateExponentElement(z, element.PowerN, element.PowerM, element.Scale, element.IgnoreComplexConjugate)

	for _, pair := range element.LockedCoefficientPairs {
		for _, relationship := range pair.OtherCoefficientRelationships {
			power1, power2 := SetCoefficientsBasedOnRelationship(element.PowerN, element.PowerM, relationship)
			relationshipScale := element.Scale * complex(pair.Multiplier, 0)
			sum += CalculateExponentElement(z, power1, power2, relationshipScale, element.IgnoreComplexConjugate)
		}
	}
	return sum
}

// LockedCoefficientPair describes how to create a new Element based on the current one.
type LockedCoefficientPair struct {
	Multiplier                    float64
	OtherCoefficientRelationships []CoefficientRelationship
}

// CalculateExponentElement calculates (z^power * zConj^conjugatePower)
//   where z is a complex number, zConj is the complex conjugate
//   and power and conjugatePower are integers.
func CalculateExponentElement(z complex128, power1, power2 int, scale complex128, ignoreComplexConjugate bool) complex128 {
	zRaisedToN := cmplx.Pow(z, complex(float64(power1), 0))
	if ignoreComplexConjugate {
		return zRaisedToN * scale
	}

	complexConjugate := complex(real(z), -1 * imag(z))
	complexConjugateRaisedToM := cmplx.Pow(complexConjugate, complex(float64(power2), 0))
	return zRaisedToN * complexConjugateRaisedToM * scale
}


// RosetteFormula uses a collection of z^m terms to calculate results.
//    This transforms the input into a circular pattern rotating around the
//    origin.
type RosetteFormula struct {
	Elements []*ZExponentialFormulaElement
}

// Calculate applies the Rosette formula to the complex number z.
func (r RosetteFormula) Calculate(z complex128) complex128 {
	sum := complex(0,0)
	for _, term := range r.Elements {
		sum += term.Calculate(z)
	}

	return sum
}

// EulerFormulaElement calculates e^(i*n*z) * e^(-i*m*zConj)
type EulerFormulaElement struct {
	Scale                  complex128
	PowerN                 int
	PowerM                 int
	// IgnoreComplexConjugate will make sure zConjugate is not used in this calculation
	//    (effectively setting it to 1 + 0i)
	IgnoreComplexConjugate bool
	// LockedCoefficientPairs will create similar terms to add to this one when calculating.
	//    This is useful when trying to force symmetry by adding another term with swapped
	//    PowerN & PowerM, or multiplying by -1.
	LockedCoefficientPairs []*LockedCoefficientPair
}

// Calculate returns the result of using the formula on the given complex number.
func (element EulerFormulaElement) Calculate(z complex128) complex128 {
	sum := CalculateEulerElement(z, element.PowerN, element.PowerM, element.Scale, element.IgnoreComplexConjugate)
	for _, pair := range element.LockedCoefficientPairs {
		for _, relationship := range pair.OtherCoefficientRelationships {
			power1, power2 := SetCoefficientsBasedOnRelationship(element.PowerN, element.PowerM, relationship)
			relationshipScale := element.Scale * complex(pair.Multiplier, 0)
			sum += CalculateEulerElement(z, power1, power2, relationshipScale, element.IgnoreComplexConjugate)
		}
	}

	return sum
}

// CalculateEulerElement calculates e^(i*n*z) * e^(-i*m*zConj)
func CalculateEulerElement(z complex128, power1, power2 int, scale complex128, ignoreComplexConjugate bool) complex128 {
	eRaisedToTheNZi := cmplx.Exp(complex(0,1) * z * complex(float64(power1), 0))
	if ignoreComplexConjugate {
		return eRaisedToTheNZi * scale
	}

	complexConjugate := complex(real(z), -1 * imag(z))
	eRaisedToTheNegativeMZConji := cmplx.Exp(complexConjugate * complex(0, -1 * float64(power2)))
	return eRaisedToTheNZi * eRaisedToTheNegativeMZConji * scale
}

// FriezeFormula is used to generate frieze patterns.
type FriezeFormula struct {
	Elements []*EulerFormulaElement
}

// Calculate applies the Frieze formula to the complex number z.
func (formula FriezeFormula) Calculate(z complex128) complex128 {
	sum := complex(0,0)
	for _, term := range formula.Elements {
		sum += term.Calculate(z)
	}

	return sum
}