package formula

import (
	"errors"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
)

// Rhombic formulas will transform points by returning the same coordinates.
type Rhombic struct {
	latticeVectors []complex128
	wavePackets    []WavePacket
}

// NewRhombicFormula returns a new formula object.
func NewRhombicFormula(packets []WavePacket, latticeHeight float64) (*Rhombic, error) {
	if latticeHeight == 0.0 {
		return nil, errors.New("rhombic lattice must specify height")
	}

	packetsWithLockedCoefficients := lockTermsBasedOnRelationship(
		[]coefficient.Relationship{
			coefficient.PlusMPlusN,
		},
		packets)

	return &Rhombic{
			wavePackets: packetsWithLockedCoefficients,
			latticeVectors: []complex128{
				complex(0.5, latticeHeight),
				complex(0.5, latticeHeight*-1),
			},
		},
		nil
}

// WavePackets returns the wave packets used.
func (r *Rhombic) WavePackets() []WavePacket {
	return r.wavePackets
}

// Calculate transforms the coordinate using the Rhombic lattice's wave packets.
func (r *Rhombic) Calculate(coordinate complex128) complex128 {
	return CalculateCoordinateUsingWavePackets(coordinate, r.LatticeVectors(), r.WavePackets())
}

// FormulaLevelTerms returns an empty list, Rhombic formulas do not have base-level terms.
func (r *Rhombic) FormulaLevelTerms() []Term {
	return nil
}

// LatticeVectors returns the lattice vectors used to create the rectangle.
func (r *Rhombic) LatticeVectors() []complex128 {
	return r.latticeVectors
}

// SymmetriesFound returns all symmetries found in this pattern.
func (r *Rhombic) SymmetriesFound() []Symmetry {
	return nil
}
