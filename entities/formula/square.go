package formula

import (
	"errors"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
)

// Square formulas will transform points by returning the same coordinates.
type Square struct {
	latticeVectors []complex128
	wavePackets    []WavePacket
}

// NewSquareFormula returns a new formula object.
func NewSquareFormula(packets []WavePacket, desiredSymmetry Symmetry) (*Square, error) {
	if desiredSymmetry != "" &&
		desiredSymmetry != P1 &&
		desiredSymmetry != P4 &&
		desiredSymmetry != P4m &&
		desiredSymmetry != P4g {
		return nil, errors.New("square lattice can apply these desired symmetries: P1, P4, P4m, P4g")
	}

	wavePacketsWithDesiredSymmetry := createNewWavePacketsBasedOnDesiredSymmetry(packets, desiredSymmetry)

	packetsWithLockedCoefficients := lockTermsBasedOnRelationship(
		[]coefficient.Relationship{
			coefficient.PlusMMinusN,
			coefficient.MinusNMinusM,
			coefficient.MinusMPlusN,
		},
		wavePacketsWithDesiredSymmetry)

	squareWallpaperFormula := &Square{
		wavePackets: packetsWithLockedCoefficients,
		latticeVectors: []complex128{
			complex(1, 0),
			complex(0, 1),
		},
	}

	return squareWallpaperFormula, nil
}

// WavePackets returns the wave packets used.
func (r *Square) WavePackets() []WavePacket {
	return r.wavePackets
}

// Calculate transforms the coordinate using the Square lattice's wave packets.
func (r *Square) Calculate(coordinate complex128) complex128 {
	return CalculateCoordinateUsingWavePackets(coordinate, r.LatticeVectors(), r.WavePackets())
}

// FormulaLevelTerms returns an empty list, Square formulas do not have base-level terms.
func (r *Square) FormulaLevelTerms() []Term {
	return nil
}

// LatticeVectors returns the lattice vectors used to create the rectangle.
func (r *Square) LatticeVectors() []complex128 {
	return r.latticeVectors
}

// SymmetriesFound returns all symmetries found in this pattern.
func (r *Square) SymmetriesFound() []Symmetry {
	symmetriesFound := []Symmetry{P1, P4}

	expectedCoefficientsBySymmetry := map[Symmetry][]coefficient.Relationship{
		P4m: {coefficient.PlusMPlusN},
		P4g: {coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum},
	}

	order := []Symmetry{P4m, P4g}
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
