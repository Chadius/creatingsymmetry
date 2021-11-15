package formula

import (
	"errors"
)

// Rectangular formulas will transform points by returning the same coordinates.
type Rectangular struct {
	latticeVectors []complex128
	wavePackets []WavePacket
}

// NewRectangularFormula returns a new formula object.
func NewRectangularFormula(packets []WavePacket, latticeHeight float64) (*Rectangular, error) {
	if latticeHeight == 0.0 {
		return nil, errors.New("rectangular lattice must specify height")
	}

	return &Rectangular{
		wavePackets: packets,
		latticeVectors: []complex128{
			complex(1, 0),
			complex(0, latticeHeight),
		},
	},
	nil
}

// WavePackets returns the wave packets used.
func (r *Rectangular) WavePackets() []WavePacket {
	return r.wavePackets
}

// Calculate transforms the coordinate using the Rectangular lattice's wave packets.
func (r *Rectangular) Calculate(coordinate complex128) complex128 {
	return coordinate
}

// FormulaLevelTerms returns an empty list, Rectangular formulas do not have base-level terms.
func (r *Rectangular) FormulaLevelTerms() []Term {
	return nil
}

// LatticeVectors returns the lattice vectors used to create the rectangle.
func (r *Rectangular) LatticeVectors() []complex128 {
	return r.latticeVectors
}