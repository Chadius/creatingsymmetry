package formula

import (
	"errors"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"math"
)

// Hexagonal formulas will transform points by returning the same coordinates.
type Hexagonal struct {
	latticeVectors []complex128
	wavePackets    []WavePacket
}

// NewHexagonalFormula returns a new formula object.
func NewHexagonalFormula(packets []WavePacket, desiredSymmetry Symmetry) (*Hexagonal, error) {
	if desiredSymmetry != "" &&
		desiredSymmetry != P1 &&
		desiredSymmetry != P3 &&
		desiredSymmetry != P31m &&
		desiredSymmetry != P3m1 &&
		desiredSymmetry != P6 &&
		desiredSymmetry != P6m {
		return nil, errors.New("hexagonal lattice can apply these desired symmetries: P1, P3, P31m, P3m1, P6, P6m")
	}

	wavePacketsWithDesiredSymmetry := createNewWavePacketsBasedOnDesiredSymmetry(packets, desiredSymmetry)

	packetsWithLockedCoefficients := lockTermsBasedOnRelationship(
		[]coefficient.Relationship{
			coefficient.PlusMMinusSumNAndM,
			coefficient.MinusSumNAndMPlusN,
		},
		wavePacketsWithDesiredSymmetry)

	return &Hexagonal{
			wavePackets: packetsWithLockedCoefficients,
			latticeVectors: []complex128{
				complex(1, 0),
				complex(-0.5, math.Sqrt(3.0)/2.0),
			},
		},
		nil
}

// WavePackets returns the wave packets used.
func (r *Hexagonal) WavePackets() []WavePacket {
	return r.wavePackets
}

// Calculate transforms the coordinate using the Hexagonal lattice's wave packets.
func (r *Hexagonal) Calculate(coordinate complex128) complex128 {
	return CalculateCoordinateUsingWavePackets(coordinate, r.LatticeVectors(), r.WavePackets())
}

// FormulaLevelTerms returns an empty list, Hexagonal formulas do not have base-level terms.
func (r *Hexagonal) FormulaLevelTerms() []Term {
	return nil
}

// LatticeVectors returns the lattice vectors used to create the rectangle.
func (r *Hexagonal) LatticeVectors() []complex128 {
	return r.latticeVectors
}

// SymmetriesFound returns all symmetries found in this pattern.
func (r *Hexagonal) SymmetriesFound() []Symmetry {
	symmetriesFound := []Symmetry{P1, P3}

	expectedCoefficientsBySymmetry := map[Symmetry][]coefficient.Relationship{
		P31m: {coefficient.PlusMPlusN},
		P3m1: {coefficient.MinusMMinusN},
		P6:   {coefficient.MinusNMinusM},
		P6m: {
			coefficient.MinusNMinusM,
			coefficient.MinusMMinusN,
			coefficient.PlusMPlusN,
		},
	}

	order := []Symmetry{P31m, P3m1, P6, P6m}
	for _, symmetryType := range order {
		coefficients := expectedCoefficientsBySymmetry[symmetryType]
		if HasSymmetry(r.WavePackets(), symmetryType, map[Symmetry][]coefficient.Relationship{
			symmetryType: coefficients,
		}) {
			symmetriesFound = append(symmetriesFound, symmetryType)
		}
	}

	return symmetriesFound
}
