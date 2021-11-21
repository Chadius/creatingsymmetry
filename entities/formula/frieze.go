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
func NewFriezeFormula(formulaLevelTerms []Term) (*Frieze, error) {
	return &Frieze{
		formulaLevelTerms: formulaLevelTerms,
	},
	nil
}

// WavePackets returns an empty array, this type of formula does not use WavePackets.
func (f *Frieze) WavePackets() []WavePacket {
	return nil
}

// Calculate applies the Frieze formula to the complex number z.
func (f *Frieze) Calculate(coordinate complex128) complex128 {
	sumOfTermCalculations := complex(0,0)
	for _, term := range f.formulaLevelTerms {
		termCalculation := f.calculateTerm(term, coordinate)
		sumOfTermCalculations += termCalculation
	}

	return sumOfTermCalculations
}

func (f *Frieze) calculateTerm(term Term, coordinate complex128) complex128 {
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


// FormulaLevelTerms returns the terms this formula will use.
func (f *Frieze) FormulaLevelTerms() []Term {
	return f.formulaLevelTerms
}

// LatticeVectors returns an empty list, this formula does not use them
func (f *Frieze) LatticeVectors() []complex128 {
	return nil
}