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
//	 MaybeFlipScale will multiply the scale by -1 if N + M is odd.
const (
	PlusNPlusM CoefficientRelationship = "+N+M"
	PlusMPlusN                         = "+M+N"
	MinusNMinusM                       = "-N-M"
	MinusMMinusN                       = "-M-N"
	PlusMPlusNMaybeFlipScale           = "+M+NF"
	MinusMMinusNMaybeFlipScale         = "-M-NF"
)

// SetCoefficientsBasedOnRelationship will rearrange powerN and powerM according to their relationship.
func SetCoefficientsBasedOnRelationship(powerN, powerM int, scale complex128, relationship CoefficientRelationship) (int, int, complex128) {
	var power1, power2 int
	switch relationship {
	case PlusNPlusM:
		power1 = powerN
		power2 = powerM
	case PlusMPlusN, PlusMPlusNMaybeFlipScale:
		power1 = powerM
		power2 = powerN
	case MinusMMinusN, MinusMMinusNMaybeFlipScale:
		power1 = -1 * powerM
		power2 = -1 * powerN
	case MinusNMinusM:
		power1 = -1 * powerN
		power2 = -1 * powerM
	}

	sumOfPowersIsOdd := (powerN + powerM) % 2 == 1
	relationshipMayFlipScale := relationship == PlusMPlusNMaybeFlipScale || relationship == MinusMMinusNMaybeFlipScale
	if sumOfPowersIsOdd && relationshipMayFlipScale {
		scale *= -1
	}

	return power1, power2, scale
}

// ZExponentialFormulaElement describes a formula of the form Scale * z^PowerN * zConjugate^PowerM.
type ZExponentialFormulaElement struct {
	Scale                  complex128
	PowerN                 int
	PowerM                 int
	// IgnoreComplexConjugate will make sure zConjugate is not used in this calculation
	//    (effectively setting it to 1 + 0i)
	IgnoreComplexConjugate bool
	// CoefficientPairs will create similar terms to add to this one when calculating.
	//    This is useful when trying to force symmetry by adding another term with swapped
	//    PowerN & PowerM, or multiplying by -1.
	CoefficientPairs LockedCoefficientPair
}

// Calculate returns the result of using the formula on the given complex number.
func (element ZExponentialFormulaElement) Calculate(z complex128) complex128 {
	sum := CalculateExponentElement(z, element.PowerN, element.PowerM, element.Scale, element.IgnoreComplexConjugate)

	for _, relationship := range element.CoefficientPairs.OtherCoefficientRelationships {
		power1, power2, scale := SetCoefficientsBasedOnRelationship(element.PowerN, element.PowerM, element.Scale, relationship)
		relationshipScale := scale * complex(element.CoefficientPairs.Multiplier, 0)
		sum += CalculateExponentElement(z, power1, power2, relationshipScale, element.IgnoreComplexConjugate)
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
	// CoefficientPairs will create similar terms to add to this one when calculating.
	//    This is useful when trying to force symmetry by adding another term with swapped
	//    PowerN & PowerM, or multiplying by -1.
	CoefficientPairs LockedCoefficientPair
}

// Calculate returns the result of using the formula on the given complex number.
func (element EulerFormulaElement) Calculate(z complex128) complex128 {
	sum := CalculateEulerElement(z, element.PowerN, element.PowerM, element.Scale, element.IgnoreComplexConjugate)

	for _, relationship := range element.CoefficientPairs.OtherCoefficientRelationships {
		power1, power2, scale := SetCoefficientsBasedOnRelationship(element.PowerN, element.PowerM, element.Scale, relationship)
		relationshipScale := scale * complex(element.CoefficientPairs.Multiplier, 0)
		sum += CalculateEulerElement(z, power1, power2, relationshipScale, element.IgnoreComplexConjugate)
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

// FriezeSymmetry notes the kinds of symmetries the formula contains.
type FriezeSymmetry struct {
	P111 bool
	P11m bool
	P211 bool
	P1m1 bool
	P11g bool
	P2mm bool
	P2mg bool
}

//AnalyzeForSymmetry scans the formula and returns a list of symmetries.
func (formula FriezeFormula) AnalyzeForSymmetry() *FriezeSymmetry {
	symmetriesFound := &FriezeSymmetry{
		P111: true,
		P11m: true,
		P211: true,
		P1m1: true,
		P11g: true,
		P2mm: true,
		P2mg: true,
	}
	for _, element := range formula.Elements {
		if element.IgnoreComplexConjugate {
			symmetriesFound.P211 = false
			symmetriesFound.P1m1 = false
			symmetriesFound.P11g = false
			symmetriesFound.P11m = false
			symmetriesFound.P2mm = false
			symmetriesFound.P2mg = false
		}

		powerSumIsEven := (element.PowerN + element.PowerM) % 2 == 0

		containsMinusNMinusM := coefficientPairsIncludes(element.CoefficientPairs.OtherCoefficientRelationships, MinusNMinusM)
		containsMinusMMinusN := coefficientPairsIncludes(element.CoefficientPairs.OtherCoefficientRelationships, MinusMMinusN)
		containsPlusMPlusN := coefficientPairsIncludes(element.CoefficientPairs.OtherCoefficientRelationships, PlusMPlusN)

		containsMinusMMinusNAndPowerSumIsOdd := coefficientPairsIncludes(element.CoefficientPairs.OtherCoefficientRelationships, MinusMMinusNMaybeFlipScale ) && !powerSumIsEven
		containsPlusMPlusNAndPowerSumIsOdd := coefficientPairsIncludes(element.CoefficientPairs.OtherCoefficientRelationships, PlusMPlusNMaybeFlipScale) && !powerSumIsEven

		containsMinusMMinusNAndPowerSumIsEven := coefficientPairsIncludes(element.CoefficientPairs.OtherCoefficientRelationships, MinusMMinusNMaybeFlipScale ) && powerSumIsEven
		containsPlusMPlusNAndPowerSumIsEven := coefficientPairsIncludes(element.CoefficientPairs.OtherCoefficientRelationships, PlusMPlusNMaybeFlipScale) && powerSumIsEven

		if !containsMinusNMinusM {
			symmetriesFound.P211 = false
		}
		if !containsPlusMPlusN {
			symmetriesFound.P1m1 = false
		}
		if !containsMinusMMinusNAndPowerSumIsOdd {
			symmetriesFound.P11g = false
		}
		if !(containsMinusMMinusN || containsMinusMMinusNAndPowerSumIsEven) {
			symmetriesFound.P11m = false
		}
		if !(
			containsMinusNMinusM &&
				(containsPlusMPlusN || containsPlusMPlusNAndPowerSumIsEven) &&
				(containsMinusMMinusN || containsMinusMMinusNAndPowerSumIsEven)) {
			symmetriesFound.P2mm = false
		}
		if !(containsMinusNMinusM && containsPlusMPlusNAndPowerSumIsOdd && containsMinusMMinusNAndPowerSumIsOdd) {
			symmetriesFound.P2mg = false
		}
	}

	return symmetriesFound
}

func coefficientPairsIncludes (relationships []CoefficientRelationship, relationshipToFind CoefficientRelationship) bool {
	for _, relationship := range relationships {
		if relationship == relationshipToFind {
			return true
		}
	}
	return false
}