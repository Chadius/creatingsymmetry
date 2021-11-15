package formula

import (
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"math/cmplx"
)

// Rosette formulas transform points around a central origin, similar to a rosette surrounding a center.
type Rosette struct {
	formulaLevelTerms []Term
}

// NewRosetteFormula returns a new formula
func NewRosetteFormula(formulaLevelTerms []Term) (*Rosette, error) {
	return &Rosette{
		formulaLevelTerms: formulaLevelTerms,
	},
	nil
}

// WavePackets returns an empty array, this type of formula does not use WavePackets.
func (r *Rosette) WavePackets() []WavePacket {
	return nil
}

// Calculate applies the Rosette formula to the complex number z.
func (r *Rosette) Calculate(coordinate complex128) complex128 {
	sumOfTermCalculations := complex(0,0)
	for _, term := range r.formulaLevelTerms {
		termCalculation := r.calculateTerm(term, coordinate)
		sumOfTermCalculations += termCalculation
	}

	return sumOfTermCalculations
}

func (r *Rosette) calculateTerm(term Term, coordinate complex128) complex128 {
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
		sum += CalculateExponentTerm(coordinate, relationshipSet.PowerN, relationshipSet.PowerM, multiplier, term.IgnoreComplexConjugate)
	}
	return sum
}

// CalculateExponentTerm calculates (z^power * zConj^conjugatePower)
//   where z is a complex number, zConj is the complex conjugate
//   and power and conjugatePower are integers.
func CalculateExponentTerm(coordinate complex128, power1, power2 int, scale complex128, ignoreComplexConjugate bool) complex128 {
	zRaisedToN := cmplx.Pow(coordinate, complex(float64(power1), 0))
	if ignoreComplexConjugate {
		return zRaisedToN * scale
	}

	complexConjugate := complex(real(coordinate), -1*imag(coordinate))
	complexConjugateRaisedToM := cmplx.Pow(complexConjugate, complex(float64(power2), 0))
	return zRaisedToN * complexConjugateRaisedToM * scale
}

// FormulaLevelTerms returns the Terms this formula will use.
func (r *Rosette) FormulaLevelTerms() []Term {
	return r.formulaLevelTerms
}

// LatticeVectors returns an empty list, this formula does not use them
func (r *Rosette) LatticeVectors() []complex128 {
	return nil
}