package formula

import (
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"math/cmplx"
)

// Frieze formulas transform points into a horizontal repeating strip, like the frieze patterns on ceilings and columns.
type Frieze struct {
	formulaLevelTerms []Term
}

// NewFriezeFormula returns a new formula
func NewFriezeFormula(formulaLevelTerms []Term) *Frieze {
	return &Frieze{
		formulaLevelTerms: formulaLevelTerms,
	}
}

// Calculate applies the Frieze formula to the complex number z.
func (r *Frieze) Calculate(coordinate complex128) complex128 {
	sumOfTermCalculations := complex(0,0)
	for _, term := range r.formulaLevelTerms {
		termCalculation := r.calculateTerm(term, coordinate)
		sumOfTermCalculations += termCalculation
	}

	return sumOfTermCalculations
}

func (r *Frieze) calculateTerm(term Term, coordinate complex128) complex128 {
	sum := complex(0.0, 0.0)

	coefficientRelationships := []coefficient.Relationship{coefficient.PlusNPlusM}
	coefficientRelationships = append(coefficientRelationships, term.CoefficientRelationships...)
	coefficientSets := coefficient.Pairing{
		PowerN: term.PowerN,
		PowerM: term.PowerM,
	}.GenerateCoefficientSets(coefficientRelationships)

	for _, relationshipSet := range coefficientSets {
		multiplier := term.Multiplier
		if relationshipSet.NegateMultiplier == true {
			multiplier *= -1
		}
		sum += CalculateEulerTerm(coordinate, relationshipSet.PowerN, relationshipSet.PowerM, multiplier, term.IgnoreComplexConjugate)
	}
	return sum
}

// CalculateEulerTerm calculates e^(i*n*z) * e^(-i*m*zConj)
func CalculateEulerTerm(z complex128, power1, power2 int, scale complex128, ignoreComplexConjugate bool) complex128 {
	eRaisedToTheNZi := cmplx.Exp(complex(0, 1) * z * complex(float64(power1), 0))
	if ignoreComplexConjugate {
		return eRaisedToTheNZi * scale
	}

	complexConjugate := complex(real(z), -1*imag(z))
	eRaisedToTheNegativeMZConji := cmplx.Exp(complexConjugate * complex(0, -1*float64(power2)))
	return eRaisedToTheNZi * eRaisedToTheNegativeMZConji * scale
}


// FormulaLevelTerms returns the Terms this formula will use.
func (r *Frieze) FormulaLevelTerms() []Term {
	return r.formulaLevelTerms
}