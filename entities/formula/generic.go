package formula

import (
	"errors"
	"github.com/chadius/creatingsymmetry/entities/formula/coefficient"
)

// TODO Do I need a formula wide multiplier?

// Generic formulas will transform points by returning the same coordinates.
type Generic struct {
	latticeVectors []complex128
	wavePackets    []WavePacket
}

// NewGenericFormula returns a new formula object.
func NewGenericFormula(packets []WavePacket, latticeWidth, latticeHeight float64, desiredSymmetry Symmetry) (*Generic, error) {
	if latticeHeight == 0.0 || latticeWidth == 0.0 {
		return nil, errors.New("generic lattice must specify dimensions")
	}

	if desiredSymmetry != P1 && desiredSymmetry != P2 && desiredSymmetry != "" {
		return nil, errors.New("generic lattice can apply these desired symmetries: P1, P2")
	}

	newWavePackets := createNewWavePacketsBasedOnDesiredSymmetry(packets, desiredSymmetry)

	return &Generic{
			wavePackets: newWavePackets,
			latticeVectors: []complex128{
				complex(1, 0),
				complex(latticeWidth, latticeHeight),
			},
		},
		nil
}

// WavePackets returns the wave packets used.
func (r *Generic) WavePackets() []WavePacket {
	return r.wavePackets
}

// Calculate transforms the coordinate using the Generic lattice's wave packets.
func (r *Generic) Calculate(coordinate complex128) complex128 {
	return CalculateCoordinateUsingWavePackets(coordinate, r.LatticeVectors(), r.WavePackets())
}

// FormulaLevelTerms returns an empty list, Generic formulas do not have base-level terms.
func (r *Generic) FormulaLevelTerms() []Term {
	return nil
}

// LatticeVectors returns the lattice vectors used to create the rectangle.
func (r *Generic) LatticeVectors() []complex128 {
	return r.latticeVectors
}

// SymmetriesFound returns all symmetries found in this pattern.
func (r *Generic) SymmetriesFound() []Symmetry {
	symmetriesFound := []Symmetry{P1}

	if HasSymmetry(r.WavePackets(), P2, map[Symmetry][]coefficient.Relationship{
		P2: {coefficient.MinusNMinusM},
	}) {
		symmetriesFound = append(symmetriesFound, P2)
	}

	return symmetriesFound
}
